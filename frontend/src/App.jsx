import { useEffect, useState } from "react";
import OrderForm from "./components/OrderForm";
import Trades from "./components/Trades";
import OrderBook from "./components/OrderBook";
import PriceChart from "./components/PriceChart";

export default function App() {
  const [trades, setTrades] = useState([]);
const API_BASE = import.meta.env.VITE_API_BASE_URL;
  const [orderBook, setOrderBook] = useState({
    buy_orders: [],
    sell_orders: [],
  });

  // WebSocket for live trades
  useEffect(() => {
    const ws = new WebSocket(`ws://${API_BASE.replace("http://", "")}/ws/trades`);

    ws.onmessage = (event) => {
      const trade = JSON.parse(event.data);
      setTrades((prev) => [trade, ...prev]);
    };

    return () => ws.close();
  }, []);

  // Poll order book
  useEffect(() => {
    const fetchOrderBook = async () => {
      const res = await fetch(`${API_BASE}/orderbook`);
      const data = await res.json();
      setOrderBook(data);
    };

    fetchOrderBook();
    const interval = setInterval(fetchOrderBook, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="container">
      {/* Heading */}
      <div className="center">
        <h1>📈 goTradingg</h1>
        <p style={{ color: "var(--muted)" }}>
          Real-time trading simulator powered by Go & WebSockets
        </p>
      </div>

      {/* Order Form */}
      <div className="card section">
        <OrderForm />
      </div>

      {/* Trades */}
      <div className="card section">
        <Trades trades={trades} />
      </div>

      {/* OrderBook + Chart */}
      <div className="grid section">
        <div className="card">
          <OrderBook orderBook={orderBook} />
        </div>

        <div className="card">
          <h3>Price Chart</h3>
          <PriceChart trades={trades} />
        </div>
      </div>
    </div>
  );
}
