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
	app.Router.
		Methods("PUT").
		Path("/api/updateWord").
		HandlerFunc(app.updateWord)
	app.Router.
		Methods("DELETE").
		Path("/api/deleteWord").
		HandlerFunc(app.deleteWord)
}

func (app *App) getWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var words []WordModel

		result, err := app.Database.Query("SELECT id, uuid, word, translated_word, created_date FROM `words`")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer result.Close()
		for result.Next() {
			var word WordModel
			err := result.Scan(&word.ID, &word.UUID, &word.Word, &word.TranslatedWord, &word.CreatedDate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			words = append(words, word)
		}
		json.NewEncoder(w).Encode(words)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *App) updateWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "PUT" {
		var updateWord UpdateWordModel

		err := json.NewDecoder(r.Body).Decode(&updateWord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		updForm, err := app.Database.Prepare("UPDATE `words` SET word = ?, translated_word = ? WHERE uuid = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		updForm.Exec(updateWord.Word, updateWord.TranslatedWord, updateWord.UUID)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *App) addWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var word WordModel

		err := json.NewDecoder(r.Body).Decode(&word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insForm, err := app.Database.Prepare("INSERT INTO `words` (uuid, word, translated_word, created_date) VALUES (?,?,?,?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		insForm.Exec(word.UUID, word.Word, word.TranslatedWord, word.CreatedDate)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *App) deleteWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		var deleteWord DeleteWordModel
		err := json.NewDecoder(r.Body).Decode(&deleteWord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		delForm, err := app.Database.Prepare("DELETE FROM `words` WHERE uuid = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		delForm.Exec(deleteWord.UUID)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
