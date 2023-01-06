package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ashwin/go-mongodb/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://<username>:<password>@cluster0.xtqw0ie.mongodb.net/?retryWrites=true&w=majority"
const dbName string = "netflix"
const colName string = "watchlist"

var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection reference is ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserterd one movie with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", deleteCount.DeletedCount)
}

func deleteAllMovies() int64 {

	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", deleteCount.DeletedCount)

	return deleteCount.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

//handle func

func GetAllMovies(c *gin.Context) {
	allMovies := getAllMovies()
	c.IndentedJSON(http.StatusOK, allMovies)
}

func CreateMovie(c *gin.Context) {
	var movie model.Netflix
	if err := c.BindJSON(&movie); err != nil {
		log.Fatal(err)
	}
	insertOneMovie(movie)

	c.IndentedJSON(http.StatusOK, movie)
}

func MarkAsWatched(c *gin.Context) {
	id := c.Param("id")
	updateOneMovie(id)

	c.IndentedJSON(http.StatusOK, id)
}

func DeleteOneMovie(c *gin.Context) {
	id := c.Param("id")
	deleteOneMovie(id)

	c.IndentedJSON(http.StatusOK, id)
}

func DeleteAllMovies(c *gin.Context) {
	count := deleteAllMovies()

	c.IndentedJSON(http.StatusOK, count)
}
