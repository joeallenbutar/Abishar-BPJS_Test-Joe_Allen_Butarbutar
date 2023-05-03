package models

import (
	"time"
)

type RequestData struct {
	Id          int       `json:"id"`
	Customer    string    `json:"customer"`
	Quantity    int 	  `json:"quantity"`
	Price       float64   `json:"price"`
	Timestamp   time.Time `json:"timestamp"`
	RequestId   int       `json:"request_id"`
}

type Order struct {
	RequestId    int       		  `json:"request_id"`
	CreatedAt    time.Time 		  `json:"created_at"`
	RequestData  []RequestData    `json:"data"`
}