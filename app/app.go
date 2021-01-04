package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/api/words").
		HandlerFunc(app.getWords)
	app.Router.
		Methods("POST").
		Path("/api/addWord").
		HandlerFunc(app.addWord)
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

func (app *App) addWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		uuid := r.FormValue("uuid")
		word := r.FormValue("word")
		translatedWord := r.FormValue("translated_word")
		createdDate := r.FormValue("created_date")
		insForm, err := app.Database.Prepare("INSERT INTO `words` (uuid, word, translated_word, created_date) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(uuid, word, translatedWord, createdDate)
		w.WriteHeader(http.StatusNoContent)
	}
}
