package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type Configuration struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

var config Configuration
var url string

type User struct {
	id       int
	username string
}

//structs for movies
type MoviesList struct {
	Movies []Movies
}

type HallsList struct {
	Halls []Halls
}

type Halls struct {
	Id    int
	Seats int
	Movie int
}

type Reservation struct {
	Id        int
	Hall      int
	Seats     []string
	Movie     int
	Useremail string
	Timing    int
}

type Movies struct {
	Id             int64
	Vote_average   float64
	Title          string
	Poster_path    string
	Original_title string
	Overview       string
	Release_date   string
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

func querymovies(movies *MoviesList) error {
	rows, err := db.Query("SELECT id, title,release_date,poster_path,vote_average,overview FROM movies")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		movie := Movies{}
		err = rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Release_date,
			&movie.Poster_path,
			&movie.Vote_average,
			&movie.Overview)

		if err != nil {
			return err
		}

		movies.Movies = append(movies.Movies, movie)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func queryhalls(halls *HallsList) error {
	rows, err := db.Query("SELECT id, seats,movie FROM halls")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		hall := Halls{}
		err = rows.Scan(
			&hall.Id,
			&hall.Seats,
			&hall.Movie)

		if err != nil {
			return err
		}

		halls.Halls = append(halls.Halls, hall)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	movies := MoviesList{}
	err := querymovies(&movies)
	// fmt.Println(movies)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	out, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

type NewReservation struct {
	Title string `json:"title,omitempty"`
}

func addReservationHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	movies := MoviesList{}
	err = querymovies(&movies)
	// fmt.Println(movies)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	out, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	params := mux.Vars(r)
	var newReservation NewReservation
	_ = json.NewDecoder(r.Body).Decode(&newReservation)
	newReservation.Title = params["title"]
	log.Printf("New Reservation full body %s\n Title %s\n Decoded is %s \n .title is %s\n Body is %s\n", r.Body, newReservation, json.NewDecoder(r.Body),
		newReservation.Title, body)

	fmt.Fprintf(w, string(out))
}

func hallsHandler(w http.ResponseWriter, r *http.Request) {

	halls := HallsList{}
	err := queryhalls(&halls)
	// fmt.Println(halls)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	out, err := json.Marshal(halls)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func InsertReservation(w http.ResponseWriter, r *http.Request) {
	reservation := &Reservation{}
	err := json.NewDecoder(r.Body).Decode(reservation) //decode the request body into struct and failed if any error occur
	if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

	log.Printf("Reservation Hall is %v + Movie is %v \n", reservation.Hall, reservation.Movie)

	// InsertReservationInDb(reservation) //Create account
	Respond(w, Message(false, "inserted succesfully"))
}

func InsertReservationInDb(reservation *Reservation) {
	for seat := range reservation.Seats {
		var sqlStatement string
		sqlStatement = "INSERT INTO reservations (hall, seat, movie, useremail,timing) Values($1,$2,$3,$4,$5)"
		var err error
		_, err = db.Exec(sqlStatement, reservation.Hall, seat, reservation.Movie, reservation.Useremail, reservation.Timing)
		if err != nil {
			panic(err)
		}
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
		sqlStatement = "INSERT INTO movies (id,title, release_date, poster_path, vote_average,overview,isAvialabe) (select $1 as id, $2 as title ,$3 as release_date,$4 as poster_path,$5 as vote_average ,$6 as overview,$7 as isAvialabe where not exists (select * from movies where id=$1))"
		var err error
		_, err = db.Exec(sqlStatement, movies.Results[movieIndex].Id, movies.Results[movieIndex].Title,
			movies.Results[movieIndex].Release_date,
			"http://image.tmdb.org/t/p/w500/"+movies.Results[movieIndex].Poster_path,
			movies.Results[movieIndex].Vote_average, movies.Results[movieIndex].Overview, true)
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

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})

	handler := c.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + webPort,
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	url = "https://api.themoviedb.org/3/movie/now_playing?page=1&language=en-US&api_key=" + os.Getenv("API_Key")

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//INTIALIZE ONLY ONCE
	// initCinema()
	//WEEKLY UPDATES
	//	go weeklyUpdate()
	router.HandleFunc("/api/getMovies/", moviesHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getHalls/", hallsHandler).Methods("GET")
	router.HandleFunc("/api/addReservation/", addReservationHandler).Methods("POST")
	router.HandleFunc("/api/insert", InsertReservation).Methods("POST")

	log.Print("Listening on " + ":" + webPort + "...")
	log.Fatal(srv.ListenAndServe())

}
