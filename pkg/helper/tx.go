package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx, err error) {
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Println("SQL error on Helper TX => Error Rollback", errRollback)
			return
		}
	} else {
		tx.Commit()
	}
}
