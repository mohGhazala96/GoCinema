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
	_ "time"

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

var db *sql.DB

func myHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title FROM movies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "Welcome to the server\n")
	fmt.Fprintln(w, "ID | Title")
	fmt.Fprintln(w, "---+--------")
	for rows.Next() {
		var (
			id    int
			title string
		)

		rows.Scan(&id, &title)

		fmt.Fprintf(w, "%2d | %s\n", id, title)
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

	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))

	return json.NewDecoder(res.Body).Decode(target)
}
func main() {
	var err error

	url := "https://api.themoviedb.org/3/movie/now_playing?page=1&language=en-US&api_key=b57cadb923f1f664952c11dbb225bb18"

	movies := new(moviesReponse)
	getJson(url, movies)
	fmt.Println(movies.Results[0].Id)
	fmt.Println(movies.Results[0].Original_language)
	//http://image.tmdb.org/t/p/w500/
	fmt.Println(movies.Results[1].Poster_path)
	fmt.Println(movies.Results[0].Title)

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

	http.HandleFunc("/", myHandler)
	//http.HandleFunc("/cache", myCachedHandler)
	log.Print("Listening on " + ":" + webPort + "...")
	http.ListenAndServe(":"+webPort, nil)

}
