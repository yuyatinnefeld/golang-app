package db

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "simple_bank"
)

func TestMain(m *testing.M) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println(conn)
	defer conn.Close()
	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
