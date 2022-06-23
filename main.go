package main

import (
	"fmt"
	"log"

	"github.com/redcowe/blog_api/db"
)

func main() {
	blogs, err := db.GetBlogs()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, value := range blogs {
		fmt.Println(value)
	}
}
