export default function Trades({ trades }) {
  return (
    <div>
      <h3>Live Trades</h3>

      <ul style={{ maxHeight: "300px", overflowY: "auto" }}>
        {trades.map((t, i) => (
          <li
            key={i}
            style={{
              display: "flex",
              justifyContent: "space-between",
              marginBottom: "6px",
              color: t.side === "BUY" ? "var(--buy)" : "var(--sell)",
            }}
          >
            <span>{t.symbol}</span>

            <span className={`badge ${t.side === "BUY" ? "badge-buy" : "badge-sell"}`}>
              {t.side}
            </span>

            <span>
              {t.quantity} @ {t.price}
            </span>
          </li>
        ))}
      </ul>
    </div>
  );
}
