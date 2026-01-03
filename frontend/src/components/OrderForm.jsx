import { useState } from "react";

export default function OrderForm() {
  const [side, setSide] = useState("BUY");
  const [price, setPrice] = useState("");
  const [quantity, setQuantity] = useState("");
  const API_BASE = import.meta.env.VITE_API_BASE_URL;

  const submitOrder = async () => {
    const parsedPrice = Number(price);
    const parsedQty = Number(quantity);

    if (!parsedPrice || !parsedQty) return;

    await fetch(`${API_BASE}/order`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        side,
        price: parsedPrice,
        quantity: parsedQty,
      }),
    });

    // reset safely
    setPrice("");
    setQuantity("");
  };

  return (
    <div style={{ marginBottom: "20px" }}>
      <h3>Place Order</h3>

      <select value={side} onChange={(e) => setSide(e.target.value)}>
        <option value="BUY">BUY</option>
        <option value="SELL">SELL</option>
      </select>

      <input
        type="number"
        placeholder="Price"
        value={price}
        onChange={(e) => setPrice(e.target.value)}
      />

      <input
        type="number"
        placeholder="Quantity"
        value={quantity}
        onChange={(e) => setQuantity(e.target.value)}
      />

      <button onClick={submitOrder}>Submit</button>
    </div>
  );
}
