package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	_ "github.com/lib/pq" // here
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	fmt.Fprintf(os.Stdout, "Starting...")
	fmt.Fprintf(os.Stdout, psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	var NODE_NAME string
	if NODE_NAME = os.Getenv("NODE_NAME"); NODE_NAME == "" {
		NODE_NAME = "?"
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := `
			INSERT INTO request_history (time)
			VALUES (NOW())
			RETURNING id
		`
		row_id := 0
		err = db.QueryRow(sqlStatement).Scan(&row_id)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Hello World. I am node: %s\n", NODE_NAME)
	})

	http.ListenAndServe(":"+PORT, nil)
}
