# Greenlight — RESTful Movie Management API

A production-ready backend API built with Go for managing movies data, handling user accounts, and supporting secure authentication flows.  

---

## Features

### **Movie Management**
- Full CRUD operations for movie resources  
- Filtering, pagination, and text-based search 
- Custom validation and structured JSON responses  

### **User Accounts**
- User registration with activation email  
- Time-limited activation tokens    
- Role-based access control (RBAC)  

### **Security**
- Rate limiting to prevent request abuse  
- Robust authentication/authorization middleware  
- Input sanitization and strict custom validation  

### **Infrastructure**
- Dockerized app + PostgreSQL using Docker Compose    
- Environment-based configuration  
- Graceful shutdown and structured logging  
- Database migrations
- Prometheus metrics

---

## 🛠️ Tech Stack

| Layer | Technology |
|-------|------------|
| Language | Go |
| Router | Chi |
| Database | PostgreSQL |
| Auth | JWT |
| Email | SMTP |
| Proxy | Caddy (planned) |
| Deployment | Docker & Docker Compose |

---

## 📁 Project Structure (simplified)

```bash
.
├── bin
├── cmd
│   ├── api
│   └── examples
├── go.mod
├── go.sum
├── internal
│   ├── data
│   ├── mailer
│   └── validator
├── Makefile
├── migrations
│   ...
├── README.md
└── remote
```

---

## API Routes

### Healthcheck

| Method | Endpoint          | Description                       |
| ------ | ----------------- | --------------------------------- |
| GET    | `/v1/healthcheck` | Returns service and system status |

### Movies

| Method | Endpoint          | Permission     | Description                           |
| ------ | ----------------- | -------------- | ------------------------------------- |
| GET    | `/v1/movies`      | `movies:read`  | List all movies (with filters/search) |
| POST   | `/v1/movies`      | `movies:write` | Create a new movie                    |
| GET    | `/v1/movies/{id}` | `movies:read`  | Fetch a single movie by ID            |
| PATCH  | `/v1/movies/{id}` | `movies:write` | Update movie fields                   |
| DELETE | `/v1/movies/{id}` | `movies:write` | Delete a movie                        |


### Users

| Method | Endpoint              | Description                                  |
| ------ | --------------------- | -------------------------------------------- |
| POST   | `/v1/users`           | Register a new user account                  |
| PUT    | `/v1/users/activated` | Activate user via token (email verification) |
| POST   | `/v1/tokens/authentication` | Create an authentication token (login) |

### Metrics

| Method | Endpoint      | Description                 |
| ------ | ------------- | --------------------------- |
| GET    | `/v1/metrics` | Prometheus metrics endpoint |

---

## Tech Stack

* **Go**
* **Chi Router**
* **PostgreSQL**
* **Prometheus**
* **Docker & Docker Compose**
