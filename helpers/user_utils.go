package helpers

import "github.com/satori/go.uuid"

func GetIndexOfTeamById(teams []*VirtualTeam, id uuid.UUID) int {
	for i, team := range teams {
		if uuid.Equal(team.ID, id) {
			return i
		}
	}
	return -1
}