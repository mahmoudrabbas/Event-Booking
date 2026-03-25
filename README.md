# 📌 Event Booking REST API (Go + Gin)

A RESTful API for managing events and user registrations, built using **Go** and the **Gin framework**.

---

## 🚀 Features

- Full CRUD for events
- User authentication (Signup & Login)
- Event registration & cancellation
- JWT-based authorization

---

## 📡 API Endpoints

### 📅 Events

| Method | Endpoint      | Description                |
| ------ | ------------- | -------------------------- |
| GET    | `/events`     | Get all available events   |
| GET    | `/events/:id` | Get a specific event by ID |
| POST   | `/events`     | Create a new event         |
| PUT    | `/events/:id` | Update an existing event   |
| DELETE | `/events/:id` | Delete an event            |

---

### 👤 Authentication

| Method | Endpoint  | Description         |
| ------ | --------- | ------------------- |
| POST   | `/signup` | Create a new user   |
| POST   | `/login`  | Authenticate a user |

---

### 🎟️ Event Registration

| Method | Endpoint               | Description                |
| ------ | ---------------------- | -------------------------- |
| POST   | `/events/:id/register` | Register user for an event |
| DELETE | `/events/:id/register` | Cancel event registration  |

---

## 🔐 Authentication (JWT Flow)

1. User signs up via `/signup`
2. User logs in via `/login`
3. Server returns a JWT token
4. Client includes token in requests:

```
Authorization: Bearer <your_token>
```

5. Protected routes:
    - Create event
    - Update event
    - Delete event
    - Register / cancel registration

---

## 📥 Request & Response Examples

### ✅ Signup

**Request:**

```json
POST /signup

{
  "email": "user@example.com",
  "password": "123456"
}
```

**Response:**

```json
{
    "message": "User created successfully"
}
```

---

### 🔑 Login

**Request:**

```json
POST /login

{
  "email": "user@example.com",
  "password": "123456"
}
```

**Response:**

```json
{
    "token": "jwt_token_here"
}
```

---

### 📅 Create Event

**Request:**

```json
POST /events

{
    "name":"event1",
    "description":"description 1",
    "location":"minia",
    "dateTime":"2026-04-02T15:30:00.000Z"
}
```

**Response:**

```json
{
    "events": [
        {
            "Id": 1,
            "Name": "event1",
            "Description": "description 1",
            "Location": "minia",
            "DateTime": "2026-04-02T15:30:00Z",
            "UserId": 1
        }
    ]
}
```

---

### 🎟️ Register for Event

**Request:**

```
POST /events/1/register
Authorization: Bearer <token>
```

**Response:**

```json
{
    "message": "Successfully registered"
}
```

---

## 📦 Project Structure

```
project/
│
├── main.go
├── go.mod
│
├── controllers/
│   ├── event_controller.go
│   ├── auth_controller.go
│
├── models/
│   ├── event.go
│   ├── user.go
│
├── routes/
│   └── routes.go
│
├── middleware/
│   └── auth_middleware.go
│
├── utils/
│   └── jwt.go
│
└── database/
    └── db.go
```

---

## 🛠️ Tech Stack

- Go (Golang)
- Gin Web Framework
- JWT Authentication
- REST API Architecture

---

## ▶️ Getting Started

```bash
git clone <https://github.com/mahmoudrabbas/Event-Booking>
cd project
go mod tidy
go run main.go
```

---

## 📌 Notes

- All protected routes require JWT authentication
- Use JSON for all requests and responses
- IDs are passed as URL parameters (`:id`)

---

## 📄 License

This project is for learning purposes and practice with Go & backend development.
