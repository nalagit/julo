package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const SVR = "root:@tcp(127.0.0.1:3306)/julo?parseTime=true"

func ExecuteSQL(query string, params interface{}) *sql.Rows {
	db, err := sql.Open("mysql", SVR)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var result *sql.Rows
	if params == "" {
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}
	} else {
		result, err = db.Query(query, params)
		if err != nil {
			panic(err.Error())
		}
	}

	return result
}

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", SVR)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
