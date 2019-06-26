package dbops


import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	//root:123456@tcp(localhost:3306)/video?charset=utf8 这一段线上可以不用这么多
	dbConn, err = sql.Open("mysql", "root:root@tcp(100.1.1.9:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
