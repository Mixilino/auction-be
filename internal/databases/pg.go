package databases

import (
	"auction-be/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connStr, err := getConnectionString()
	if err != nil {
		panic("Error opening env file" + err.Error())
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Error opening database" + err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic("Error when pinging database")
	}

	return db
}

func getConnectionString() (string, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s", config.GetDBUser(),
		config.GetDBName(), config.GetDBPassword(), config.GetDBHost(), config.GetDBPort(), config.GetSSLMode())
	return connStr, nil
}
