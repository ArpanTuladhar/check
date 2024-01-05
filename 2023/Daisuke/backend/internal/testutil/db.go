package testutil

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func InitDB(t *testing.T) (gormDB *gorm.DB, dbName string) {
	env := LoadEnv()
	dbName, dbCloseFunc, err := createCleanDB(env)
	if err != nil {
		t.Fatal(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		dbName,
		url.QueryEscape("Asia/Tokyo"),
	)

	gormDB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := dbCloseFunc(); err != nil {
			t.Fatal(err)
		}
		gormSQLDB, err := gormDB.DB()
		if err != nil {
			t.Fatal(err)
		}

		if err := gormSQLDB.Close(); err != nil {
			t.Fatal(err)
		}
	})

	return gormDB, dbName
}

func createCleanDB(e *TestEnv) (dbName string, closeFunc func() error, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", e.DBUser, e.DBPass, e.DBHost, e.DBPort)
	dbName = fmt.Sprintf("%s_%s", e.DBName, uuid.NewString())
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	createDB(sqlDB, dbName)
	createTables(sqlDB)
	//loadFixtures(sqlDB)
	closeFunc = func() error {
		err = dropDatabase(sqlDB, dbName)
		if err != nil {
			return fmt.Errorf("drop error %w", err)
		}

		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("sqlDB.Close() error %w", err)
		}

		return nil
	}

	return dbName, closeFunc, nil
}

func createDB(sqlDB *sql.DB, dbName string) {
	createFmt := fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbName)
	_, err := sqlDB.Exec(createFmt)
	if err != nil {
		panic(" sqlDB.Exec:" + err.Error())
	}

	if _, err := sqlDB.Exec(fmt.Sprintf("USE `%s`;", dbName)); err != nil {
		panic(err)
	}
}

func createTables(sqlDB *sql.DB) {
	_, thisFilePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller error")
	}

	dumpPath := filepath.Join(filepath.Dir(thisFilePath), "..", "..", "..")
	dumpPath = filepath.Join(dumpPath, "backend", "database", "docker", "sqls", "import", "create_tables.sql")

	dumpBytes, err := os.ReadFile(dumpPath)
	if err != nil {
		panic(err)
	}

	regexpNewline := regexp.MustCompile(`\r\n|\r|\n`)
	dumpStr := regexpNewline.ReplaceAllString(string(dumpBytes), "")

	for _, stmt := range strings.Split(dumpStr, ";") {
		if stmt == "" {
			continue
		}
		if _, err := sqlDB.Exec(stmt); err != nil {
			panic(err)
		}
	}
}

func dropDatabase(sqlDB *sql.DB, dbName string) error {
	_, err := sqlDB.Exec(fmt.Sprintf("DROP DATABASE `%s`;", dbName))
	if err != nil {
		return fmt.Errorf("drop error %w", err)
	}
	return nil
}
