package main

import (
	// "articles/core/models"
	// "articles/services/mysql"
	"fmt"
	"net/http"

	"services/mysql"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	mysql.CheckDB()
	router = gin.New()
	router.NoRoute(noRouteHandler())
	serviceLayer := router.Group("/v1")
	service.InitRoutes(version1)

}

func main() {
	fmt.Println("Server Running on Port: ", 9090)
	http.ListenAndServe(":9090", router)
}

func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var statuscode int
		var message string = "Not Found"
		var data interface{} = nil
		var listError []models.ErrorModel = nil
		var endpoint string = c.Request.URL.String()
		var method string = c.Request.Method

		var tempEr models.ErrorModel
		tempEr.ErrorCode = 4041
		tempEr.Hints = "Not Found !! \n Routes In Valid. You enter on invalid Page/Endpoint"
		tempEr.Info = "visit http://localhost:9090/v1/docs to see the available routes"
		listError = append(listError, tempEr)
		statuscode = 404
		responseModel := &models.ResponseModel{
			statuscode,
			message,
			data,
			listError,
			endpoint,
			method,
		}
		var content gin.H = responseModel.NewResponse()
		c.JSON(statuscode, content)
	}
}
