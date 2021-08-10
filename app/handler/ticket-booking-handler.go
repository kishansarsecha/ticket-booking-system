package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sujithps/ticket-booking-system/app/handler/contract"
	"github.com/sujithps/ticket-booking-system/app/service"
	"net/http"
)

func BookTicket(context *gin.Context) (int, contract.APIResponse) {
	var bookingRequest contract.BookingRequest
	err := context.ShouldBindJSON(&bookingRequest)

	if err != nil {
		println(fmt.Sprintf("error while booking the ticket %s", err.Error()))
		return http.StatusBadRequest, contract.CreateErrorApiResponse(err)
	}

	responseData := service.BookTicket(bookingRequest)

	return http.StatusCreated, contract.CreateSuccessApiResponse(responseData)
}
