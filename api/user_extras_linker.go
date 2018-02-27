package api

import (
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"../helpers"
	"log"
	"errors"
)

func ExtraUserEndpoints(api iris.Party) {
	api.Post("/user/add-team/{type:string}", func(c iris.Context) {
		virtualTeamType := c.Params().Get("type")

		type AddedTeam struct {
			UserId uuid.UUID `json:"user_id"`
			Team *helpers.VirtualTeam `json:"team"`
		}

		payload := new(AddedTeam)

		err := c.ReadJSON(payload)
		if err != nil {
			log.Println("c.ReadJSON(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		teamResponse, err := helpers.CreateNewVirtualTeam(payload.Team)
		if err != nil {
			log.Println("helpers.CreateNewVirtualTeam(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		user, err := helpers.GetUserByID(payload.UserId)
		if err != nil {
			log.Println("helpers.GetUserByID(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		switch virtualTeamType {
		case "old":
			user.OldTeams = append(user.OldTeams, teamResponse)
			break
		case "saved":
			user.SavedTeams = append(user.SavedTeams, teamResponse)
			break
		case "playing":
			user.PlayingTeams = append(user.PlayingTeams, teamResponse)
			break
		default:
			log.Println("INVALID TEAM TYPE", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: errors.New("invalid team name").Error(),
			})
			break
		}


	})
}
