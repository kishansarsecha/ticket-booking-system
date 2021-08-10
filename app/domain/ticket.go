package domain

import "time"

type Ticket struct {
	Id      string  `json:"id,omitempty"`
	Catalog Catalog `json:"catalog" binding:"required"`
	Slot    Slot    `json:"slot" binding:"required"`
}

type Catalog struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Slot struct {
	Id   string    `json:"id" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
}
