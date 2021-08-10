package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sujithps/ticket-booking-system/app/handler"
	"net/http"
)

func StartServer() {

	server := gin.Default()

	server.POST("/book", func(context *gin.Context) {
		context.JSON(http.StatusCreated, handler.BookTicket())
	})

	if err := server.Run(":8080"); err != nil {
		_ = fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}

	fmt.Println("Server Started")
}

func pingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
