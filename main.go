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

var Config Configuration
var conf map[string]string

type User struct {
	id       int
	username string
}

var db *sql.DB

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(Config)        // output: [UserA, UserB]
	fmt.Println(Config.Dbname) // output: [UserA, UserB]
	conf := make(map[string]string)
	conf["DBHOST"] = Config.Host
	conf["DBPORT"] = Config.Port
	conf["DBUSER"] = Config.User
	conf["DBPASS"] = Config.Password
	conf["DBNAME"] = Config.Dbname

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf["DBHOST"], conf["DBPORT"],
		conf["DBUSER"], conf["DBPASS"], conf["DBNAME"])

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
