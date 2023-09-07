# Go-Gin-API

A comprehensive backend API built with Go and the Gin framework. This application provides CRUD operations for books, user authentication, logging, rate limiting, and more.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Authentication**: Secure user registration and login using JWT.
- **CRUD for Books**: Create, read, update, and delete operations for books.
- **Logging**: Logs user activities and system events.
- **Rate Limiting**: Prevents abuse by limiting the number of API requests.
- **Email Utilities**: Send emails for notifications or password resets.
- **Password Utilities**: Securely hash and verify passwords.

## Prerequisites

- Go (version 1.16 or newer)
- PostgreSQL
- SMTP server (for email utilities)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/astergginete/golang-backend-api.git
   cd golang-backend-api
   ```

2. Install the required Go packages:
   ```bash
   go mod download
   ```

3. Set up your PostgreSQL database and update the connection string in the configuration.

4. Run the database migrations:
   ```bash
   # Assuming you have a migration tool set up
   migration_tool run internal/migrations/
   ```

5. (Optional) Set up your SMTP server details in the configuration for email utilities.

## Usage

1. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

2. The server will start on the default port (e.g., `8080`). You can access the API at `http://localhost:8080`.

## API Endpoints

### User

- **Register**: `POST /users/register`
- **Login**: `POST /users/login`
- **Get User Profile**: `GET /users/profile`
- **Update User Profile**: `PUT /users/profile`
- **Delete User**: `DELETE /users/delete`

### Books

- **List All Books**: `GET /books`
- **Get Single Book**: `GET /books/:id`
- **Create Book**: `POST /books`
- **Update Book**: `PUT /books/:id`
- **Delete Book**: `DELETE /books/:id`

### Logs

- **Get All Logs**: `GET /logs`
- **Get Log by ID**: `GET /logs/:id`

(Add more endpoints as necessary)

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Push to your fork.
5. Submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
