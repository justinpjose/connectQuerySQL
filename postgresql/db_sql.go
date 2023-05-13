package postgresql

import (
	"connectQuerySQL/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // nolint: golint
)

func OpenSQLWithoutSSH(dbCfg config.DBConfig) (*sql.DB, func(), error) {
	db, err := openSQL(dbCfg, false)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database - err: %v", err)
	}

	close := func() {
		db.Close()
	}

	return db, close, nil
}

// openSQL - opens a PostgreSQL database and returns SQL db
func openSQL(dbCfg config.DBConfig, ssh bool) (*sql.DB, error) {
	dbInfo := fmt.Sprintf(
		"sslmode=disable host=%s port=%d user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.Username, dbCfg.Password, dbCfg.Name)

	driverName := postgresDriverName
	if ssh {
		driverName = postgresSSHDriverName
	}

	db, err := sql.Open(driverName, dbInfo)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
