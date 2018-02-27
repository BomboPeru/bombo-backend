package api

import (
	"log"

	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"../helpers"
)


// LinkWithVirtualTeamType ...
func LinkWithVirtualTeamType(api iris.Party) {
	api.Get("/virtual_team/{id:string}", func(c iris.Context) {
		virtualteamID := c.Params().Get("ID")
		ID, err := uuid.FromString(virtualteamID)
		if err != nil {
			log.Println("uuid.FromString(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		virtualteam, err := helpers.GetVirtualTeamByID(ID)
		if err != nil {
			log.Println("GetVirtualTeamByID(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
		} else {
			c.StatusCode(iris.StatusOK)
		}

		c.JSON(Response{
			Data:  virtualteam,
			Error: err.Error(),
		})

	})

	api.Post("/virtual_team/create", func(c iris.Context) {
		virtualteam := new(helpers.VirtualTeam)
		err := c.ReadJSON(virtualteam)
		if err != nil {
			log.Println("c.ReadJSON(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		// If not have Id
		if virtualteam.ID.String() == "" {
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
			virtualteam.ID = newID
		}

		returnedVirtualTeam, err := helpers.CreateNewVirtualTeam(virtualteam)
		if err != nil {
			log.Println("CreateNewVirtualTeam(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  returnedVirtualTeam,
			Error: err.Error(),
		})
	})

	api.Post("/virtual_team/update", func(c iris.Context) {
		virtualteam := new(helpers.VirtualTeam)
		err := c.ReadJSON(virtualteam)
		if err != nil {
			log.Println("c.ReadJSON(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		// If not have Id
		if virtualteam.ID.String() == "" {
			log.Println("Update method needs Id")
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		updatedVirtualTeam, err := helpers.UpdateVirtualTeam(virtualteam)
		if err != nil {
			log.Println("UpdateVirtualTeam(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  updatedVirtualTeam,
			Error: err.Error(),
		})

	})

	api.Post("/virtual_team/delete", func(c iris.Context) {
		type IDForm struct {
			ID uuid.UUID `json:"id"`
		}

		idForm := new(IDForm)
		err := c.ReadJSON(idForm)
		if err != nil {
			log.Println("c.ReadJSON(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		returnedVirtualTeam, err := helpers.DeleteVirtualTeamByID(idForm.ID)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
		} else {
			c.StatusCode(iris.StatusOK)
		}

		c.JSON(Response{
			Data:  returnedVirtualTeam,
			Error: err.Error(),
		})

	})
}
