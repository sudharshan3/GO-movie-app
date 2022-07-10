package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/sudharshan3/GO-movie-app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sudharshan:shansud3198@cluster0.yemeq.mongodb.net/?retryWrites=true&w=majority"
const dbname = "movieapp"
const collname = "movies"

var collection *mongo.Collection

func init() {
	ConnectionOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), ConnectionOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB conneciton Success!")
	collection = client.Database(dbname).Collection(collname)

}
func InsertOneMovie(movie model.Movies) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie", inserted.InsertedID)
}
