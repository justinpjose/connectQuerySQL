package backbonev2

import (
	"connectQuerySQL/config"
	"connectQuerySQL/postgresql"
	"database/sql"
	"fmt"
	"log"
)

type Service struct {
	Db     *sql.DB
	Tables dbTables
}

// ConnectToDB connects to the backboneV2 database
// It returns a service struct containing the db and tables, and also a closeFunc to safely close all connections when needed
func ConnectToDB(dbCfg config.DBConfig, sshCfg config.SSHConfig) (*Service, func(), error) {
	log.Print("connecting to DB")

	// open connection to DB
	db, closeFunc, err := postgresql.OpenSQLViaSSH(dbCfg, sshCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db via ssh - err: %v", err)
	}

	// get tables
	backboneV2Tables := getDbTables()

	log.Print("connected to DB successfully")

	return &Service{
		Db:     db,
		Tables: backboneV2Tables,
	}, closeFunc, nil
}
