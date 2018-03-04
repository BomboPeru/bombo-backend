package external

import (
	"../helpers"
	"regexp"
	"strings"
	"fmt"
	"../util"
	"strconv"
	"github.com/k0kubun/pp"
)

const second = "suplentes"
const principal = "alineaciones_iniciales"


const goalKeeper = "goalKeeper"
const midFielder = "midFielder"
const defender = "defender"
const forward = "forward"


/*
type Players struct {
	Name       string    `json:"name"`
	GoalKeeper []*Player `json:"goal_keeper"`
	MidFielder []*Player `json:"mid_fielder"`
	Defender   []*Player `json:"defender"`
	Forward    []*Player `json:"forward"`
	Coach      *Player   `json:"coach"`
}
*/
func GetPointsToPlayerByEvent(teamScore int, typePlaying string, position string, event *Event) int {
	points := 0

	if typePlaying == principal {
		points += 2
	}

	switch event.Type {
	case "substitution-in":
		points += 1
	case "substitution-out":
		points += -1
	case "soccer-ball":
		assistSearcher, _ := regexp.Compile(`\([^\n\r)]+\)`)
		as := assistSearcher.Find([]byte(event.Metadata))
		assist := strings.Replace(string(as), "(", "", -1)
		assist = strings.Replace(assist, ")", "", -1)
		fmt.Printf("assist by %s\n", assist)
		switch position {
		case goalKeeper:
			points += 8
		case midFielder:
			points += 7
		case defender:
			points += 8
		case forward:
			points += 6
		}
	case "y-card":
		points += -1
	case "yr-card":
		points += -2
	case "r-card":
		points += -2
	case "soccer-ball-own":
		points += -3
	}

	if teamScore<1 {
		switch position {
		case goalKeeper:
			points += 6
		case defender:
			points += 3

		}
	}

	return points
}


func GetFilledTeamWithEvent(realTeam *helpers.VirtualTeam, event *MatchEvents) *helpers.VirtualTeam {

	team := new(helpers.VirtualTeam)
	*team = *realTeam

	for typ, tev := range event.Home.Events {
		if typ == principal {
			for _, e := range tev {
				for _, p := range team.Players.GoalKeeper {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, goalKeeper, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, midFielder, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, forward, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, defender, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}
			}
		} else {
			for _, e := range tev {
				for _, p := range team.Players.GoalKeeper {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, goalKeeper, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, midFielder, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, forward, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, defender, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}
			}
		}
	}

	for typ, tev := range event.Away.Events {
		if typ == principal {
			for _, e := range tev {
				for _, p := range team.Players.GoalKeeper {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, goalKeeper, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, midFielder, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, forward, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, defender, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}
			}
		} else {
			for _, e := range tev {
				for _, p := range team.Players.GoalKeeper {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, goalKeeper, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, midFielder, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, forward, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, defender, ie)
							totalPoints += partialPoints

							points := strconv.Itoa(totalPoints)
							p.Points = points
						}

					}
				}
			}
		}
	}
	return team
}

func FillTeamsUser(user *helpers.User, events []*MatchEvents) {

	for _, team := range user.PlayingTeams {
		for _, event := range events {
			filledTeam := GetFilledTeamWithEvent(team, event)
			pp.Println(filledTeam)
		}
		return

	}

}
