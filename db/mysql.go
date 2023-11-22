package db

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/vpovarna/go-todo-api/config"
)

// Can't this be injected as well? Do we need a struct?
var ctx = context.Background()

// TODO: Add close function ?!
func CreateMySQLConnection(envVar config.ToDoServiceConfig) *sqlx.DB {
	mysqlDB := envVar.MysqlDB

	dsm := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", mysqlDB.Username, mysqlDB.Password, mysqlDB.Url, mysqlDB.DBName)
	db, err := sqlx.ConnectContext(
		ctx,
		mysqlDB.Driver,
		dsm,
	)

	if err != nil {
		log.Fatal("Unable to connect to the database. Error: %s", err)
	}

	//TODO: Move this to healthCheck handler ?!
	err = db.Ping()
	if err != nil {
		log.Warn("DB is not ready. Error:", err)
	} else {
		log.Info("Successfully connected to DB:", mysqlDB.DBName)
	}

	return db
}
