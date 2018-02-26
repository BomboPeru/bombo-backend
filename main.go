package main

import (
	"fmt"
	"github.com/kataras/iris"
	"./api"
)

const apiVersion = "1.0"

func main() {
	app := iris.Default()

	apiParty := app.Party(fmt.Sprintf("/api/v%s", apiVersion))

	
	api.LinkWithUserType(apiParty)
	
	api.LinkWithPlayerType(apiParty)
	
	api.LinkWithVirtualTeamType(apiParty)

	app.Logger().SetLevel("debug")

	app.Run(iris.Addr(":8080"))

}
