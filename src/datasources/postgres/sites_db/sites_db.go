package sites_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB

	username = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = os.Getenv("POSTGRES_HOST")
	db       = os.Getenv("POSTGRES_DB")
)

func init() {

	datasourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		host,
		db,
	)

	var err error
	Client, err = sql.Open("postgres", datasourceName)

	if err != nil {
		panic(err)
	}

	// if err = Client.Ping(); err != nil {
	// 	panic(err)
	// }

	log.Println("database successfully configured")
}
