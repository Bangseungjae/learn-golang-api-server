package code3_2

import (
	"context"
	"database/sql"
)

func (r *Repository) Update(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.Exec("/* 변경 처리 1 */")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("/* 변경 처리 2 */")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("/* 변경 처리 3 */")
	if err != nil {
		tx.Rollback()
		return err
	}

	// 다른 처리
	return tx.Commit()
}

type Repository struct {
	db *sql.DB
}
