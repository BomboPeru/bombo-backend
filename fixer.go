package main

import (
	"os"
	"log"
)

func init() {
	// Try to create ./database folder
	err := os.Mkdir("./database", 0755)
	if os.IsExist(err) {
		log.Println("Folder for database already exist")
	}
}