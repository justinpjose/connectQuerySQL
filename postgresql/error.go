package postgresql

import (
	"database/sql"
	"fmt"
	"log"
)

func TxBeginError(e error) error {
	return fmt.Errorf("could not begin transaction: %w", e)
}

func Rollback(tx *sql.Tx, originalErr error) error {
	rollErr := tx.Rollback()
	if rollErr != nil {
		return rollbackError(originalErr, rollErr)
	}

	return nil
}

func rollbackError(originalErr, rollErr error) error {
	return fmt.Errorf("could not rollback transaction: %w. original error: %v", rollErr, originalErr)
}

func Commit(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Fatal(commitError(err))
	}
}

func commitError(e error) error {
	return fmt.Errorf("could not commit transaction: %w", e)
}
