package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "192.168.20.101:5432"
	user     = "postgres"
	dbname   = "iot"
	sslmode  = "disable"
	password = "lgh"
)

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, user, dbname, sslmode, password))
	if err != nil {

	}
	defer db.Close()
}
