package main

import (
	"connectQuerySQL/backbonev2"
	"connectQuerySQL/config"
	"log"
)

func main() {
	dbCfg, sshCfg, err := config.Get()
	if err != nil {
		log.Fatalf("failed to get configurations - err: %v", err)
	}
	defer config.Unset()

	// Change receiver of the method accordingly to the relevant service
	service, closeConn, err := backbonev2.ConnectToDB(*dbCfg, *sshCfg)
	if err != nil {
		log.Fatalf("failed to connect to db - err: %v", err)
	}
	defer closeConn()

	log.Print("Running the service")
	run(service)
}
