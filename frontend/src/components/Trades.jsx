export default function Trades({ trades }) {
  return (
    <div>
      <h3>Live Trades</h3>

      <ul style={{ maxHeight: "300px", overflowY: "auto" }}>
        {trades.map((t, i) => (
          <li key={i}>
            {t.symbol} | {t.quantity} @ {t.price}
          </li>
        ))}
      </ul>
    </div>
  );
}
