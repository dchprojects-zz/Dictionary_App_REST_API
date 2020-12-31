package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type WordModel struct {
	ID             int    `json:"id"`
	UUID           string `json:"uuid"`
	Word           string `json:"word"`
	TranslatedWord string `json:"translated_word"`
	CreatedDate    string `json:"created_date"`
}

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func main() {
	database, err := CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8000", app.Router))
}

func CreateDatabase() (*sql.DB, error) {

	user := "d9d9vs9"
	password := "Vagina$2020"
	dbName := "dbword"

	db, err := sql.Open("mysql", user+":"+password+"@/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/words").
		HandlerFunc(app.getWords)
}

func (app *App) getWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var words []WordModel

	result, err := app.Database.Query("SELECT id, uuid, word, translated_word, created_date FROM `words`")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var word WordModel
		err := result.Scan(&word.ID, &word.UUID, &word.Word, &word.TranslatedWord, &word.CreatedDate)
		if err != nil {
			panic(err.Error())
		}
		words = append(words, word)
	}
	json.NewEncoder(w).Encode(words)
}
