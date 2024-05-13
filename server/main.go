package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hsn0918/mongodb_project/svc"
)

var svcContext *svc.ServiceContext

func init() {
	svcContext = svc.NexServiceContext()
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		book, _ := svcContext.BookModel.FindOne(c, "6641e33e922ae44e35b9fa3b")
		c.JSON(200, book)

	})

	r.Run(":8888")
}
