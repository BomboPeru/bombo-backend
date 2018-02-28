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

		// If virtualTeam not have Id
		if uuid.Equal(payload.Team.ID, uuid.Nil){
			newID, err := uuid.NewV4()
			if err != nil {
				log.Println("uuid.NewV4(), ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: err.Error(),
				})
				return
			}
			payload.Team.ID = newID
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
				Error: errors.New("invalid team type").Error(),
			})
			break
		}


		finalUser, err := helpers.UpdateUser(user)
		if err !=nil {
			log.Println("helpers.UpdateUser(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data: finalUser,
			Error: nil,
		})

	})


	api.Post("/user/remove-team/{type:string}", func(c iris.Context) {
		virtualTeamType := c.Params().Get("type")

		type AddedTeam struct {
			UserId uuid.UUID `json:"user_id"`
			TeamId uuid.UUID `json:"team_id"`
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
			index := helpers.GetIndexOfTeamById(user.OldTeams, payload.TeamId)
			if index == -1 {
				log.Println("Team not exist, ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: errors.New("team id not exist").Error(),
				})
				return
			}

			user.OldTeams = append(
				user.OldTeams[:index],
				user.OldTeams[index+1:]...
			)

			rUser, err := helpers.UpdateUser(user)

			if err != nil {
				log.Println("helpers.UpdateUser(), ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: err.Error(),
				})
				return
			}

			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data: rUser,
				Error: nil,
			})
			break
		case "saved":
			index := helpers.GetIndexOfTeamById(user.SavedTeams, payload.TeamId)
			if index == -1 {
				log.Println("Team not exist, ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: errors.New("team id not exist").Error(),
				})
				return
			}

			user.SavedTeams = append(
				user.SavedTeams[:index],
				user.SavedTeams[index+1:]...
			)

			rUser, err := helpers.UpdateUser(user)

			if err != nil {
				log.Println("helpers.UpdateUser(), ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: err.Error(),
				})
				return
			}

			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data: rUser,
				Error: nil,
			})
			break
		case "playing":
			index := helpers.GetIndexOfTeamById(user.PlayingTeams, payload.TeamId)
			if index == -1 {
				log.Println("Team not exist, ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: errors.New("team id not exist").Error(),
				})
				return
			}

			user.PlayingTeams = append(
				user.PlayingTeams[:index],
				user.PlayingTeams[index+1:]...
			)

			rUser, err := helpers.UpdateUser(user)

			if err != nil {
				log.Println("helpers.UpdateUser(), ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: err.Error(),
				})
				return
			}

			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data: rUser,
				Error: nil,
			})
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
