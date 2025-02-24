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

	if _, err := db.Exec(q1); err != nil {
		log.Printf("initialize table error: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q2); err != nil {
		log.Printf("initialize table error: %s\n", err.Error())
		return err
	}

	return nil
}
