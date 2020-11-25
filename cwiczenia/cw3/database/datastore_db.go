package database

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

var datastoreClient *datastore.Client

func init () {
	ctx := context.Background()
	c, err := datastore.NewClient(ctx, "mcourser-dev")

	if err != nil {
		log.Panic(err)
	}

	log.Println("Created db connection")

	datastoreClient = c
}

func getAllFromDB() ([]*DBItem, error) {
	ctx := context.Background()
	var output []*DBItem

	log.Println("Going to retriev from db")
	_, err := datastoreClient.GetAll(ctx, datastore.NewQuery("db_item_entry"), &output)
	log.Println("Retrieved from db")
	if err != nil {
		return nil, err
	}

	return output, nil
}

func addToDB (name, description string, priority Priority) (*DBItem, error) {
	var priorityInt int64 = 0

	if priority == MEDIUM {
		priorityInt = 1
	} else if priority == HIGH {
		priorityInt = 2
	}

	item := makeDBItem(name, description, priorityInt)
	_, err := datastoreClient.Put(
		context.Background(),
		datastore.IncompleteKey("db_item_entry", nil),
		&item,
	)

	return &item, err
}