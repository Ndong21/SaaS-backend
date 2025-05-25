# SaaS Backend
This repository contains code for creating backend endpoints for a SaaS management system for small and medium-sized businesses. The APIs are developed using SQLC and the Gin web framework in Go.

## Project Structure
`cmd/api`: This folder holds the application's main entry point.

`db/migrations`: Contains SQL files that create/update the database schema (tables and columns) used by the APIs.

`internal/stocks/handler`: Contain API route definitions and handler functions for the stocks module.

`internal/stocks/query`: This folder contain sql query files. These files define the database queries used by the application, which are processed by sqlc to generate type-safe Go code for interacting with the database.

`internal/stocks/repo`: This directory contains repository code that acts as an abstraction layer between the database
and the application logic. It provides functions to interact with database using generated sqlc code.

## Setting up the project
- Set up environment variables:
  - Create a new database for use in this project
  - Configure the neccessary variables in the .env file such as the database connection details
- Install Dependencies:
  -Ensure you have Go, golang-migrate and sqlc installed. Then install the required Go modules: `go mod tidy`

- Run the application: `go run cmd/api/...`



