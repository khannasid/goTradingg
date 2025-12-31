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

	broadcaster := transport.NewTradeBroadcaster()
	handler := transport.NewHTTPHandler(ob, broadcaster)

	log.Println("🚀 goTradingg server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}