import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement,
  Tooltip,
} from "chart.js";

ChartJS.register(LineElement, CategoryScale, LinearScale, PointElement, Tooltip);

export default function PriceChart({ trades }) {
  if (!trades || trades.length === 0) {
    return (
      <div style={{ color: "var(--muted)", textAlign: "center" }}>
        No trades yet. Chart will update after first match.
      </div>
    );
  }

  const prices = trades.slice(0, 20).reverse();

  const data = {
    labels: prices.map((_, i) => i + 1),
    datasets: [
      {
        label: "Price",
        data: prices.map((t) => t.price),
        borderColor: "#38bdf8",
        tension: 0.4,
      },
    ],
  };

  return <Line data={data} />;
}

