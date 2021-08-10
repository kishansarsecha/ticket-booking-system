package handler

import (
	"github.com/sujithps/ticket-booking-system/app/domain"
	"github.com/sujithps/ticket-booking-system/app/http"
	"time"
)

func BookTicket() http.APIResponse {
	data := domain.Ticket{
		Entity:  domain.Entity{
			Id: "1",
		},
		Catalog: domain.Catalog{
			Id:   "c1",
			Name: "Movie 1",
		},
		Slot:    domain.Slot{
			Id:   "s1",
			Date: time.Now(),
		},
	}
	return http.CreateSuccessApiResponse(data)
}