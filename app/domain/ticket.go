package domain

import "time"

type Ticket struct {
	Id      string  `json:"id"`
	Catalog Catalog `json:"catalog"`
	Slot    Slot    `json:"slot"`
}

type Catalog struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Slot struct {
	Id   string    `json:"id"`
	Date time.Time `json:"date"`
}
