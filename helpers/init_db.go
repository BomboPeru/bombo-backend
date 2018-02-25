package helpers

import (
	"log"

	"github.com/asdine/storm"
)

// DbUser ...
var DbUser *storm.DB

// DbPlayer ...
var DbPlayer *storm.DB

// DbVirtualTeam ...
var DbVirtualTeam *storm.DB

func init() {
	var err error

	DbUser, err = storm.Open("./database/user.db")
	if err != nil {
		log.Panicln("storm.Open() at User, ", err.Error())
	}

	DbPlayer, err = storm.Open("./database/player.db")
	if err != nil {
		log.Panicln("storm.Open() at Player, ", err.Error())
	}

	DbVirtualTeam, err = storm.Open("./database/virtualteam.db")
	if err != nil {
		log.Panicln("storm.Open() at VirtualTeam, ", err.Error())
	}

}
