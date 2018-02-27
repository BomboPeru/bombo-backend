package api

import (
	"log"

	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"../helpers"
)


// LinkWithPlayerType ...
func LinkWithPlayerType(api iris.Party) {
	api.Get("/player/{id:string}", func(c iris.Context) {
		playerID := c.Params().Get("ID")
		ID, err := uuid.FromString(playerID)
		if err != nil {
			log.Println("uuid.FromString(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		player, err := helpers.GetPlayerByID(ID)
		if err != nil {
			log.Println("GetPlayerByID(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
		} else {
			c.StatusCode(iris.StatusOK)
		}

		c.JSON(Response{
			Data:  player,
			Error: err.Error(),
		})

	})

	api.Get("/player/all", func(c iris.Context) {
		players, err := helpers.GetAllPlayers()
		if err != nil {
			log.Println("GetAllPlayers(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
		} else {
			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data:  players,
				Error: nil,
			})
		}

	})

	api.Post("/player/create", func(c iris.Context) {
		player := new(helpers.Player)
		err := c.ReadJSON(player)
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
		if uuid.Equal(player.ID, uuid.Nil){
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
			player.ID = newID
		}

		returnedPlayer, err := helpers.CreateNewPlayer(player)
		if err != nil {
			log.Println("CreateNewPlayer(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  returnedPlayer,
			Error: err.Error(),
		})
	})

	api.Post("/player/update", func(c iris.Context) {
		player := new(helpers.Player)
		err := c.ReadJSON(player)
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
		if uuid.Equal(player.ID, uuid.Nil){
			log.Println("Update method needs Id")
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		updatedPlayer, err := helpers.UpdatePlayer(player)
		if err != nil {
			log.Println("UpdatePlayer(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  updatedPlayer,
			Error: err.Error(),
		})

	})

	api.Post("/player/delete", func(c iris.Context) {
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

		returnedPlayer, err := helpers.DeletePlayerByID(idForm.ID)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
		} else {
			c.StatusCode(iris.StatusOK)
		}

		c.JSON(Response{
			Data:  returnedPlayer,
			Error: err.Error(),
		})

	})
}
