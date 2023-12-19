package todo

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type TodoDBKey struct{}

func WithTodoDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, &TodoDBKey{}, db)
}

func ExtractTodoDB(ctx context.Context) (*gorm.DB, error) {
	if v := ctx.Value(&TodoDBKey{}); v != nil {
		db, ok := v.(*gorm.DB)
		if ok {
			return db, nil
		}
	}
	return nil, errors.New("ExtractTodoDB: failed to extract DB")
}
