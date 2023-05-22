package models

import (
	"database/sql"
	"fmt"
)

type DatabaseModel struct {
	DB *sql.DB
}

type Bucket struct {
	Id         int    `json:"id"`
	BucketName string `json:"bucket_name"`
}

// StoreBucket stores a new bucket in the database
func (m *DatabaseModel) StoreBucket(bucketName string) error {
	stmt := `INSERT INTO buckets(bucket_name) VALUES ($1)`
	_, err := m.DB.Exec(stmt, bucketName)
	if err != nil {
		return err
	}

	return nil
}

func (m *DatabaseModel) StoreFiles(bucketName string, fileName string) error {
	stmt := `INSERT INTO files (bucket_name, file_name) VALUES ($1, $2)`
	_, err := m.DB.Exec(stmt, bucketName, fileName)
	if err != nil {
		return err
	}

	return nil
}

func (m *DatabaseModel) CheckIfFileExists(bucketName, fileName string) error {
	stmt := "SELECT file_name FROM files WHERE bucket_name = $1 AND file_name = $2"
	res, err := m.DB.Exec(stmt, bucketName, fileName)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows > 0 {
		return fmt.Errorf("file already exists")
	}

	return nil
}

// VerifyBucketName function verifies if a bucket name already exists in the database
// returns an error if the bucket name already exists
func (m *DatabaseModel) VerifyBucketName(bucketName string) error {
	stmt := `SELECT bucket_name FROM buckets WHERE bucket_name = $1`
	res, err := m.DB.Exec(stmt, bucketName)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()

	if rows > 0 {
		return fmt.Errorf("bucket name already exists")
	}

	return nil
}
