package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type databaseFacotry struct {
	db *sql.DB
}

var instance = databaseFacotry{}

func getInstance() *databaseFacotry {
	return &instance
}

func (f *databaseFacotry) initDB() error {
	dbHost := "127.0.0.1"
	dbUser := "postgres"
	dbPassowrd := "postgres"
	dbDatabase := "postgres"
	dbConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassowrd, dbDatabase)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	f.db = db
	return nil
}

func GetDB() (*sql.DB, error) {
	f := getInstance()

	if f.db != nil {
		return f.db, nil
	}

	if err := f.initDB(); err != nil {
		return nil, err
	}

	return f.db, nil
}

func GetTx() (*sql.Tx, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CloseDB() error {
	f := getInstance()
	if f.db == nil {
		return nil
	}

	return f.db.Close()
}
