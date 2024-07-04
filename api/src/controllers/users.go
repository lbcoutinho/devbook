package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllUsers"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserById"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
