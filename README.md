# 🚀 Go Learning Journey - Todo REST API

> **Learning in Progress** 📚  
> This repository documents my journey learning Go programming language by building a practical REST API with modern tools and best practices.

## 🎯 Project Overview

A REST API for managing todo items built with Go, featuring:

- **RESTful Architecture** with clean endpoints
- **Database Integration** using GORM ORM with PostgreSQL
- **Containerized Development** with Docker Compose
- **Modern Go Patterns** and best practices
- **Auto-migration** for database schema management

## 🛠️ Tech Stack

| Category             | Technology              |
| -------------------- | ----------------------- |
| **Language**         | Go 1.21+                |
| **Web Framework**    | Gin                     |
| **Database**         | PostgreSQL 15           |
| **ORM**              | GORM                    |
| **Containerization** | Docker & Docker Compose |
| **Database UI**      | Adminer                 |

## 🏗️ Project Structure

```
playground/
├── main.go              # Main application entry point
├── docker-compose.yml   # Docker services configuration
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksums
└── README.md           # Project documentation
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Git

### 1. Clone & Setup

```bash
git clone <your-repo-url>
cd playground
```

### 2. Start Database Services

```bash
docker-compose up -d
```

This will start:

- **PostgreSQL** database on `localhost:2345`
- **Adminer** database UI on `localhost:8081`

### 3. Run the Application

```bash
go run main.go
```

The API will be available at `http://localhost:8080`

## 📊 Database Access

### Via Adminer Web UI

1. Open http://localhost:8081
2. Login with:
   - **System**: PostgreSQL
   - **Server**: postgres
   - **Username**: todo_admin
   - **Password**: admin
   - **Database**: tododb

### Via Command Line

```bash
# Connect to PostgreSQL container
docker exec -it go-db psql -U todo_admin -d tododb

# List tables
\dt

# View table structure
\d todo_items
```

## 🔌 API Endpoints

### Base URL: `http://localhost:8080/v1`

| Method | Endpoint | Description                        | Example                         |
| ------ | -------- | ---------------------------------- | ------------------------------- |
| `GET`  | `/items` | Get all todo items with pagination | `GET /v1/items?page=1&limit=10` |
| `POST` | `/items` | Create a new todo item             | `POST /v1/items`                |

### 📝 API Examples

#### Create Todo Item

```bash
curl -X POST http://localhost:8080/v1/items \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go fundamentals"}'
```

**Response:**

```json
{ "data": 1 }
```

#### Get Todo Items

```bash
curl http://localhost:8080/v1/items
```

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "title": "Learn Go fundamentals",
      "status": "Doing",
      "created_at": "2025-08-05T09:45:13.663192+07:00",
      "updated_at": "2025-08-05T09:45:13.663192+07:00"
    }
  ]
}
```

## 🎓 Learning Objectives

### ✅ Completed

- [x] Go project setup with modules
- [x] HTTP server with Gin framework
- [x] Database integration with GORM
- [x] PostgreSQL connection and configuration
- [x] Docker containerization
- [x] REST API design and implementation
- [x] JSON request/response handling
- [x] Database migrations and schema management
- [x] Error handling and validation
- [x] API pagination

### 🎯 Next Steps

- [ ] Authentication & Authorization (JWT)
- [ ] Input validation with custom validators
- [ ] Middleware implementation (logging, CORS)
- [ ] Unit testing with testify
- [ ] API documentation with Swagger
- [ ] Configuration management with Viper
- [ ] Graceful shutdown handling
- [ ] Metrics and monitoring
- [ ] Deployment with Docker

## 🚦 Development Commands

```bash
# Install dependencies
go mod tidy

# Run the application
go run main.go

# Start database only
docker-compose up postgres -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Reset database (caution: deletes all data)
docker-compose down -v && docker-compose up -d
```

## 🐛 Troubleshooting

### Common Issues

1. **Port already in use**

   ```bash
   # Kill process using port 8080
   lsof -ti:8080 | xargs kill -9
   ```

2. **Database connection refused**

   ```bash
   # Check if PostgreSQL container is running
   docker-compose ps
   ```

3. **Permission denied errors**
   ```bash
   # Reset database with proper permissions
   docker-compose down -v
   docker-compose up -d
   ```

## 📚 Resources & References

- [Go Official Documentation](https://golang.org/doc/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM Documentation](https://gorm.io/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker Compose Reference](https://docs.docker.com/compose/)

## 🤝 Contributing

This is a personal learning project, but feedback and suggestions are welcome! Feel free to:

- Open issues for questions or suggestions
- Share learning resources
- Provide code review feedback

## 📄 License

This project is for educational purposes. Feel free to use it as a reference for your own learning journey.

---

**Happy Coding!** 🎉  
_Building confidence in Go, one API endpoint at a time._
