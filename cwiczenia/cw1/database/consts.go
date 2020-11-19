package database

import (
	"github.com/sonyarouje/simdb/db"
	"log"
)

var driver *db.Driver


func init () {
	var err error
	driver, err = db.New("data")

	if err != nil {
		log.Panic("Could not connect to db")
	}
}