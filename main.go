package main

import (
	ps "ddd-sample/src/presentation/screening"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/screening/v1/", ps.MakeHandler())

	http.Handle("/", accessControl(mux))

	log.Fatal(http.ListenAndServe(":4000", nil))
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
