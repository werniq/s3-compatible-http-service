package driver

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// OpenDb function opens a connection to the database
// database connection string is stored in the .env file
// returns a pointer to the database and an error if the connection could not be established

var (
	DATABASE_URL = "host=localhost port=5432 dbname=s3 user=postgres password=password sslmode=disable"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	stmt := "CREATE TABLE IF NOT EXISTS buckets (id SERIAL PRIMARY KEY, bucket_name VARCHAR(255) NOT NULL)"
	_, err = db.Exec(stmt)
	if err != nil {
		return nil, err
	}

	stmt = "CREATE TABLE IF NOT EXISTS files (id SERIAL PRIMARY KEY, bucket_name VARCHAR(255) NOT NULL, file_name VARCHAR(255) NOT NULL)"
	_, err = db.Exec(stmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}
