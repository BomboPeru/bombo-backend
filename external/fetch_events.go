package external

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const getAllEndpoints = "http://open.bombo.pe/api/v1.0/get-endpoints"
const baseOpenBombo = "http://open.bombo.pe"

type MiniWork struct {
	Name string `json:"name"`
	URL string `json:"url"`
	Type string `json:"type"`
	Endpoint string `json:"endpoint"`
}

func GetAllActiveEvents() ([]*MatchEvents, error) {
	responseEndpoints, err := http.Get(getAllEndpoints)
	type structEndpointResponse struct {
		Data []MiniWork `json:"data"`
		Error error `json:"error"`
	}

	dataEndpoints, err := ioutil.ReadAll(responseEndpoints.Body)
	if err != nil {
		return []*MatchEvents{}, err
	}

	endpoints := new(structEndpointResponse)

	err = json.Unmarshal(dataEndpoints, endpoints)
	if err != nil {
		return []*MatchEvents{}, err
	}

	allEvents := make([]*MatchEvents, 0)

	for _, mw := range endpoints.Data {
		finalEndpoint := baseOpenBombo + mw.Endpoint
		response , err := http.Get(finalEndpoint)
		if err != nil {
			return []*MatchEvents{}, err
		}

		type structResponse struct {
			Data *MatchEvents `json:"data"`
			Error error `json:"error"`
		}
		resp := new(structResponse)

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []*MatchEvents{}, err
		}

		err = json.Unmarshal(data, resp)
		if err != nil {
			return []*MatchEvents{}, err
		}

		allEvents = append(allEvents, resp.Data)
	}

	return allEvents, nil

}
