package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redcowe/blog_api/db"
	"github.com/redcowe/blog_api/model"
)

func main() {

	r := gin.Default()

	//default route
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "get off my route!!"})
	})

	//route for displaying all blogs
	r.GET("/blogs", func(ctx *gin.Context) {
		blogs, _ := db.GetBlogs()
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		ctx.JSON(http.StatusOK, blogs)
	})

	r.POST("/blog", func(ctx *gin.Context) {
		blog := model.BlogPost{
			Title:          "Test Post 2: Eletric Boogaloo",
			Author:         "Ya mama ",
			CreateDateTime: time.Now().String(),
			UpdateDateTime: time.Now().String(),
			Body:           []string{"First paragraph", "second paragraph", "https://images.unsplash.com/photo-1490730141103-6cac27aaab94?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8ZnJlZXxlbnwwfHwwfHw%3D&w=1000&q=80", "closing paragraph"},
			Tags:           []string{"Test", "Picture", "Go", "Employed"},
		}
		db.AddBlog(&blog)
		ctx.JSON(http.StatusOK, gin.H{"message": "ok :)"})
	})
	r.Run()
}
