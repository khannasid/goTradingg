import { useEffect, useState } from "react";
import OrderForm from "./components/OrderForm";
import Trades from "./components/Trades";
import OrderBook from "./components/OrderBook";

export default function App() {
  const [trades, setTrades] = useState([]);
  const [orderBook, setOrderBook] = useState({
    buy_orders: [],
    sell_orders: [],
  });

  // WebSocket for live trades
  useEffect(() => {
  let ws;

  try {
    ws = new WebSocket("ws://localhost:8080/ws/trades");

    ws.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.onmessage = (event) => {
      const trade = JSON.parse(event.data);
      setTrades((prev) => [trade, ...prev]);
    };

    ws.onerror = (err) => {
      console.error("WebSocket error", err);
    };
  } catch (err) {
    console.error("WebSocket init failed", err);
  }

  return () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.close();
    }
  };
}, []);


  // Poll order book (simple & reliable for now)
  useEffect(() => {
    const fetchOrderBook = async () => {
      try {
        const res = await fetch("http://localhost:8080/orderbook");
        const data = await res.json();
        setOrderBook(data);
      } catch (err) {
        console.error("Failed to fetch orderbook", err);
      }
    };

    fetchOrderBook();
    const interval = setInterval(fetchOrderBook, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div style={{ padding: "24px" }}>
      <h1>📈 goTradingg</h1>

      <div className="card">
        <OrderForm />
      </div>

      <div style={{ display: "flex", gap: "24px", marginTop: "24px" }}>
        <div className="card">
          <OrderBook orderBook={orderBook} />
        </div>

        <div className="card">
          <Trades trades={trades} />
        </div>
      </div>
    </div>
  );
}
