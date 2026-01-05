package main

import (
	"log"
	"net/http"
	"os"

	"goTradingg/internal/orderbook"
	"goTradingg/internal/transport"
)

func main(){
	ob := &orderbook.OrderBook{
		Symbol : "TSLA",
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	broadcaster := transport.NewTradeBroadcaster()
	handler := transport.NewHTTPHandler(ob, broadcaster)

	log.Println("🚀 goTradingg server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}