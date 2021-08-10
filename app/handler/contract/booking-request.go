package contract

import "github.com/sujithps/ticket-booking-system/app/domain"

type BookingRequest struct {
	Catalog domain.Catalog
	Slot domain.Slot
}