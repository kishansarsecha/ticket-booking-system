package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sujithps/ticket-booking-system/app/domain"
	"github.com/sujithps/ticket-booking-system/app/http"
)

func BookTicket(context *gin.Context) (int, http.APIResponse) {
	var ticket domain.Ticket
	err := context.ShouldBindJSON(&ticket)

	if err != nil {
		println(fmt.Sprintf("error while booking the ticket %s", err.Error()))
		return 400, http.CreateErrorApiResponse(err)
	}

	responseData := domain.Ticket{
		Id: "1",
		Catalog: ticket.Catalog,
		Slot: ticket.Slot,
	}

	return 201, http.CreateSuccessApiResponse(responseData)
}
