package integration

import (
	"context"
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/testutil"
	"gorm.io/gorm"
)

type TodoDBTestHelper struct {
	ctx    context.Context
	gormDB *gorm.DB
	dbName string
}

func InitTodoDBTestHelper(t *testing.T) *TodoDBTestHelper {
	gormDB, dbName := testutil.InitDB(t)

	ctx := context.Background()

	return &TodoDBTestHelper{ctx: ctx, gormDB: gormDB, dbName: dbName}
}
