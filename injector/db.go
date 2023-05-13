package injector

import (
	"connectQuerySQL/config"
	"connectQuerySQL/postgresql"
	"fmt"
	"log"

	"github.com/River-Island/product-gopkgs/logger"
	injPostgresql "github.com/River-Island/product-injector-api/pkg/persistence/postgresql"
)

type Service struct {
	Db  *injPostgresql.DbImpl
	Log logger.Logger
}

// ConnectToDB connects to the injector database
// It returns a service struct containing the db and tables, and also a closeFunc to safely close all connections when needed
func ConnectToDB(dbCfg config.DBConfig, sshCfg config.SSHConfig) (*Service, func(), error) {
	log.Print("connecting to DB")

	// open connection to DB
	db, closeFunc, err := postgresql.OpenSQLXViaSSH(dbCfg, sshCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db via ssh - err: %v", err)
	}

	log.Print("connected to DB successfully")

	injectorDBImp := &injPostgresql.DbImpl{DB: db}

	log := logger.NewLogger("injector-api")

	return &Service{
		Db:  injectorDBImp,
		Log: log,
	}, closeFunc, nil
}
