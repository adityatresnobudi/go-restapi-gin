package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeTable(db *sql.DB) error {
	q1 := `
		CREATE TABLE IF NOT EXISTS accounts (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		account_number VARCHAR(255) NOT NULL,
  		account_holder VARCHAR(255) not null,
  		balance FLOAT NOT NULL, 
		roles VARCHAR(50) NOT NULL CHECK (roles IN ('user', 'admin')),
  		created_at TIMESTAMPTZ DEFAULT NOW(),
  		updated_at TIMESTAMPTZ DEFAULT NOW()
	);`

	q2 := `
		CREATE TABLE IF NOT EXISTS transactions (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		account_id_from VARCHAR(255) NOT NULL,
  		account_id_to VARCHAR(255) not null,
  		amount FLOAT NOT NULL, 
  		created_at TIMESTAMPTZ DEFAULT NOW(),
  		updated_at TIMESTAMPTZ DEFAULT NOW()
	);`

	q3 := `
		CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY NOT NULL,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`

	if _, err := db.Exec(q1); err != nil {
		log.Printf("initialize table accounts: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q2); err != nil {
		log.Printf("initialize table transactions: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q3); err != nil {
		log.Printf("initialize table users: %s\n", err.Error())
		return err
	}

	return nil
}
