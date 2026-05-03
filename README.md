## 🚀 Features

- RESTful APIs built using Gin

- Modular and scalable project structure (Clean Architecture)

- PostgreSQL integration using GORM

- Request validation

- Secure password hashing using bcrypt

- JWT-based authentication with middleware

- Token blacklisting (logout invalidation using Redis)

- Environment-based configuration (.env)

- DTO layer for clean request/response handling

- Centralized error handling

- Rate limiting using Redis (5 requests/minute)

## ⚙️ In Progress

- Implementing Role Based Access Control

- Advanced rate limiting strategies (per-route / per-user)

## 🛠 Tech Stack

- Go (Golang)

- Gin Framework

- GORM ORM

- PostgreSQL

- Redis

## 📂 Folder Strcture

.
├── cmd
│ └── main.go
├── config
│ ├── db.go
│ ├── env.go
│ └── redis.go
├── docs
│ ├── docs.go
│ ├── swagger.json
│ └── swagger.yaml
├── go.mod
├── go.sum
├── helpers
│ └── strc.go
├── internal
│ ├── dto
│ │ ├── product.go
│ │ └── user.go
│ ├── handlers
│ │ ├── auth.go
│ │ └── product.go
│ ├── middlewares
│ │ ├── auth.go
│ │ └── ratelimit.go
│ ├── models
│ │ ├── product.go
│ │ └── user.go
│ ├── repositories
│ │ ├── auth.go
│ │ ├── product.go
│ │ └── redis.go
│ ├── routes
│ │ └── routes.go
│ ├── services
│ │ ├── auth.go
│ │ └── product.go
│ └── utils
│ ├── error.go
│ ├── jwt.go
│ ├── pass.go
│ └── response.go
└── README.md

## 🎨 Design Decisions

# Centralized Error Handling

- A unified error structure ensures consistent API responses.
- This also makes it easier to extend later with logging, error codes, and tracing.

# Folder Structure

- The project follows a layered architecture:
  `Handler → Service → Repository → Database`

-- Improves separation of concerns

-- Makes the system easier to test and maintain

-- Allows independent scaling of components

- The use of internal/ ensures encapsulation and prevents unintended external usage.

# Rate Limiting

- Currently implemented using Redis with a fixed window approach.

- Trade-off:
  - At window boundaries, users may send bursts (e.g., multiple requests at the edge of a time window).

Future improvements:

- Sliding window log (more accurate control)
- Token bucket (better handling of bursts)

## ▶️ Run Locally

1. Clone repo
2. Create `.env` file
3. Run:
   go run main.go

`Make Sure GO Is Installed On Your Device`
