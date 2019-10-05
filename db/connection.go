package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetConnection(host, port, user, dbname, password, sslmode string) *gorm.DB {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	connection, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatalln(err)
	}
	return connection
}
