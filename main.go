package main

import (
	"fmt"
	"github.com/kataras/iris"
	"./api"
	"github.com/iris-contrib/middleware/cors"
)

const apiVersion = "1.0"

func main() {
	app := iris.Default()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"},
		AllowCredentials: true,
	})

	crsConf := cors.NewAllowAllAppMiddleware()

	app.Configure(crsConf)
	app.Use(crs)


	apiParty := app.Party(fmt.Sprintf("/api/v%s", apiVersion))

	
	api.LinkWithUserType(apiParty)
	
	api.LinkWithPlayerType(apiParty)
	
	api.LinkWithVirtualTeamType(apiParty)

	api.LinkExtrasUtil(apiParty)

	app.Logger().SetLevel("debug")
	//
	app.Run(iris.Addr(":8080"))

	////jaimePoints, err := external.GetPlayersPointsFromCSV("./points/puntaje_fecha_29.csv")
	////if err != nil {
	////	panic(err)
	////}
	////id, _ := uuid.FromString("08e4b723-8e19-4828-8604-0df593dcad62")
	////user, err := helpers.GetUserByID(id)
	//
	//
	//
	//resp, err := http.Get("http://open.bombo.pe/api/v1.0/premier_league")
	//if err != nil {
	//	panic(err)
	//}
	//
	//data, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//type Response struct {
	//	Data external.League `json:"data"`
	//	Error error `json:"error"`
	//}
	//pLeague := new(Response)
	//
	//err = json.Unmarshal(data, pLeague)
	//if err != nil {
	//	panic(err)
	//}
	//
	//events, err := external.GetAllActiveEvents()
	//if err != nil {
	//	panic(err)
	//}
	//
	//bregyPoints := external.GetPlayersPointsFromPlayersList(pLeague.Data, events)
	//
	//pp.Println(bregyPoints[0:10])
}
