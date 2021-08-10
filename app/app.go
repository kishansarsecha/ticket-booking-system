package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sujithps/ticket-booking-system/app/handler"
	"net/http"
)

func StartServer() {
	server := gin.Default()

	v1 := server.Group("/v1")

	{
		v1.POST("/book", handleBookTicket())

		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"status": "running"})
		})
	}

	if err := server.Run(":8080"); err != nil {
		_ = fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}

	fmt.Println("Server Started")
}

func handleBookTicket() func(context *gin.Context) {
	return func(context *gin.Context) {
		status, response := handler.BookTicket(context)
		context.JSON(status, response)
	}
}
