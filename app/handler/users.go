package handler

import (
	"goproj/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		users := []model.User{}
		respondJSON(w, http.StatusOK, users)
	} else {
		respondError(w, http.StatusMethodNotAllowed, "StatusMethodNotAllowed")
	}
}

// getUserOr404 gets a task instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
