package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Db connection was failed: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Db ping was failed: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (adapter *Adapter) CloseDbConnection() {
	err := adapter.db.Close()
	if err != nil {
		log.Fatalf("Could not close connection to the database due to an error: %v", err)
	}
}

func (adapter *Adapter) AddToHistory(result int, operation string) error {
	queryString, args, err := sq.Insert("OperationsHistory").Columns("Result", "Operation", "CreatedAt").Values(result, operation, time.Now().Local()).ToSql()
	if err != nil {
		return err
	}

	_, err = adapter.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
