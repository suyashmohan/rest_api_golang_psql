# What is this?

This repository contains code to demonstrate a basic REST API written in Go language and PostgreSQL.

The APIs lets you create and save `Notes`.

# Requirements
* https://github.com/julienschmidt/httprouter
* https://github.com/lib/pq
* https://golang.org/x/crypto/bcrypt
* https://github.com/dgrijalva/jwt-go
* https://github.com/go-yaml/yaml

# How To Run
The project uses Docker Compose

* `docker-compose up`

# APIs

### Note

Needs Auth Bearer Token in Headers.

* POST("/note") - Create Note
* GET("/note/:id") - Get a Note
* PUT("/note/:id") - Update a Note
* DELETE("/note/:id") - Delete a Note

### User

* POST("/user") - Create a User
* POST("/user/verify") - Verify and Generate a Auth Token