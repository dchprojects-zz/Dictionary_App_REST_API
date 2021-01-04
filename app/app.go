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
