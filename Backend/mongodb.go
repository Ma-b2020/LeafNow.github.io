package db

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	ctx    *context.Context
	client *mongo.Client
}

// Requires the MongoDB Go Driver
// https://go.mongodb.org/mongo-driver
func NewDB() *Database {
	db := Database{}
	ctx := context.TODO()
	db.ctx = &ctx
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://Tie2:Tie2@clustersample.noid7fp.mongodb.net/test")
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	db.client = client

	if err != nil {
		log.Fatal(err)
	}
	return &db
}
func (db *Database) GetSuggestions(query string) []primitive.M {

	coll := db.client.Database("sampleDB").Collection("employeeCollection")
	cursor, err := coll.Aggregate(*db.ctx, bson.A{
		bson.D{
			{"$search",
				bson.D{
					{"index", "titleIndex"},
					{"autocomplete",
						bson.D{
							{"query", query},
							{"path", "synonyms"},
						},
					},
					{"returnStoredSource", true},
				},
			},
		},
		bson.D{},
		bson.D{},
		bson.D{
			{"$project",
				bson.D{
					{"title", 1},
					{"picture", 1},
					{"type", 1},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	var suggestions []primitive.M

	cursor.All(*db.ctx, &suggestions)

	// for cursor.Next(db.ctx) {
	// 	var suggestion bson.M
	// 	err := cursor.Decode(&suggestion)

	// 	if err != nil {
	// 		print("Error")
	// 	}

	// 	suggestions = append(suggestions, suggestion)
	// }

	// defer cursor.Close(db.ctx)

	return suggestions
}

func (db *Database) HttpHandler(resp http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("q")
	suggestions := db.suggest(query)
	fmt.Println(suggestions)
}
