package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redcowe/blog_api/model"
	"go.mongodb.org/mongo-driver/bson"
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

//helper function for initalizing checking and initialized db
func checkInitDB() error {
	if !initialized {
		err := initDB()
		if err != nil {
			return err
		}
	}
	return nil
}

func AddBlog(b *model.BlogPost) error {
	err := checkInitDB()
	if err != nil {
		return err
	}

	collection := DB.Database("blog_db").Collection("blogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Inserting document
	_, err2 := collection.InsertOne(ctx, b)
	if err2 != nil {
		log.Fatal("Unable to insert into DB", err)
		return err2
	}

	return nil
}

//Returns all blogs in db
func GetBlogs() ([]bson.M, error) {
	checkInitDB()

	collection := DB.Database("blog_db").Collection("blogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	//variable for holding blogs
	blogs := []bson.M{}
	cur.All(ctx, &blogs)

	return blogs, nil
}

func RemoveBlog() error {
	//TODO
	return nil
}

func UpdateBlog(b *model.BlogPost) error {
	//TODO
	return nil
}
