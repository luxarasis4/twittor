package middleware

import (
	"net/http"

	"github.com/luxarasis4/twittor/database"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !database.CheckConnection() {
			http.Error(w, "Lost database connection", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
