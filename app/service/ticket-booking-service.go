package service

import (
	"github.com/sujithps/ticket-booking-system/app/domain"
	"github.com/sujithps/ticket-booking-system/app/handler/contract"
)

func BookTicket(request contract.BookingRequest) domain.Ticket {
	return domain.Ticket{
		Id:      "1",
		Catalog: request.Catalog,
		Slot:    request.Slot,
	}
}