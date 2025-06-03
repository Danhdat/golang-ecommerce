# Golang E-commerce Website

E-commerce website built with Golang, Gin framework, GORM, and MySQL.

## Features
- Product management (CRUD)
- Category management  
- Product search and filtering
- Shopping cart system
- User authentication
- Order processing
- Payment integration

## Tech Stack
- **Backend**: Golang with Gin framework
- **ORM**: GORM
- **Database**: MySQL
- **Architecture**: MVC pattern

## Setup
1. Clone repository
2. Install dependencies: `go mod download`
3. Setup MySQL database
4. Run: `go run main.go`

## API Endpoints
- GET /products - List all products
- GET /categories/:id - Products by category
- GET /products/search?q=keyword - Search products

## Project Structure