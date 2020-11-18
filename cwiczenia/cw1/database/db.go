package database

import (
	"github.com/sonyarouje/simdb/db"
)
// https://godoc.org/github.com/sonyarouje/simdb/db

var driver, err = db.New("data")


func addToDB (name, description string, priority Priority) (error, *DBItem) {
	var priorityInt uint = 0

	if priority == MEDIUM {
		priorityInt = 1
	} else if priority == HIGH {
		priorityInt = 2
	}

	item := makeDBItem(name, description, priorityInt)
	return driver.Insert(item), &item
}
