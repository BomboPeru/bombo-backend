package helpers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/asdine/storm"
	"github.com/satori/go.uuid"
)

type Players struct {
	Name       string    `json:"name"`
	GoalKeeper []*Player `json:"goal_keeper"`
	MidFielder []*Player `json:"mid_fielder"`
	Defender   []*Player `json:"defender"`
	Forward    []*Player `json:"forward"`
	Coach      *Player   `json:"coach"`
}

// VirtualTeam ...
type VirtualTeam struct {
	ID        uuid.UUID `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Players   Players   `json:"players" form:"players"`
	Balance   int       `json:"balance" form:"balance"`
	LeagueImg string    `json:"league_img" form:"league_img"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	Ranking   int       `json:"ranking" form:"ranking"`
	Points    int       `json:"points" form:"points"`
}

// GetVirtualTeamByID ...
func GetVirtualTeamByID(id uuid.UUID) (*VirtualTeam, error) {
	virtualteam := new(VirtualTeam)

	if err := DbVirtualTeam.One("ID", id, virtualteam); err != nil {
		if err == storm.ErrNotFound {
			log.Println("VirtualTeam not found")
			return virtualteam, err
		}
		log.Println("DbVirtualTeam.Find(), ", err.Error())
		return virtualteam, err
	}
	return virtualteam, nil
}

// GetVirtualTeamByAnyField ...
func GetVirtualTeamByAnyField(field string, value interface{}) (*VirtualTeam, error) {
	virtualteam := new(VirtualTeam)

	if err := DbVirtualTeam.One(field, value, virtualteam); err != nil {
		if err == storm.ErrNotFound {
			log.Println("VirtualTeam not found")
			return virtualteam, err
		}
		log.Println("DbVirtualTeam.Find(), ", err.Error())
		return virtualteam, err
	}
	return virtualteam, nil
}

// GetAllVirtualTeams ...
func GetAllVirtualTeams() ([]*VirtualTeam, error) {
	virtualteams := make([]*VirtualTeam, 0)
	err := DbVirtualTeam.All(&virtualteams)
	if err != nil {
		return virtualteams, err
	}
	return virtualteams, nil
}

// CreateNewVirtualTeamWithRawData ...
func CreateNewVirtualTeamWithRawData(data []byte) (*VirtualTeam, error) {
	virtualteam := new(VirtualTeam)
	err := json.Unmarshal(data, virtualteam)
	if err != nil {
		log.Println("json.Unmarshal(), ", err.Error())
		return virtualteam, err
	}
	err = DbVirtualTeam.Save(virtualteam)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			log.Println("VirtualTeam already exist")
			return virtualteam, err
		}
		log.Println("json.Unmarshal(), ", err.Error())
		return virtualteam, err
	}

	return virtualteam, nil
}

// CreateNewVirtualTeam ...
func CreateNewVirtualTeam(virtualteam *VirtualTeam) (*VirtualTeam, error) {

	err := DbVirtualTeam.Save(virtualteam)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			log.Println("VirtualTeam already exist")
			return virtualteam, err
		}
		log.Println("json.Unmarshal(), ", err.Error())
		return virtualteam, err
	}

	return virtualteam, nil
}

// DeleteVirtualTeamByID ...
func DeleteVirtualTeamByID(id uuid.UUID) (*VirtualTeam, error) {

	virtualteam, err := GetVirtualTeamByID(id)
	if err != nil {
		log.Println("GetVirtualTeamByID(), ", err.Error())
		return virtualteam, err
	}

	err = DbVirtualTeam.DeleteStruct(virtualteam)
	if err != nil {
		log.Println("DeleteStruct(), ", err.Error())
		return virtualteam, err
	}

	return virtualteam, nil

}

// DeleteVirtualTeamByStruct ...
func DeleteVirtualTeamByStruct(virtualteam *VirtualTeam) (*VirtualTeam, error) {
	err := DbVirtualTeam.DeleteStruct(virtualteam)
	if err != nil {
		log.Println("DeleteStruct(), ", err.Error())
		return virtualteam, err
	}

	return virtualteam, nil
}

// UpdateVirtualTeam ...
func UpdateVirtualTeam(virtualteam *VirtualTeam) (*VirtualTeam, error) {
	err := DbVirtualTeam.Update(virtualteam)
	if err != nil {
		log.Println("DbVirtualTeam.Update(), ", err.Error())
		return virtualteam, err
	}

	finalTeam, err := GetVirtualTeamByID(virtualteam.ID)
	if err != nil {
		log.Println("GetVirtualTeamByID(), ", err.Error())
		return finalTeam, err
	}

	return finalTeam, nil
}
