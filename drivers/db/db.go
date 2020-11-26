package db

import (
	"database/sql"
	"fmt"
	"time"

	"lemontech.com/metaq/domain"
	"lemontech.com/metaq/drivers/store"
)

var db *sql.DB

// Connect will create the db connection
func Connect() (err error) {
	db, err = sql.Open("mysql", store.ENV.DBURL)
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	return
}

// CheckURL will try to connect
func CheckURL(url string) (err error) {
	db, err = sql.Open("mysql", url)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Ping()
	return
}

// ShowDBs returns a list of available dbs
func ShowDBs() (dbs domain.Databases, err error) {
	statement := "SHOW DATABASES"
	rows, err := db.Query(statement)
	defer rows.Close()

	if err != nil {
		return
	}

	_, result := parseArr("databases", rows)
	for i := range result {
		dbs = append(dbs, domain.Database{Name: result[i][1], Selected: false})
	}

	return
}

// Query will execute the provided statement and return a dataset
func Query(dbs domain.Databases, statement string) (data domain.Datasets, err error) {
	for i := range dbs {
		if dbs[i].Selected {
			prestatement := fmt.Sprintf("use %s", dbs[i].Name)
			tx, _ := db.Begin()
			r, _ := tx.Query(prestatement)
			defer r.Close()
			rows, errq := tx.Query(statement)
			defer tx.Rollback()

			if errq != nil {
				err = errq
				return
			}
			defer rows.Close()

			keys, result := parseArr(dbs[i].Name, rows)
			data.Headers = keys
			data.Rows = append(data.Rows, result...)
		}
	}
	return
}

func parseArr(name string, rows *sql.Rows) (keys []string, result [][]string) {
	keys, _ = rows.Columns()
	for rows.Next() {
		cols := make([]interface{}, len(keys))
		for i := range keys {
			cols[i] = &cols[i]
		}
		rows.Scan(cols...)
		res := []string{name}
		for i := range keys {
			res = append(res, string(cols[i].([]byte)))
		}
		result = append(result, res)
	}
	keys = append([]string{"database"}, keys...)
	return
}
