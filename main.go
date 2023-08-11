package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	log.Println("Init in main package")
}

func insertWallet(db *sql.DB, owner string, balance float64) error {
	query := `
			INSERT INTO wallet (owner, balance)
			VALUES ($1, $2) RETURNING id
	`
	row := db.QueryRow(query, owner, balance)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	log.Println("last insert id: ", id)
	return nil
}

func updateWalletBalance(db *sql.DB, walletID int, newBalance float64) error {
	query := `
			UPDATE wallet SET balance = $1 WHERE id = $2
	`
	_, err := db.Exec(query, newBalance, walletID)
	if err != nil {
		return err
	}

	return nil
}

func deleteWallet(db *sql.DB, walletID int) error {
	query := `
			DELETE FROM wallet WHERE id = $1
	`
	_, err := db.Exec(query, walletID)
	if err != nil {
		return err
	}

	return nil
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
	query := `CREATE TABLE IF NOT EXISTS wallet (
		id SERIAL PRIMARY KEY,
		owner VARCHAR(255) NOT NULL,
		balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
	)`

	_, err = conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	err = insertWallet(conn, "John Doe", 100.00)
	if err != nil {
		log.Fatal("insert error:", err)
	}

	err = updateWalletBalance(conn, 1, 200.00)
	if err != nil {
		log.Fatal("update error:", err)
	}

	err = deleteWallet(conn, 1)
	if err != nil {
		log.Fatal("delete error:", err)
	}

	log.Println("done")
}
