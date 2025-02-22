package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"payment_getway/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table created successfully")
	}
}
func SaveUserToDB(db *sql.DB, user model.User) error {
	createUserTable(db)
	query := `INSERT INTO users (id, email, password,created_at) VALUES ($1, $2, $3,$4)`
	_, err := db.Exec(query, user.ID, user.Email, user.Password, time.Now())
	return err
}
func GenerateJWT(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 20).Unix(),
		"iat":   time.Now().Unix(),
	}
	// http.SetCookie()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}
func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
