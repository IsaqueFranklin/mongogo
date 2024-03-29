package common

import (
  "context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
  return db.Collection(col)
}


func InitDB() error {
  uri := os.Getenv("MONGO_URI")
  
  if uri == "" {
    return errors.New("No mongoDB URI.")
  }

  client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

  ir err != nil {
    return err
  }

  db = client.Database("go_demo")

  return nil
}


func CloseDB() error {
  return db.Client().Disconnect(context.Background())
}
