package orderbook

import (
	"testing"
	"time"
	"goTradingg/internal/models"
)

func TestOrderBookInitialization(t *testing.T) {
	ob := &OrderBook{
		Symbol: "TSLA",
	}

	if ob.Symbol != "TSLA" {
		t.Fatalf("expected symbol TSLA, got %s", ob.Symbol)
	}

	if len(ob.BuyOrders) != 0 {
		t.Fatalf("expected empty buy orders, got %d", len(ob.BuyOrders))
	}

	if len(ob.SellOrders) != 0 {
		t.Fatalf("expected empty sell orders, got %d", len(ob.SellOrders))
	}
}

func TestSimpleBuySellMatch(t *testing.T) {
	ob := &OrderBook{
		Symbol: "TSLA",
	}

	buy := models.Order{
		ID:        "buy-1",
		Side:      models.Buy,
		Price:     110,
		Quantity:  10,
		Timestamp: time.Now(),
	}

	sell := models.Order{
		ID:        "sell-1",
		Side:      models.Sell,
		Price:     100,
		Quantity:  10,
		Timestamp: time.Now(),
	}

	// Add orders
	ob.AddOrder(buy)
	ob.AddOrder(sell)

	// Execute matching
	trades := ob.Match()

	if len(trades) != 1 {
		t.Fatalf("expected 1 trade, got %d", len(trades))
	}

	trade := trades[0]

	if trade.Symbol != "TSLA" {
		t.Fatalf("expected symbol TSLA, got %s", trade.Symbol)
	}

	if trade.Price != 100 {
		t.Fatalf("expected trade price 100, got %f", trade.Price)
	}

	if trade.Quantity != 999 {
		t.Fatalf("expected trade quantity 999, got %d", trade.Quantity)
	}
}
