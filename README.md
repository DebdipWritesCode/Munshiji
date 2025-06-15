# 📊 Munshiji — The MUN Scoresheet Manager

**Munshiji** (मुन्शीजी) is a powerful, AI-assisted scoring and feedback system tailored for Model United Nations (MUN) conferences. It functions like an intelligent digital accountant, helping executive boards manage scores, generate feedback, and export professional reports with ease.

---

## 📚 Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture Overview](#architecture-overview)
- [Database Schema](#database-schema)
- [Authentication Flow](#authentication-flow)
- [Parameter Rule Engine](#parameter-rule-engine)
- [AI Feedback & ZIP Export](#ai-feedback--zip-export)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Future Roadmap](#future-roadmap)
- [Screenshots](#screenshots)
- [License](#license)

---

## 🚀 Features

- 🧑‍⚖️ Full CRUD for users, score sheets, delegates, parameters, and scores.
- 🧠 LLM (OpenAI) integration for generating detailed delegate feedback.
- 🔐 Secure authentication using access/refresh tokens with cookies and memory-only storage.
- 📄 PDF generation with pie charts + performance metrics per delegate.
- 📦 Combined ZIP export of Excel sheets + individual feedback reports.
- 🧮 Dynamic scoring logic based on custom parameter rules (average, absolute, weighted).
- 🧰 Form validations using Zod.
- 📈 Charts.js visualizations for scoring breakdowns.
- 🖱️ Beautiful UI with TailwindCSS v4 + shadcn/ui components.
- 🔄 Session throttling for AI (5-minute cooldown per user).
- 🐳 Dockerized PostgreSQL setup.

---

## 🛠️ Tech Stack

### Backend
- **Language**: Golang
- **Database**: PostgreSQL (via `lib/pq`)
- **Codegen**: [`sqlc`](https://sqlc.dev) for typed DB access
- **Server**: gRPC + gRPC Gateway
- **Auth**: Access/refresh token with HMAC symmetric key
- **Sessions**: Stored in PostgreSQL
- **AI**: OpenAI API (batched + concurrent)
- **Build Tools**: `make`, `golang-migrate`

### Frontend
- **Framework**: React (TypeScript)
- **Routing**: React Router
- **State**: Redux Toolkit
- **Validation**: Zod
- **Styling**: TailwindCSS v4
- **Components**: shadcn/ui
- **Data Fetching**: Axios with interceptor
- **Charts**: Chart.js

---

## 🧱 Architecture Overview

```

monorepo/
│
├── backend/                # Go gRPC server with PostgreSQL
│   ├── db/                 # sqlc-generated DB access
│   ├── gapi/               # gRPC handlers and services
│   ├── token/              # JWT auth with HMAC
│   ├── proto/              # Protocol buffers + gateway annotations
│   ├── Makefile            # Backend commands (migrate, sqlc, proto, run)
│   └── main.go             # Entry point with concurrent gRPC + HTTP
│
├── frontend/               # Vite + React frontend
│   └── ...                 # Components, Redux, UI

```
---

## 🗃️ Database Schema Overview

- **users**: Auth (email, password)
- **sessions**: Refresh token sessions
- **score_sheets**: Top-level sheet tied to users
- **delegates**: Linked to specific score sheet
- **parameters**: Scoring rules & metadata
- **scores**: Stores value + note (delegate × parameter)
- **ai_sessions**: Tracks last AI request time per user

---

## 🔐 Authentication Flow

- **Access Token**:
  - Expiry: 15 minutes
  - Stored in-memory via Redux Toolkit
  - Sent via `Authorization: Bearer <token>` header

- **Refresh Token**:
  - Expiry: 48 hours
  - Stored as HTTP-only cookie
  - Auto-refreshed on page reload via `/refresh-access-token`

- **Security**:
  - No localStorage/sessionStorage used
  - Backend validates token on every protected route via middleware
  - Frontend Axios interceptor attaches access token automatically

---

## 🧮 Parameter Rule Engine

Each parameter has a `rule_type`:
- `average`: Mean of all score values.
- `absolute`: Sum of absolute values.
- `special`: Weighted formula:
```

score = (avg\_score \* score\_weight) + (num\_scores \* length\_weight)

````

Example:
```ts
Scores: [4, 6, 7, 8]
avg_score = 6.25
length = 4

score = 6.25 * 0.7 + 4 * 0.3 = 5.575
````

* Dynamic recalculations happen in Redux via `scoresSlice`.
* Top 3 delegates are ranked and highlighted dynamically.

---

## 🤖 AI Feedback & ZIP Export

* Batch scores in groups of 5 → each group has a custom prompt.
* Responses from OpenAI handled in concurrent goroutines.
* Each response is rendered to a delegate-specific PDF:

  * Includes feedback
  * Parameter-wise pie chart
* Excel sheet generated from current score data.
* All PDFs + Excel files zipped and downloaded on frontend.

---

## ⚙️ Getting Started

```bash
# Clone the repo
git clone https://github.com/DebdipWritesCode/Munshiji
cd munshiji
```

### Backend Setup

```bash
cd backend
go mod tidy

# Create and apply DB migrations
make migrateup

# (Optional) Generate SQLC + Proto code
make sqlc
make proto

# Run server
make run
```

### Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

---

## 🔑 Environment Variables

### Backend `.env` (`backend/app.env`)

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
DB_DRIVER=
DB_SOURCE=

REFRESH_TOKEN_DURATION=48h
ACCESS_TOKEN_DURATION=15m

HTTP_SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:9000
ENVIRONMENT=development

TOKEN_SYMMETRIC_KEY=
OPENAI_API_KEY=
```

### Frontend `.env` (`frontend/.env`)

```
VITE_BACKEND_URL=http://localhost:8080
```

---

## 🛣️ Future Roadmap

* 📬 Redis + email verification flow
* 📤 Request-based access to other's score sheets
* 🔒 Role-based permissions for shared sheets
* 🧾 Better session analytics dashboard
* 🌐 Deployment via Docker Compose + CI/CD

---

## 🖼️ Screenshots

### 🧾 Delegate Score Sheet UI
Displays a dynamic and editable score sheet with parameter-wise scoring, special rule evaluation, and total calculation. Top 3 delegates are highlighted based on total score.

![Delegate Score Sheet UI](frontend/public/score-sheet-ui.png)

---

### 📝 Feedback PDF Preview
Automatically generated PDF per delegate, including AI-powered feedback and parameter-wise performance pie chart.

![Feedback PDF Preview](frontend/public/feedback-pdf-preview.png)


---

## 📄 License

MIT License © 2025 [Debdip Mukherjee](https://github.com/DebdipWritesCode)

---

> *Built with care and concurrency 💙*
