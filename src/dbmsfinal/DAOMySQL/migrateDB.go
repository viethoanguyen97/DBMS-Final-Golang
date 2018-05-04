package queryMySQL

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:01120309@tcp(localhost:3306)/DBMSFinal?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	err = DB.Ping()
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect database succesfully")
}

func CloseDB() {
	DB.Close()
}
