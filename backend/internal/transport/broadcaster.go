/*
Each WebSocket connection runs in its own goroutine and owns a buffered channel. 
The broadcaster fans out trade events using non-blocking sends, so slow clients 
don’t affect system throughput. Channels are just synchronization primitives; 
goroutines perform the work.
*/
package transport

import (
	"sync"

	"goTradingg/internal/models"
)

type TradeBroadcaster struct {
	mu sync.Mutex
	clients map[chan models.Trade]struct{}
}

func NewTradeBroadcaster() *TradeBroadcaster {
	return &TradeBroadcaster{
		clients: make(map[chan models.Trade]struct{}),
	}
}

func (b *TradeBroadcaster) Register(ch chan models.Trade) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.clients[ch] = struct{}{}
}

func (b *TradeBroadcaster) Unregister(ch chan models.Trade) {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.clients, ch)
	close(ch)
}

func (b *TradeBroadcaster) Broadcast(trade models.Trade) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for ch := range b.clients {
		select {
		case ch <- trade:
		default:
			// drop if client is slow (important!)
		}
	}
}
