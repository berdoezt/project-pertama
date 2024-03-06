package lib

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=password dbname=hacktiv sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
