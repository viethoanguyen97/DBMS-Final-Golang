package queryMySQL

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:01120309@tcp(localhost:3306)/SampleOrderRDBMSData?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect database succesfully")
}

func CloseDB() {
	db.Close()
}

