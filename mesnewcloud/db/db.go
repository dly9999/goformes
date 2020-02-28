package Db

import (
	"fmt"
	"goformescloud/mesnewcloud/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Dbc struct {
	Sql *sqlx.DB
}

var db Dbc

func init() {
	dbconfig, _, _, _ := config.Getconfig()
	costr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbconfig.Usr, dbconfig.Psd, dbconfig.Server, dbconfig.Port, dbconfig.Dbname)
	fmt.Println(costr)
	db1, err := sqlx.Open("mysql", costr) //"root:root@tcp(127.0.0.1:3306)/test"
	if err != nil {
		fmt.Println(err)
	}

	db.Sql = db1
}
func ReDb() Dbc {
	return db
}
func CloseDb() {
	defer db.Sql.Close()
}
