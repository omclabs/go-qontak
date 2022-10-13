package configs

import (
	"database/sql"
	"fmt"
	"omclabs/go-qontak/app/helpers"
	"os"
	"time"
)

func NewMysqlConn() *sql.DB {
	driver := os.Getenv("MYSQL_DRIVER")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	databaseName := os.Getenv("MYSQL_DATABASE")

	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, username, password, host, port, databaseName)
	db, err := sql.Open(driver, connectionString)
	helpers.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
