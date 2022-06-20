package main

import (
	"fmt"
	"time"

	"github.com/redcowe/blog_api/db"
	"github.com/redcowe/blog_api/model"
)

func main() {
	fmt.Println("Sitting pretty :)")
	myBlogPost := model.BlogPost{
		Title:           "My First Blog Post!",
		Author:          "Joshua Cowell",
		CreatedDateTime: time.Now().String(),
		UpdateDateTime:  time.Now().String(),
		Body:            "First blog post!!!!",
		Tags:            []string{},
	}
	db.AddBlog(&myBlogPost)
}
