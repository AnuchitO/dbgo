package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	log.Println("Init in main package")
}

func main() {
	conn, err := sql.Open("postgres", "postgres://batsjuib:Voopy92tnjyUMBQhi0EDtxTdg--aA-rK@tiny.db.elephantsql.com/batsjuib")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create Wallet table into database
	sql := `CREATE TABLE IF NOT EXISTS wallet (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
	)`

	_, err = conn.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}
