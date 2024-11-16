# Student CRUD API

## Overview
This is a simple student CRUD REST API built using Golang and Gin.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/InosRahul/student-crud-api.git
    cd student-crud-api
    ```

2. Run migrations:
    ```bash
    make migrate
    ```

3. Start the server:
    ```bash
    make run
    ```

## Environment Variables

- `PORT`: Port number to run the server on (default: 8080)

## Endpoints

- **GET /api/v1/students**: Get all students.
- **GET /api/v1/students/{id}**: Get a specific student by ID.
- **POST /api/v1/students**: Create a new student.
- **PUT /api/v1/students/{id}**: Update an existing student.
- **DELETE /api/v1/students/{id}**: Delete a student record.

## Running Tests

Run unit tests using:
```bash
make test
