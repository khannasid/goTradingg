export default function OrderBook({ orderBook }) {
  const buyOrders = orderBook?.buy_orders || [];
  const sellOrders = orderBook?.sell_orders || [];

  return (
    <div>
      <h3>Order Book</h3>

      <div style={{ display: "flex", gap: "30px" }}>
        <div>
          <h4>BUY</h4>
          {buyOrders.map((o, i) => (
            <div key={i}>
              {o.quantity} @ {o.price}
            </div>
          ))}
        </div>

        <div>
          <h4>SELL</h4>
          {sellOrders.map((o, i) => (
            <div key={i}>
              {o.quantity} @ {o.price}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
