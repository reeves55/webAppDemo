package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var(
	username string = ""
	password string = ""
	dbname = ""
	db, err = sql.Open("mysql", "user:password@/dbname")
)

func init()  {

}