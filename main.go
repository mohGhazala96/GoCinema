package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	databaseConfig := make(map[string]string)
	databaseConfig["DBHOST"] = config.Host
	databaseConfig["DBPORT"] = config.Port
	databaseConfig["DBUSER"] = config.User
	databaseConfig["DBPASS"] = config.Password
	databaseConfig["DBNAME"] = config.Dbname

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseConfig["DBHOST"], databaseConfig["DBPORT"],
		databaseConfig["DBUSER"], databaseConfig["DBPASS"], databaseConfig["DBNAME"])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT id, username FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.id, &user.username)
		if err != nil {
			panic(err)
		}
		fmt.Println(user.username)
		fmt.Println(user.id)

	}

}
