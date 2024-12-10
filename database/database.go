package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type postgresConfig struct {
	user     string
	password string
	host     string
	port     string
	dbName   string
	sslmode  string
}

func ConnectPostgres() error {

	var err error
	// connect postgres URL
	var conf = postgresConfig{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbName:   os.Getenv("DB_NAME"),
		sslmode:  os.Getenv("DB_MODE"),
	}

	// set postgres URL
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", conf.user, conf.password, conf.host, conf.port, conf.dbName, conf.sslmode)

	DB, err = sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("can not connect to database;", err)
	}

	// ping postgres DB
	if err = DB.Ping(); err != nil {
		log.Fatal("ping: ", err)
		return err
	}

	fmt.Println("Successfully connected to postgres DB")
	return nil
}
