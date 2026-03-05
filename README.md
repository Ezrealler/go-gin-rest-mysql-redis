# Go Gin REST API Blog System

A RESTful backend service built with **Go + Gin + MySQL + Redis** implementing a simple blog system with authentication, posts, comments, and categories.

This project demonstrates common backend architecture including **JWT authentication, Redis session management, CRUD APIs, pagination, and Swagger documentation**.

---

# Tech Stack

* **Go**
* **Gin Web Framework**
* **MySQL**
* **Redis**
* **JWT Authentication**
* **Swagger API Documentation**
* **Docker (MySQL container)**

---

# Project Features

### Authentication

* User login
* JWT token generation
* Redis-based token session storage
* Logout support

### User Module

* Create user
* Update user
* Delete user
* Get user information

### Post Module

* Create post
* Update post
* Delete post
* List posts with pagination
* Filter posts by category

### Comment Module

* Add comment to post
* List comments of a post

### Category Module

* Create category
* List categories
* Assign category to posts

---

# Project Structure

```
go-gin-rest-mysql-redis
│
├── configuration      # database and config setup
├── docs               # swagger generated docs
├── middleware         # authentication middleware
├── model              # data models
├── repository         # database queries
├── router             # route initialization
├── service            # business logic / controllers
├── util               # jwt, redis utilities
├── resource           # configuration files
├── dump.sql           # database schema
└── main.go            # application entry
```

---

# Getting Started

## 1 Clone Repository

```
git clone https://github.com/shahzadthathal/go-gin-rest-mysql-redis.git
cd go-gin-rest-mysql-redis
```

---

# 2 Install Dependencies

```
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get github.com/DATA-DOG/go-sqlmock
go get github.com/stretchr/testify/assert
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/segmentio/ksuid
go get github.com/gomodule/redigo/redis
```

---

# 3 Install Swagger

```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

Generate docs:

```
swag init
```

---

# 4 Setup Database

Start MySQL (Docker recommended)

Import database schema:

```
docker exec -i demo-mysql mysql -uroot -proot ms_account_dev < dump.sql
```

---

# 5 Configure Database

Edit configuration inside:

```
resource/config.yaml
```

Example:

```
DB:
  USER_NAME: root
  PASSWORD: root
  HOST_NAME: 127.0.0.1:3307
  NAME: ms_account_dev
```

---

# 6 Run Tests

```
go test -v
```

---

# 7 Run Application

```
go run main.go
```

---

# Swagger API Documentation

Open in browser:

```
http://localhost:8999/swagger/index.html
```

---

# Default Admin Account

```
username: admin
password: admin1234
```

---

# Example API

## Login

```
POST /api/login
```

Body:

```
{
  "username":"admin",
  "password":"admin1234"
}
```

---

## Create Post

```
POST /api/posts
```

Header:

```
Authorization: Bearer <token>
```

Body:

```
{
  "title": "My first post",
  "description": "hello world",
  "category_id": 1,
  "status": 1
}
```

---

# Future Improvements

* Post like system
* Nested comments
* Redis hot ranking
* API rate limiting
* File upload support
* CI/CD integration

---

# License

MIT License
