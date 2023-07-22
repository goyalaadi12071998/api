package middlewares

import (
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(string(rune(time.Now().Unix())) + ":" + r.Method + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
