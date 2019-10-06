package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bukhavtsov/user-directory/db"
	"github.com/bukhavtsov/user-directory/pkg/api"
	"github.com/bukhavtsov/user-directory/pkg/data"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	host     = "localhost"
	port     = "postgres"
	user     = "postgres"
	dbname   = "postgres"
	password = "postgres"
	sslmode  = "disable"
)

var (
	serverEndpoint = os.Getenv("SERVER_ENDPOINT")
)

func init() {
	if serverEndpoint == "" {
		serverEndpoint = ":8080"
	}
}

func main() {
	r := mux.NewRouter()
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()
	api.ServeUserResource(r, data.NewUser(conn))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	})
	handler := cors.Default().Handler(r)

	log.Println("serving server at ", serverEndpoint)
	log.Fatal(http.ListenAndServe(":8080", handler))
}