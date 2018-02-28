package helpers

import (
	"github.com/asdine/storm"
	"github.com/satori/go.uuid"
	"log"
	"encoding/json"
)

// Player ...
type Player struct {
    ID uuid.UUID `json:"id" form:"id"`
    Name string `json:"name" form:"name"`
    JNumber string `json:"j_number" form:"j_number"`
    Points string `json:"points" form:"points"`
    Team string `json:"team"`
    Cost int `json:"cost"`
	Nation     string `json:"nation"`
}

// GetPlayerByID ...
func GetPlayerByID(id uuid.UUID) (*Player, error) {
	player := new(Player)

	if err := DbPlayer.One("ID", id, player); err != nil {
		if err == storm.ErrNotFound {
			log.Println("Player not found")
			return player, err
		}
		log.Println("DbPlayer.Find(), ", err.Error())
		return player, err
	}
	return player, nil
}

// GetPlayerByAnyField ...
func GetPlayerByAnyField(field string, value interface{}) (*Player, error) {
	player := new(Player)

	if err := DbPlayer.One(field, value, player); err != nil {
		if err == storm.ErrNotFound {
			log.Println("Player not found")
			return player, err
		}
		log.Println("DbPlayer.Find(), ", err.Error())
		return player, err
	}
	return player, nil
}

// GetAllPlayers ...
func GetAllPlayers() ([]*Player, error) {
	players := make([]*Player, 0)
	err := DbPlayer.All(&players)
	if err != nil {
		return players, err
	}
	return players, nil
}


// CreateNewPlayerWithRawData ...
func CreateNewPlayerWithRawData(data []byte) (*Player, error) {
	player := new(Player)
	err := json.Unmarshal(data, player)
	if err != nil {
		log.Println("json.Unmarshal(), ", err.Error())
		return player, err
	}
	err = DbPlayer.Save(player)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			log.Println("Player already exist")
			return player, err
		}
		log.Println("json.Unmarshal(), ", err.Error())
		return player, err
	}

	return player, nil
}

// CreateNewPlayer ...
func CreateNewPlayer(player *Player) (*Player, error) {

	err := DbPlayer.Save(player)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			log.Println("Player already exist")
			return player, err
		}
		log.Println("json.Unmarshal(), ", err.Error())
		return player, err
	}

	return player, nil
}

// DeletePlayerByID ...
func DeletePlayerByID(id uuid.UUID) (*Player, error) {

	player, err := GetPlayerByID(id)
	if err != nil {
		log.Println("GetPlayerByID(), ", err.Error())
		return player, err
	}

	err = DbPlayer.DeleteStruct(player)
	if err != nil {
		log.Println("DeleteStruct(), ", err.Error())
		return player, err
	}

	return player, nil

}

// DeletePlayerByStruct ...
func DeletePlayerByStruct(player *Player) (*Player, error) {
	err := DbPlayer.DeleteStruct(player)
	if err != nil {
		log.Println("DeleteStruct(), ", err.Error())
		return player, err
	}

	return player, nil
}


// UpdatePlayer ...
func UpdatePlayer(player *Player) (*Player, error) {
	err := DbPlayer.Update(player)
	if err != nil {
		log.Println("DbPlayer.Update(), ", err.Error())
		return player, err
	}

	finalPlayer, err := GetPlayerByID(player.ID)
	if err != nil {
		log.Println("GetPlayerByID(), ", err.Error())
		return finalPlayer, err
	}
	return finalPlayer, nil
}