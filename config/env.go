package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ToDoServiceConfig struct {
	MysqlDB Database
}

type Database struct {
	Driver   string
	Url      string
	Username string
	Password string
	DBName   string
}

func LoadEnv() ToDoServiceConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load the config file as env variables")
	}

	mysqlUsername := os.Getenv("DB_USER")
	mysqlPassword := os.Getenv("DB_PASSWORD")
	mysqlHost := os.Getenv("DB_URL")
	mysqlDBName := os.Getenv("DB_NAME")

	return ToDoServiceConfig{
		MysqlDB: Database{
			Driver:   "mysql",
			Url:      mysqlHost,
			Username: mysqlUsername,
			Password: mysqlPassword,
			DBName:   mysqlDBName,
		},
	}
}
