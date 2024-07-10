package handler

import (
	"encoding/json"
	"net/http"
	"payment_getway/controller"
	"payment_getway/db"
	"payment_getway/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	database, err := db.DbIn()
	if err != nil {
		panic(err)
	}
	defer database.Close()
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)

	err = controller.SaveUserToDB(database, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": user.ID.String()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	database, err := db.DbIn()
	if err != nil {
		panic(err)
	}
	defer database.Close()
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err = json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user model.User
	err = database.QueryRow("SELECT id, email, password FROM users WHERE email = $1", credentials.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := controller.GenerateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
