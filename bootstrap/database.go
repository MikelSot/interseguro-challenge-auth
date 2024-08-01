package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	onceDB sync.Once
	db     *sql.DB
	err    error
)

func newDataBase() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	onceDB.Do(func() {
		dns := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			user,
			password,
			host,
			port,
			dbname,
			sslMode,
		)

		db, err = sql.Open("postgres", dns)
	})
	if err != nil {
		return err
	}

	return nil
}

func getConnectionDB() *sql.DB {
	if err := newDataBase(); err != nil {
		log.Fatalf("no se pudo obtener una conexion a la base de datos de extraccion: %v", err)
	}

	return db
}
