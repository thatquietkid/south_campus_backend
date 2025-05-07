# South Campus Backend

![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go\&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-4.x-blue?logo=go\&label=Echo)
![Docker](https://img.shields.io/badge/Docker-Ready-blue?logo=docker)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Ready-blue?logo=postgresql)

This is a backend API for the South Campus App (University of Delhi), built using Go and Echo framework. It handles authentication and management of courses, cafeteria items, events, bus routes, hostels, complaints, and more, with role-based access control using JWT.

> \[!NOTE]
> This API is live and is being hosted on [Render](https://render.com/) -> [Visit this link](https://south-campus-backend.onrender.com).

---

## 📦 Project Structure

```
south_campus_backend/
├── config/              # Database configuration
├── handlers/            # Route handler functions
├── middleware/          # Custom middleware for auth and admin
├── models/              # GORM models
├── routes/              # Route registration
├── main.go              # Entry point
└── README.md            # Project documentation
```

---

## 🔧 Setup

Make sure you have [Docker](https://www.docker.com/products/docker-desktop) and [Docker Compose](https://docs.docker.com/compose/) installed.

### 1. Clone and Navigate

```bash
git clone https://github.com/thatquietkid/south_campus_backend.git
cd south_campus_backend
```

### 2. Configure Database

Edit your PostgreSQL credentials in `config/db.go`:

```go
dsn := "host=localhost user=south_campus_user password=12345 dbname=south_campus port=5432 sslmode=disable"
```

### 3. Auto-Migrate Tables

The following tables are migrated:

* `users`
* `courses`
* `cafeteria_items`
* `events`
* `bus_routes`
* `hostels`
* `complaints`

### 4. Start the Server

```bash
sudo docker-compose up --build
```

---

## 🔐 Authentication (JWT)

Login endpoint issues a JWT if credentials match:

```http
POST /login
Content-Type: application/json

{
  "username": "adminuser",
  "password": "adminpass"
}
```

Returns:

```json
{
  "token": "<JWT_TOKEN>"
}
```

Use this token in headers:

```http
Authorization: Bearer <JWT_TOKEN>
```

---

## 🧑‍⚖️ Roles

* `role = admin` → Full access (POST/DELETE)
* `role = user` → Read-only access

---

## 📚 API Endpoints

### Public

* `GET /` – Welcome message
* `GET /courses` – List all courses
* `GET /course-syllabus/:code` – Get course syllabus by course code
* `GET /course-attendance` – Get course attendance info
* `GET /cafeteria-items` – List all cafeteria items
* `GET /hostels` – List all hostels
* `GET /bus-routes` – List all bus routes
* `GET /events` – List all campus events
* `GET /complaints` – List all complaints
* `POST /complaints` – Submit a complaint

### Authenticated (JWT required)

* `POST /login` – Login and get JWT token

### Admin-only (JWT required + role `admin`)

* `POST /courses` – Create a course
* `DELETE /courses/:id` – Delete a course
* `POST /cafeteria-items` – Add cafeteria item
* `DELETE /cafeteria-items/:id` – Delete cafeteria item

---

## 🥪 Sample Users

Insert this sample admin user (password is bcrypt-hashed):

```sql
INSERT INTO users (username, password, role) VALUES (
  'adminuser',
  '$2a$10$7ykk8XAxmIbX8BML6Yo6A.K8U5ZChItutZ1zCDmf6hsmNzEZ16Yl2',
  'admin'
);
```

---

## 📨 Postman Sample Requests

### Login

```bash
curl -X POST http://localhost:1323/login \
  -H "Content-Type: application/json" \
  -d '{"username": "adminuser", "password": "adminpass"}'
```

### Create Course (Admin Only)

```bash
curl -X POST http://localhost:1323/courses \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"name": "Math 101", "description": "Intro to Math"}'
```

### Get Courses

```bash
curl http://localhost:1323/courses
```

---

## 📄 License

[MIT License](LICENSE).

---

## Built by [@thatquietkid](https://github.com/thatquietkid/)
