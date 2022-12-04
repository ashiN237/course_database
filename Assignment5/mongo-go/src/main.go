package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func MongoOperation() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	log.Println(os.Getenv("MONGO_URI\n"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Connection to MongoDB\n")

	var posts []Post

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: status code %s\n", resp.StatusCode)
		return err
	}

	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &posts); err != nil {
		fmt.Println(err)
		return err
	}

	col := client.Database("test").Collection("sample_collections")

	if err = insert(col, posts); err != nil {
		return err
	}

	if err = findOne(col); err != nil {
		return err
	}

	if err = findMany(col); err != nil {
		return err
	}

	if err = col.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}

func insert(col *mongo.Collection, posts []Post) error {
	interface_posts := make([]interface{}, len(posts))
	for i := range posts {
		interface_posts[i] = posts[i]
	}
	_, err := col.InsertMany(context.Background(), interface_posts)
	return err
}

func findOne(col *mongo.Collection) error {
	var doc bson.Raw
	findOptions := options.FindOne()
	err := col.FindOne(context.Background(), bson.D{}, findOptions).Decode(&doc)

	if err == mongo.ErrNoDocuments {
		log.Println("Documents not found")
		return err
	} else if err != nil {
		return err
	}
	fmt.Println(doc.String())
	return err
}

func findMany(col *mongo.Collection) error {
	filter := struct {
		UserId int
	}{2}

	findOptions := options.Find()
	cur, err := col.Find(context.Background(), filter, findOptions)
	if err != nil {
		return err
	}

	for cur.Next(context.Background()) {
		var doc Post
		if err = cur.Decode(&doc); err != nil {
			return err
		}

		fmt.Printf("%d %s\n", doc.Id, doc.Title)
	}

	return err
}

func main() {
	if err := MongoOperation(); err != nil {
		log.Fatal(err)
	}
	log.Println("normal end.")
}