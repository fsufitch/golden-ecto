package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
//	_ "github.com/lib/pq"
//	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/sqlite3"
//	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
//	_ "gopkg.in/doug-martin/goqu.v3/adapters/mysql"
)


var Dialect string
var DbConn *sql.DB
var DbQ *goqu.Database

func VerifyInit() {
	if DbConn == nil {
		panic("DB not initialized!")
	}
	err := DbConn.Ping()
	if err != nil {
		panic(err)
	}
}

func Initialize(driver string, source string) {
	conn, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}
	Dialect = driver
	DbConn = conn
	DbQ = goqu.New(driver, DbConn)
}
