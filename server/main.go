package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hsn0918/mongodb_project/apis"
	"github.com/hsn0918/mongodb_project/svc"
)

var SvcContext *svc.ServiceContext

func init() {
	SvcContext = svc.NexServiceContext()
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {

	})
	college := r.Group("/college")
	{
		college.GET("/collegename", apis.CollegeName)
		college.GET("/dean", apis.Dean)
		college.GET("/collegeteacher", apis.CollegeTeacher)
	}
	r.Run(":8888")
}
