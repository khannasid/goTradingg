package models

import "time"

type OrderSide string

const Buy OrderSide = "BUY"
const Sell OrderSide = "SELL"

type Order struct {
	ID string
	Symbol string
	Side OrderSide
	Price float64
	Quantity int
	Timestamp time.Time
}

type Trade struct{
	BuyOrderID string
	SellOrderID string
	Symbol string
	Price float64
	Quantity int
	Timestamp time.Time
}