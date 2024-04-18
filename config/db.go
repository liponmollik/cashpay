package db

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db              *sql.DB
	connectionError error
	once            sync.Once
)

// InitializeDB initializes the database connection pool.
func InitializeDB(dataSourceName string) error {
	once.Do(func() {
		db, connectionError = sql.Open("mysql", dataSourceName)
	})
	return connectionError
}

// GetDB returns a database connection from the pool.
func GetDB() *sql.DB {
	return db
}
