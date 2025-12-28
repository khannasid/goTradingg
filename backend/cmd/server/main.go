package main

import (
	"log"
	"net/http"

	"goTradingg/internal/orderbook"
	"goTradingg/internal/transport"
)

func main(){
	ob := &orderbook.OrderBook{
		Symbol : "TSLA",
	}

	handler := transport.NewHTTPHandler(ob)

	log.Println("🚀 goTradingg server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}