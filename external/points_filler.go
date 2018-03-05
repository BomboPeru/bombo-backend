package external

import (
	"../helpers"
	"regexp"
	"strings"
	"../util"
	"strconv"
	pretty "github.com/k0kubun/pp"
	"fmt"
	"io/ioutil"
	"github.com/gocarina/gocsv"
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
type PlayerPoints struct {
	Name string `json:"name" csv:"Name"`
	Team string `json:"team" csv:"Team"`
	Position string `json:"position" csv:"Position"`
	Cost float64 `json:"cost" csv:"Cost"`
	Points float64 `json:"points" csv:"Points"`
}

func GetPlayersPointsFromCSV(filePath string) ([]*PlayerPoints, error) {
	playersPoints := make([]*PlayerPoints, 0)

	data, err :=ioutil.ReadFile(filePath)
	if err != nil{
		return playersPoints, err
	}

	err = gocsv.UnmarshalBytes(data, &playersPoints)
	if err != nil {
		return playersPoints, err
	}

	return playersPoints, nil
}


func FillUserTeamWithPlayerPoints(user *helpers.User, pPoints []*PlayerPoints) {
	for _, eachPlayerTeam := range user.PlayingTeams {
		totalPoints := 0.0
		for _, p := range eachPlayerTeam.Players.Defender {
			for _, pp := range pPoints {
				if util.MatchNames(p.Name, pp.Name) {
					totalPoints += pp.Points
					strPoints := strconv.FormatFloat(pp.Points, 'f', -1, 64)
					p.Points = strPoints
					break
				}
			}
		}

		for _, p := range eachPlayerTeam.Players.MidFielder {
			for _, pp := range pPoints {
				if util.MatchNames(p.Name, pp.Name) {
					totalPoints += pp.Points
					strPoints := strconv.FormatFloat(pp.Points, 'f', -1, 64)
					p.Points = strPoints
					break
				}
			}
		}

		for _, p := range eachPlayerTeam.Players.GoalKeeper {
			for _, pp := range pPoints {
				if util.MatchNames(p.Name, pp.Name) {
					totalPoints += pp.Points
					strPoints := strconv.FormatFloat(pp.Points, 'f', -1, 64)
					p.Points = strPoints
					break
				}
			}
		}

		for _, p := range eachPlayerTeam.Players.Forward {
			for _, pp := range pPoints {
				if util.MatchNames(p.Name, pp.Name) {
					totalPoints += pp.Points
					strPoints := strconv.FormatFloat(pp.Points, 'f', -1, 64)
					p.Points = strPoints
					break
				}
			}
		}

		eachPlayerTeam.Points = totalPoints

	}
}

func FillPointsInPlayersList(league League, matchs []*MatchEvents) []*helpers.Player {
	//players := make([]*helpers.Player, 0)

	allLeaguePlayers := make([]*Player, 0)
	for _, team := range league {
		allLeaguePlayers = append(allLeaguePlayers, team.GoalKeeper...)
		allLeaguePlayers = append(allLeaguePlayers, team.MidFielder...)
		allLeaguePlayers = append(allLeaguePlayers, team.Defender...)
		allLeaguePlayers = append(allLeaguePlayers, team.Forwarder...)
	}
	return []*helpers.Player{}
}

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
		//fmt.Printf("assist by %s\n", assist)
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

func GetFilledTeamWithMatchEvent(realTeam *helpers.VirtualTeam, event *MatchEvents) *helpers.VirtualTeam {

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
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, goalKeeper, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, midFielder, ie)
							totalPoints += float64(partialPoints)


						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, forward, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, defender, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

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
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, goalKeeper, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, midFielder, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, forward, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Home.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, defender, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Away.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

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
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, goalKeeper, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Home.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, midFielder, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, forward, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, principal, defender, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

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
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, goalKeeper, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.MidFielder {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, midFielder, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points
					}
				}

				for _, p := range team.Players.Forward {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, forward, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}

				for _, p := range team.Players.Defender {
					if util.MatchNames(e.Name, p.Name) {
						score, _ := strconv.Atoi(event.Away.Score)
						if event.Away.Score == "-" {
							score = 0
						}
						totalPoints := 0.0
						for _, ie := range e.Events {

							partialPoints := GetPointsToPlayerByEvent(score, second, defender, ie)
							totalPoints += float64(partialPoints)

						}
						score2, _ := strconv.Atoi(event.Home.Score)
						if score > score2 {
							totalPoints *= 1.1
						}
						if score < score2 {
							totalPoints *= 0.9
						}
						points := strconv.FormatFloat(totalPoints, 'f', -1, 64)
						p.Points = points

					}
				}
			}
		}
	}
	totalTeamPoints := 0.0

	for _, p := range team.Players.Forward {
		points, _ := strconv.ParseFloat(p.Points, 64)
		totalTeamPoints += float64(points)
	}

	for _, p := range team.Players.Defender {
		points, _ := strconv.ParseFloat(p.Points, 64)
		totalTeamPoints += float64(points)
	}

	for _, p := range team.Players.MidFielder {
		points, _ := strconv.ParseFloat(p.Points, 64)
		totalTeamPoints += float64(points)
	}

	for _, p := range team.Players.GoalKeeper{
		points, _ := strconv.ParseFloat(p.Points, 64)
		totalTeamPoints += float64(points)
	}

	team.Points = totalTeamPoints

	return team
}

func FillTeamsUser(user *helpers.User, events []*MatchEvents) {

	for _, team := range user.PlayingTeams {
		for _, event := range events {
			filledTeam := GetFilledTeamWithMatchEvent(team, event)

			pretty.Println(filledTeam.Points, fmt.Sprintf("%s vs %s", event.Home.Name, event.Away.Name))
		}
		return

	}

}
