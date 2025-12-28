package orderbook

import (
	"sort"
	"sync"
	"time"
	"goTradingg/internal/models" // Remember! In Go, the module root is backend/, not the repo root. So no need to add backend here.
								// As we have initialized goTradingg as module name in go.mod file inside backend/, go treats backend/ as module root.
)

type OrderBook struct {
	Symbol string

	BuyOrders []models.Order
	SellOrders []models.Order

	mu sync.Mutex
}

// this sortBooks function is a receiver function of OrderBook struct
// when a function is declared like:
// func (some variable with struct "kinda OOPs like concept") <func_name> (params) { ... }
// so it gets easy to call this function, when ever we declare an instance of the struct
// this function will be present to be used. Eg: 
// ob := OrderBook{Symbol: "BTCUSD"}
// ob.sortBooks()  --> this is how we can call this function

func (ob *OrderBook) sortBooks(){
	sort.Slice(ob.BuyOrders, func(i, j int)bool{
		if ob.BuyOrders[i].Price == ob.BuyOrders[j].Price{
			return ob.BuyOrders[i].Timestamp.Before(ob.BuyOrders[j].Timestamp) // here, the earlier timestamp order has higher priority (will be first)
		}
		return ob.BuyOrders[i].Price > ob.BuyOrders[j].Price
	})

	sort.Slice(ob.SellOrders, func(i, j int)bool{
		if ob.SellOrders[i].Price == ob.SellOrders[j].Price{
			return ob.SellOrders[i].Timestamp.Before(ob.SellOrders[j].Timestamp)
		}
		return ob.SellOrders[i].Price < ob.SellOrders[j].Price
	})
}

func (ob *OrderBook) AddOrder(order models.Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock() // defer is used to ensure ki the lock gets unlocked at the end of the function execution

	if order.Side == models.Buy {
		ob.BuyOrders = append(ob.BuyOrders, order)
	}else{
		ob.SellOrders = append(ob.SellOrders, order)
	}
	ob.sortBooks()
}

// See this function returns a slice of trades that were executed as a result of matching orders in the order book.

func (ob *OrderBook) Match() []models.Trade{
	ob.mu.Lock() //obviously we have to lock it, as changes are happening at the central data structure aka the orderbook 
	defer ob.mu.Unlock()

	// here in this function, we have to return the trade struct (defined in models package)
	trades := []models.Trade{}

	for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0{
		buy := ob.BuyOrders[0]
		sell := ob.SellOrders[0]

		if buy.Price < sell.Price{
			break
		}

		qty := min(buy.Quantity, sell.Quantity)

		trade := models.Trade{
			BuyOrderID:  buy.ID,
			SellOrderID: sell.ID,
			Symbol:      ob.Symbol,
			Price:       sell.Price,
			Quantity:    qty,
			Timestamp:   time.Now(),
		}

		trades = append(trades, trade)

		buy.Quantity -= qty
		sell.Quantity -= qty

		if buy.Quantity == 0{
			ob.BuyOrders = ob.BuyOrders[1:]
		}else{
			ob.BuyOrders[0] = buy
		}

		if sell.Quantity == 0 {
			ob.SellOrders = ob.SellOrders[1:]
		} else {
			ob.SellOrders[0] = sell
		}
	}

	return trades
}

func (ob *OrderBook) Snapshot() (buys []models.Order, sells []models.Order){
	ob.mu.Lock()
	defer ob.mu.Unlock()

	// Return copies to avoid external mutation
	buys = append([]models.Order(nil), ob.BuyOrders...)
	sells = append([]models.Order(nil), ob.SellOrders...)

	return
}