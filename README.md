# Golang Clean

The project has been developed using the Test-Driven Development (TDD) methodology and structured according to the principles of Clean Architecture.

## Objective

In this project, our goal is to provide basic HTTP APIs for managing user data. The project has been developed in compliance with Clean Architecture principles, ensuring tight control over dependencies within internal layers while facilitating extensibility outward.

## Used Technologies and Architecture

The project has been developed using the following technologies and architectural principles:

- **Programming Language:** Go
- **HTTP Framework:** Fiber (Version 2)
- **Database:** MongoDB

## Project Architecture

The project follows the Clean Architecture principles and is structured into the following layers:

1. **Application Layer:** This layer handles HTTP requests and manages the business logic. It communicates with the Domain Layer.
2. **Domain Layer:** This layer defines the business logic. User objects, business rules, and fundamental rules are defined here.
3. **Infrastructure Layer:** This layer provides access to the database and other external resources. Dependencies are managed in this layer.
4. **Presentation Layer:** This layer contains routers that handle HTTP requests and user interface code.

## API Endpoints

- `GET /user`: Lists all users.
- `GET /user/:id`: Retrieves a specific user by ID.
- `POST /user`: Creates a new user.

## How to Run

1. Clone the project folder to your computer.
2. Open a terminal and navigate to the project folder: `cd project-folder`
3. Install the required dependencies: `go mod tidy`
4. To start the application: `go run main.go`
5. The application will run at [http://localhost:8080](http://localhost:8080).
