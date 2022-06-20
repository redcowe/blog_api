package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redcowe/blog_api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB, initialized = &mongo.Client{}, false

func initDB() error {

	//Loading ENV Variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
		return err
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not found.")
	}

	//Initalizing connection to DB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	DB = client
	initialized = true

	return nil
}

func AddBlog(b *model.BlogPost) error {
	if !initialized {
		initDB()
	}

	collection := DB.Database("blog_db").Collection("blogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Inserting document
	res, err := collection.InsertOne(ctx, b)

	if err != nil {
		log.Fatal("Unable to insert into DB", err)
		return err
	}
	fmt.Println(res.InsertedID)

	return nil
}
