package api

import (
	"../external"
	"../helpers"
	"github.com/k0kubun/pp"
)

func init() {

	users, err := helpers.GetAllUsers()
	if err != nil {
		panic(err)
	}
	points, err := external.GetPlayersPointsFromCSV("./points/puntaje_fecha_29.csv")
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		external.FillUserTeamWithPlayerPoints(user, points)
		if err != nil {
			panic(err)
		}

		upUser, err := helpers.UpdateUser(user)
		if err != nil {
			panic(err)
		}

		if len(user.PlayingTeams) >0 {
			pp.Println(upUser.PlayingTeams[0].Points)
		}
	}

}
