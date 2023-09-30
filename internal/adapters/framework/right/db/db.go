package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	DB *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("DB Connection Failure : %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB Connection Failure : %v", err)
	}

	return &Adapter{DB: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("DB Close Failure")
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	qs, args, err := squirrel.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.DB.Exec(qs, args...)
	if err != nil {
		return err
	}

	return nil
}
