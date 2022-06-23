package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redcowe/blog_api/db"
)

func main() {
	blogs, err := db.GetBlogs()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()
	r.GET("/blogs", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, blogs)
	})
	r.Run()
}
