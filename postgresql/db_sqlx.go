package postgresql

import (
	"connectQuerySQL/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint: golint
)

func OpenSQLXWithoutSSH(dbCfg config.DBConfig) (*sqlx.DB, func(), error) {
	db, err := openSQLX(dbCfg, false)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database - err: %v", err)
	}

	close := func() {
		db.Close()
	}

	return db, close, nil
}

// openSQLX - opens a PostgreSQL database and returns sqlX db
func openSQLX(dbCfg config.DBConfig, ssh bool) (*sqlx.DB, error) {
	dbInfo := fmt.Sprintf(
		"sslmode=disable host=%s port=%d user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.Username, dbCfg.Password, dbCfg.Name)

	driverName := postgresDriverName
	if ssh {
		driverName = postgresSSHDriverName
	}

	db, err := sqlx.Connect(driverName, dbInfo)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
