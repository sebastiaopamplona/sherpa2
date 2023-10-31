package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func commitTxAndRollbackOnError(tx *sqlx.Tx, err error) error {
	if tx == nil {
		return fmt.Errorf("nil tx")
	}

	var errOut error
	if err == nil {
		if commitErr := tx.Commit(); commitErr != nil {
			errOut = commitErr
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				errOut = fmt.Errorf("%v; %v", errOut, rollbackErr)
			}
		}
	}

	if err != nil {
		errOut = err
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			errOut = fmt.Errorf("%v; %v", errOut, rollbackErr)
		}
	}

	return errOut
}
