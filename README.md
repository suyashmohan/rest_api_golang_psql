# What is this?

This repository contains code to demonstrate a basic REST API written in Go language and PostgreSQL.

The APIs lets you create and save `Notes`.

# Requirements
* https://github.com/julienschmidt/httprouter
* https://github.com/lib/pq
* https://golang.org/x/crypto/bcrypt
* https://github.com/gbrlsnchs/jwt

# How To Run
The project has a basic Makefile to help out in running the project

* `make dep` should install the dependencies in local vendor folder
* `make run` should start the server

# APIs

### Note

Needs Auth Token in Headers.

* POST("/note") - Create Note
* GET("/note/:id") - Get a Note
* PUT("/note/:id") - Update a Note
* DELETE("/note/:id") - Delete a Note

### User

* POST("/user") - Create a User
* POST("/user/verify") - Verify and Generate a Auth Token