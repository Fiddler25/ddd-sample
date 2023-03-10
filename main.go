package main

import (
	"ddd-sample/ent"
	"ddd-sample/screening"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, client := ent.New()
	defer client.Close()

	repo := screening.NewRepository(client)
	svc := screening.NewService(repo)

	mux := http.NewServeMux()
	mux.Handle("/screening/v1/", screening.MakeHandler(svc))

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
