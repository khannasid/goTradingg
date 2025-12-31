package models

import "time"

type OrderSide string

const Buy OrderSide = "BUY"
const Sell OrderSide = "SELL"

type Order struct {
	ID        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Side      OrderSide `json:"side"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
}

type Trade struct{
	BuyOrderID  string    `json:"buy_order_id"`
	SellOrderID string    `json:"sell_order_id"`
	Symbol      string    `json:"symbol"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Timestamp   time.Time `json:"timestamp"`
}