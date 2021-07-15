package handler

import (
	"goproj/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllWords(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		words := []model.Word{}
		respondJSON(w, http.StatusOK, words)
	} else {
		respondError(w, http.StatusMethodNotAllowed, "StatusMethodNotAllowed")
	}
}

// getWordOr404 gets a task instance if exists, or respond the 404 error otherwise
func getWordOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Word {
	word := model.Word{}
	if err := db.First(&word, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &word
}
