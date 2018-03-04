package main

import (
	"fmt"
	"github.com/kataras/iris"
	"./api"
	"./helpers"
	"github.com/iris-contrib/middleware/cors"
	"github.com/satori/go.uuid"
	"./external"
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

	app.Use(crs)

	apiParty := app.Party(fmt.Sprintf("/api/v%s", apiVersion))

	
	api.LinkWithUserType(apiParty)
	
	api.LinkWithPlayerType(apiParty)
	
	api.LinkWithVirtualTeamType(apiParty)

	app.Logger().SetLevel("debug")


	//app.Run(iris.Addr(":8080"))

	id, _ := uuid.FromString("08e4b723-8e19-4828-8604-0df593dcad62")
	user, err := helpers.GetUserByID(id)
	if err != nil {
		panic(err)
	}

	events, err := external.GetAllActiveEvents()
	if err != nil {
		panic(err)
	}

	external.FillTeamsUser(user, events)


}
