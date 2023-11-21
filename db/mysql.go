package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/vpovarna/go-todo-api/config"
)

func CreateMySQLConnection(envVar config.ToDoServiceConfig) *sqlx.DB {
	mysqlDB := envVar.MysqlDB

	dsm := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", mysqlDB.Username, mysqlDB.Password, mysqlDB.Url, mysqlDB.DBName)
	db, err := sqlx.Connect(
		mysqlDB.Driver,
		dsm,
	)

	if err != nil {
		log.Fatal("Unable to db to the database. Error: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Warn("DB is not ready. Error:", err)
	} else {
		log.Info("Successfully connected to DB:", mysqlDB.DBName)
	}

	return db
}
