package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port string `envconfig:"PORT" default:"8080" required:"true" split_words:"true"`
}

type DBConfig struct {
	Username string `envconfig:"DB_USERNAME" default:"root" required:"true" split_words:"true"`
	Password string `envconfig:"DB_PASSWORD" default:"Xonen@3616" required:"true" split_words:"true"`
	Host     string `envconfig:"DB_HOST" default:"127.0.0.1" required:"true" split_words:"true"`
	Port     string `envconfig:"DB_PORT" default:"3306" required:"true" split_words:"true"`
	Name     string `envconfig:"DB_NAME" default:"mysql_database" required:"true" split_words:"true"`
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

func LoadAppConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	c := &Config{}
	if err := envconfig.Process("", c); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	return c
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
}

func (c *Config) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", c.ConnectionString())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to the database")

	return db, nil
}
