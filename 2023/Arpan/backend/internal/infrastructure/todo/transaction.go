package todo

import (
	"context"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
)

type transactor struct {
	conn *TodoConn
}

func NewTransactor(conn *TodoConn) gateway.Transactor {
	return &transactor{conn: conn}
}

func (t transactor) Transaction(ctx context.Context, worker func(context.Context) error) error {
	tx := t.conn.GormDB.Begin()

	newCtx := WithTodoDB(ctx, tx)

	if err := worker(newCtx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
