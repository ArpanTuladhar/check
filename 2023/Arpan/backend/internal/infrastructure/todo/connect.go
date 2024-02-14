package todo

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoConn struct {
	GormDB *gorm.DB
}

func NewOwnerSQLHandler(conf *config.Config) (*TodoConn, func(), error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s",
		conf.DBUser,
		conf.DBPass,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
		url.QueryEscape(conf.DBLoc),
	)

	db, closer, err := newSQLHandler(dsn, conf)
	return &TodoConn{GormDB: db}, closer, err
}

func newSQLHandler(dsn string, conf *config.Config) (*gorm.DB, func(), error) {
	if _, err := time.LoadLocation("UTC"); err != nil {
		return nil, nil, fmt.Errorf(": %w", err)
	}

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}))
	if err != nil {
		return nil, nil, fmt.Errorf(": %w", err)
	}

	db, err := conn.DB()
	if err != nil {
		return nil, nil, fmt.Errorf(": %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, nil, fmt.Errorf(": %w", err)
	}

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn, func() { db.Close() }, nil
}

type binder struct {
	todoConn *TodoConn
}

func NewConnectionBinder(todoConn *TodoConn) gateway.Binder {
	return &binder{
		todoConn: todoConn,
	}
}

func (b binder) Bind(ctx context.Context) context.Context {
	return WithTodoDB(ctx, b.todoConn.GormDB)
}
