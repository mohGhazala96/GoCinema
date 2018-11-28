package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Configuration struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

var config Configuration

type User struct {
	id       int
	username string
}

var db *sql.DB

func myHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "Welcome to the server\n")
	fmt.Fprintln(w, "ID | Name")
	fmt.Fprintln(w, "---+--------")
	for rows.Next() {
		var (
			id       int
			username string
		)

		rows.Scan(&id, &username)

		fmt.Fprintf(w, "%2d | %s\n", id, username)
	}
}

func main() {
	config := Configuration{}
	config.Dbname = os.Getenv("DATABASE_NAME")
	config.Host = os.Getenv("DATABASE_HOST")
	config.User = os.Getenv("DATABASE_USER")
	config.Password = os.Getenv("DATABASE_PASSWORD")
	config.Port = os.Getenv("DATABASE_PORT")
	var webPort = os.Getenv("WEB_PORT")

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// rows, err := db.Query(`SELECT id, username FROM users`)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	user := User{}
	// 	err = rows.Scan(&user.id, &user.username)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(user.username)
	// 	fmt.Println(user.id)

	// }

	http.HandleFunc("/", myHandler)
	//http.HandleFunc("/cache", myCachedHandler)
	log.Print("Listening on " + ":" + webPort + "...")
	http.ListenAndServe(":"+webPort, nil)

}
