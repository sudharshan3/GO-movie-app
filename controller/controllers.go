package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudharshan3/GO-movie-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//functions
func InsertOneMovie(movie model.Movies) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie", inserted.InsertedID)
}
func UpdateOneModel(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated one movie", updated.ModifiedCount)
}
func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted one movie", deletecount)
}
func DeleteAllMovies() int64 {
	deletecount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all movies", deletecount.DeletedCount)
	return deletecount.DeletedCount
}

//methods
func GetAllMovies() []model.Movies {
	var movies []model.Movies
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := GetAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	var movie model.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	params := mux.Vars(r)
	UpdateOneModel(params["id"])
	json.NewEncoder(w).Encode("Movie marked watched:")

}
func DeleteaMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	params := mux.Vars(r)
	DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie deleted")
}

func DeletemyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	DeleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted")
}
