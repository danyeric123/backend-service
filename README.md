# Go Backend Template

This is a basic template for a Go backend with PostgreSQL. It includes best practices and boilerplate code to get you started quickly.

## Project Structure

```markdown
go-backend-template/ 
  ├── cmd/ 
  │ └── main.go 
  ├── db/ 
  │ └── db.go 
  ├── handlers/ 
  │ └── handler.go 
  ├── Dockerfile 
  ├── docker-compose.yaml 
  ├── go.mod 
  ├── go.sum 
  ├── init.sql 
  ├── Makefile 
  └── README.md
```

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. Build the Docker images:

    ```sh
    make build
    ```

2. Run the application:

    ```sh
    make run
    ```

3. Access the application at `http://localhost:8080`.

### Cleaning Up

To stop and remove the Docker containers, run:

```sh
make clean
```

### Database Initialization

The init.sql file contains SQL statements to initialize the database with some boilerplate data. This file is automatically executed when the PostgreSQL container is started.