
# SE Backend Project

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Testing](#testing)

## Prerequisites

Before you begin, ensure you have met the following requirements:
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install)

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repo>
   ```

2. **Fetch the latest changes:**
   ```bash
   git fetch
   ```

3. **Checkout the develop branch:**
   ```bash
   git checkout develop
   ```
4. **Update dependencys:**
   ```bash
   go mod tidy
   ```

5. **Copy the example environment configuration:**
   ```bash
   cp .env.example .env
   ```

6. **Edit the `.env` file:**
   - Change the necessary keys according to your environment.

## Running the Application

1. **Start the application with Docker Compose:**
   ```bash
   docker-compose up
   ```

2. **Open pgAdmin** to manage your PostgreSQL database.

3. **Run the application:**
   ```bash
   go run main.go
   ```

## Testing

You can test the API endpoints using [Postman](https://www.postman.com/):
- Set up your requests based on the API specifications.
- Ensure the application is running before testing.

```bash
# Example command or API request
curl -X GET "http://localhost:8000/user"
```
