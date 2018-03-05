package api

import (
	"github.com/kataras/iris"
	"../helpers"
	"sort"
)

type Position struct {
	Place int `json:"place"`
	User string `json:"user"`
	TeamName string `json:"team_name"`
	Points float64 `json:"points"`
}

type byPoints []*Position

func (s byPoints) Len() int {
	return len(s)
}
func (s byPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPoints) Less(i, j int) bool {
	return s[i].Points < s[j].Points
}

func LinkExtrasUtil(api iris.Party) {
	api.Get("/extra/get-ranking", func(c iris.Context) {
		users, err := helpers.GetAllUsers()
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"data": nil,
				"error": err,
			})
			return
		}
		allPlayingPositions := make([]*Position, 0)
		for _, user := range users {
			for _, team := range user.PlayingTeams {
				allPlayingPositions = append(allPlayingPositions, &Position{
					User: user.Username,
					TeamName: team.Name,
					Place: 0,
					Points: team.Points,
				})
			}
		}

		sort.Sort(byPoints(allPlayingPositions))
		offset := len(allPlayingPositions)
		for i, p := range allPlayingPositions {
			p.Place = offset-i
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"data": allPlayingPositions,
			"error": nil,
		})
	})
}
