# Go User API

A RESTful API built using Go, Fiber, and PostgreSQL.

## Features

* Create User
* Get User By ID
* Get All Users
* Update User
* Delete User
* Dynamic Age Calculation
* Input Validation
* Structured Logging with Zap

## Tech Stack

* Go
* Fiber
* PostgreSQL
* Validator
* Zap Logger

## API Endpoints

### Create User

POST /users

### Get All Users

GET /users

### Get User By ID

GET /users/:id

### Update User

PUT /users/:id

### Delete User

DELETE /users/:id

## Run Project

go run cmd/server/main.go
