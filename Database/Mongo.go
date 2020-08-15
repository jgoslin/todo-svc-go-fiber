package Database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
	CTX    context.Context
}

const (
	URI = "mongodb://localhost:27017"
	DB  = "Sample"
)

func (mgi *MongoInstance) Connect() error {
	//Setup Mongo DB Connection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	//Connect to DB and validate connection.
	if client != nil && err == nil {
		db := client.Database(DB)
		if db != nil {
			mgi.Client = client
			mgi.DB = db
			mgi.CTX = ctx
		} else {
			return errors.New("unable to connect to database")
		}
	} else {
		return err
	}

	//Test connection to Mongo
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
