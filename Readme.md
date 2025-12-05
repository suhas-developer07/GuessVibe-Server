# **VI-Sense**

*A dual-language (Go + Python) cognitive prediction engine using LLMs.*

---

##  Overview

**VI-Sense** is an intelligent system designed to **predict what’s on a user’s mind** by continuously asking context-aware questions.
It establishes a real-time session with the user, processes each answer, and uses an LLM-powered reasoning loop to generate the next best question — ultimately converging on a prediction.

The system is built with:

* **Golang** (WebSocket server + gRPC client)
* **Python** (LLM inference service using LangChain + Groq API)
* **Redis** (session storage)
* **MongoDB** (user data persistence)

---

##  How VI-Sense Works

### 1️ **Client → Go Service (WebSockets)**

* The mobile app connects to the Go service over **WebSockets**.
* Go creates a **unique session** per game and stores session metadata in **Redis**.

### 2️ **Go → Python (gRPC)**

* For each user answer, the Go server sends the data to the Python LLM service using **gRPC `.proto` contracts**.
* Python receives the context and generates the **next question** using LangChain + Groq.

### 3️ **Python → Go → User**

* Python returns the generated question via gRPC.
* Go pushes it to the client over WebSockets in real-time.

### 4️ **Prediction Loop**

* Redis stores all Q/A pairs per session.
* As context grows, the LLM becomes more confident and predicts the final intent or thought.

---

##  Tech Stack

### **Backend Services**

| Component          | Technology                                   |
| ------------------ | -------------------------------------------- |
| **Go Service**     | Echo HTTP Framework, WebSockets, gRPC Client |
| **Python Service** | LangChain, Groq API, gRPC Server             |
| **LLM Provider**   | Groq (Mixtral, Llama, etc.)                  |
| **Databases**      | MongoDB (user data), Redis (session storage) |

---

##  Project Structure (Simplified)

```
VI-Sense-backend/
│── Go/                 # WebSocket + gRPC client
│   ├── cmd/            #Entry point + database
│   ├── internal/         
│   ├── pkg/
│   └── proto
│
│── Python/             # LLM microservice
│   ├── app/
│   ├── proto/
│   ├── Dockerfile
│   └── main.py
│
└── README.md
```

---

##  Environment Variables

Create a `.env` file in both Go and Python services:

```
MONGODB_URI=
REDIS_URL=
DATABASE_NAME=
PORT=
LLM_PORT=
LLM_HOST=
GROQ_API_KEY=
```

---

##  Local Development Setup

### 1️ Clone the Repository

```bash
git clone https://github.com/suhas-developer07/VI-Sense-backend.git
```

---

### 2️ Run the Go Service

```bash
cd Go
make run  or go run ./cmd
```

Starts:

* WebSocket server
* gRPC client
* Redis integration

---

### 3️ Run the Python LLM Service

```bash
cd Python
pip install -r requirement.txt
python main.py
```

Starts:

* gRPC server
* LangChain pipeline
* Groq-powered question generator

---

##  Communication Flow

1. **Mobile App** connects → Go WebSocket
2. Go creates a session → stores in Redis
3. User answers → Go sends to Python via gRPC
4. LLM → generates next question
5. Python → Go → User
6. Loop continues until prediction

---

##  Deployment

Currently deployed using:

* **Railway.app** for both Go and Python services
* Uses internal networking for gRPC communication

https://sixthsense-production.up.railway.app/

---

##  Features

* Real-time question-answer loop
* Context-aware LLM prompting
* Predictive reasoning engine
* Redis-backed session storage
* Microservice architecture (Go ↔ Python)
* Scalable WebSocket infrastructure
* LangChain prompt orchestration
* gRPC high-performance communication

---

##  Live App Preview

### ** Try the App on Expo**

🔗 **[https://expo.dev/accounts/manojrustcult/projects/guess-game/builds/078401f4-441e-4282-9130-832fd703b5e5](https://expo.dev/accounts/manojrustcult/projects/guess-game/builds/078401f4-441e-4282-9130-832fd703b5e5)**

---
##  License

Licensed under the **Apache License 2.0**.
See [`LICENSE`](LICENSE) file for more details.

---

##  Contributing

Pull requests are welcome!
For major changes, please open an issue to discuss what you would like to improve.

---

##  Support the Project

If you like VI-Sense, please ⭐ the repository on GitHub — it helps a lot!

---
