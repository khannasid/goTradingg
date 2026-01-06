# goTradingg 🚀
A **production-grade, real-time trading simulator** built with **Go and React**, featuring a
concurrency-safe matching engine, WebSocket-based live trade streaming, cloud deployment,
and CI/CD automation.

> Designed to demonstrate **backend engineering, real-time systems, distributed systems fundamentals,
and production DevOps practices**.

---

## 🔥 Key Highlights (Recruiter-Friendly)
- Real-time trading system (BUY/SELL matching)
- Go-based concurrency-safe matching engine
- WebSocket (WSS) live trade streaming
- Dockerized backend & frontend
- Cloud deployment (Railway + Vercel)
- CI/CD using GitHub Actions
- Unit testing with Go testing framework
- Environment-based configuration & CORS handling
- Designed for horizontal scalability and distributed extensions

---

## 🧠 System Architecture

Browser (React UI)
|
| HTTPS / WSS
v
Frontend (Vercel)
|
| REST / WebSocket
v
Backend API (Go, Railway)
|
v
In-memory OrderBook & Matching Engine


---

## 🧩 Core Components

### 1️⃣ Matching Engine (Go)
- Implements **price–time priority**
- Thread-safe using mutexes
- Deterministic trade execution
- Designed for future sharding by symbol

### 2️⃣ Real-Time Streaming
- WebSocket-based pub/sub broadcaster
- Multiple clients supported concurrently
- Secure WSS communication in production

### 3️⃣ REST APIs
- Place BUY / SELL orders
- Fetch live order book
- JSON contract enforcement

---

## ⚙️ Technology Stack

### Backend
- Go (Golang)
- net/http
- Goroutines, mutexes
- WebSockets
- Docker

### Frontend
- React (Vite)
- WebSockets
- Chart.js
- Environment-based config

### DevOps / Cloud
- Docker & Docker Compose
- Railway (Backend hosting)
- Vercel (Frontend hosting)
- GitHub Actions (CI/CD)

---

## 🧪 Testing
- Unit tests written for order matching logic
- Tests executed automatically in CI pipeline
- Deployment blocked if tests fail

```bash
go test ./...
```
## 🔁 CI/CD Pipeline
- On every push to main:
- Backend tests executed
- Backend auto-deployed to Railway
- Frontend auto-deployed to Vercel

## ☁️ Deployment
- Backend: Railway (Docker-based deployment)
- Frontend: Vercel
- Secure HTTPS + WSS
- Environment-driven configuration

## 🚀 Future Enhancements (Distributed Systems)
- Symbol-based sharding of order books
- Message queue (Kafka / NATS) for ingestion
- Persistent storage (Postgres / Redis)
- Leader election & replication
- Fault tolerance & replay mechanisms
- Observability (metrics, tracing)

## 🏆 Learning Outcomes
- Real-time systems design
- Concurrency & correctness
- Docker & cloud deployments
- CI/CD pipelines
- Production debugging (CORS, WSS, env config)
- Distributed systems fundamentals

## 👨‍💻 Author

Siddhant Khanna
Backend Engineer
