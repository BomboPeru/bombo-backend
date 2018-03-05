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

	app.Run(iris.Addr(":8080"))

}
