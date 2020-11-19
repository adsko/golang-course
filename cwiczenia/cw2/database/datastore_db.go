package database

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

var datastoreClient *datastore.Client

func init () {
	ctx := context.Background()
	c, err := datastore.NewClient(ctx, "testApp")

	if err != nil {
		log.Panic(err)
	}

	datastoreClient = c
}

