package main

import (
	"log"
	"net/http"

	"github.com/bukhavtsov/user-directory/db"
	"github.com/bukhavtsov/user-directory/pkg/api"
	"github.com/bukhavtsov/user-directory/pkg/data"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = "postgres"
	user     = "postgres"
	dbname   = "postgres"
	password = "postgres"
	sslmode  = "disable"
)

func main() {
	r := mux.NewRouter()
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()
	api.ServeUserResource(r, data.NewUser(conn))
	log.Fatal(http.ListenAndServe(":8080", r))
}
