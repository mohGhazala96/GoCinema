package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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

//structs for movies
type Movies struct {
	Vote_count        int64
	Id                int64
	Video             bool
	Vote_average      float64
	Title             string
	Popularity        float64
	Poster_path       string
	Original_language string
	Original_title    string
	Genre_ids         []int64
	Backdrop_path     string
	Adult             bool
	Overview          string
	Release_date      string
}
type Dates struct {
	Maximum string
	Minimum string
}
type moviesReponse struct {
	Results       []Movies
	Page          int64
	Total_results int64
	Dates         Dates
	Total_pages   int64
}

//
var db *sql.DB

func myHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title,release_date,poster_path FROM movies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintln(w, "ID | Title")
	fmt.Fprintln(w, "---+--------")
	for rows.Next() {
		var (
			id           int
			title        string
			release_date string
			poster_path  string
		)

		rows.Scan(&id, &title, &release_date, &poster_path)
		fmt.Fprintf(w, "%5d | %s | %s | %s  \n", id, title, release_date, poster_path)
	}
	rows, err = db.Query("SELECT id, seats,movie FROM halls")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintln(w, "ID | seats")
	fmt.Fprintln(w, "---+--------")
	for rows.Next() {
		var (
			id    int64
			seats int
			movie int64
		)

		rows.Scan(&id, &seats, &movie)
		fmt.Fprintf(w, "%3d | %d | %d  \n", id, seats, movie)
	}
}
func getJson(url string, target interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	payload := strings.NewReader("{}")

	req, err := http.NewRequest("GET", url, payload)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
func insertMovies(movies *moviesReponse) {
	for movieIndex := range movies.Results {
		var sqlStatement string
		sqlStatement = "INSERT INTO movies (id,title, release_date, poster_path, vote_average,isAvialabe) (select $1 as id, $2 as title ,$3 as release_date,$4 as poster_path,$5 as vote_average ,$6 as isAvialabe where not exists (select * from movies where id=$1))"
		var err error
		_, err = db.Exec(sqlStatement, movies.Results[movieIndex].Id, movies.Results[movieIndex].Title,
			movies.Results[movieIndex].Release_date,
			"http://image.tmdb.org/t/p/w500/"+movies.Results[movieIndex].Poster_path,
			movies.Results[movieIndex].Vote_average, true)
		if err != nil {
			panic(err)
		}
	}
}

func removeUnavialabeMoviesFromCinema(movies *moviesReponse) {
	rows, err := db.Query("SELECT id FROM movies")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id int64
		)
		rows.Scan(&id)
		var isFound bool
		for movieIndex := range movies.Results {
			if movies.Results[movieIndex].Id == id {
				isFound = true
			}
		}
		if !isFound {
			//update here
			var sqlStatement string
			sqlStatement = "UPDATE movies SET isAvialabe = $2 WHERE id = $1;"
			var err error
			_, err = db.Exec(sqlStatement, id, false)
			if err != nil {
				panic(err)
			}
		}
	}
}

func updateHalls(movies *moviesReponse) {
	rows, err := db.Query("SELECT id,movie FROM halls")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var movieIndex int = 0
	for rows.Next() {
		var (
			id int64
		)
		rows.Scan(&id)

		var sqlStatement string
		sqlStatement = "UPDATE halls SET movie = $2 WHERE id = $1;"
		var err error
		_, err = db.Exec(sqlStatement, id, movies.Results[movieIndex].Id)
		if err != nil {
			panic(err)
		}
		movieIndex++

	}
}

func updateCinema() {
	url := "https://api.themoviedb.org/3/movie/now_playing?page=1&language=en-US&api_key=b57cadb923f1f664952c11dbb225bb18"
	movies := new(moviesReponse)
	getJson(url, movies)
	insertMovies(movies) //inserting new movies
	removeUnavialabeMoviesFromCinema(movies)
	updateHalls(movies)
}

func weeklyUpdate() {
	fmt.Println("Updating Cinema")
	currentTime := time.Now()
	newDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+6, 0, 0, 0, 0, currentTime.Location())
	difference := newDay.Sub(currentTime)
	if difference < 0 {
		newDay = newDay.Add(24 * time.Hour)
		difference = newDay.Sub(currentTime)
	}
	for {
		time.Sleep(difference)
		difference = 24 * time.Hour
		updateCinema() // updates movies and halls

	}
}

func insertHalls(movies *moviesReponse) {
	for movieCount := 0; movieCount < 20; movieCount++ {
		var sqlStatement string
		sqlStatement = "INSERT INTO halls (id,seats,movie) (select $1 as id, $2 as seats ,$3 as movie where not exists (select * from halls where id=$1))"
		var err error
		_, err = db.Exec(sqlStatement, movieCount, 200, movies.Results[movieCount].Id)
		if err != nil {
			panic(err)
		}
	}
}
func initCinema() {
	url := "https://api.themoviedb.org/3/movie/now_playing?page=1&language=en-US&api_key=b57cadb923f1f664952c11dbb225bb18"
	movies := new(moviesReponse)
	getJson(url, movies)
	insertMovies(movies)
	insertHalls(movies)
}

func main() {
	var err error
	config := Configuration{}
	config.Dbname = os.Getenv("DATABASE_NAME")
	config.Host = os.Getenv("DATABASE_HOST")
	config.User = os.Getenv("DATABASE_USER")
	config.Password = os.Getenv("DATABASE_PASSWORD")
	config.Port = os.Getenv("DATABASE_PORT")
	var webPort = os.Getenv("WEB_PORT")

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
	//INTIALIZE ONLY ONCE
	initCinema()
	//WEEKLY UPDATES
	//	go weeklyUpdate()
	http.HandleFunc("/", myHandler)
	//http.HandleFunc("/cache", myCachedHandler)
	log.Print("Listening on " + ":" + webPort + "...")
	http.ListenAndServe(":"+webPort, nil)
}
