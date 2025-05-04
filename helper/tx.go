package helper

import (
	"database/sql"
	"log"
)

func NewTx(tx *sql.Tx, err *error) error {
	if *err != nil {
		tx.Rollback()
		log.Print(err)
		return *err
	} else {
		tx.Commit()
		return nil
	}
}
