package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vpovarna/go-todo-api/config"
	"log"
)

func CreateMySQLConnection(envVar config.ToDoServiceConfig) *sqlx.DB {
	mysqlDB := envVar.MysqlDB

	db, err := sqlx.Connect(
		mysqlDB.Driver,
		fmt.Sprintf("%s:%s@(%s)/%s", mysqlDB.Username, mysqlDB.Password, mysqlDB.Url, mysqlDB.DBName),
	)

	if err != nil {
		log.Fatalf("Unable to db to the database. Error: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB is not ready. Error: %s", err)
	} else {
		log.Printf("Successfully connected to DB: %s", mysqlDB.DBName)
	}

	return db
}
