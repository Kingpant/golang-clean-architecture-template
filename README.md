# Golang Clean Architecture Template

This is a template for building applications in Go using the Clean Architecture principles. It provides a structured way to organize your code, making it easier to maintain and test.

## Features

-   **Clean Architecture**: The project follows the Clean Architecture principles, separating concerns into different layers.
    -   **Domain Layer**: Contains the core business logic and domain entities.
    -   **Usecases Layer**: Contains application-specific business rules and orchestrates the flow of data between the domain and the interface layers.
    -   **Interface Layer**: Contains the interfaces for external systems, such as HTTP handlers, database repositories, and external APIs.
    -   **Infrastructure Layer**: Contains the implementation details for the interfaces defined in the Interface Layer, such as database connections and external service clients.

## Tools
This template uses several tools and libraries to facilitate development:
-   **Fiber**: A fast HTTP web framework for Go.
-   **Bun**: An ORM for Go that provides a simple and efficient way to interact with databases.
-   **GoDotEnv**: A library for loading environment variables from `.env` files.
-   **Zap**: A logging library for Go that provides structured logging.
-   **Swag**: A tool for generating API documentation from Go code.
-   **GoMock**: A mocking framework for Go that allows you to create mock implementations of interfaces for testing.


## Getting Started
0. Install the required tools:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   go install github.com/segmentio/goline@latest
   go install go.uber.org/mock/mockgen@latest
    ```
    Make sure to add the `$GOPATH/bin` directory to your `PATH` environment variable.
    - For example, you can add the following line to your `.bashrc` or `.zshrc` file:
      ```bash
      export PATH=$PATH:$GOPATH/bin
      ```
    - swag will be used to generate API documentation
    - goline will be used to format the code use with **emeraldwalk.runonsave** extension as you can see in the `.vscode/settings.json` file
    - mockgen will be used to generate mock implementations of interfaces for testing.

1. Clone the repository:
   ```bash
   git clone https://github.com/Kingpant/golang-clean-architecture-template.git
   ```
2. Change into the project directory:
   ```bash
   cd golang-clean-architecture-template
   ```
3. Install the dependencies:
   ```bash
   go mod tidy
   ```
4. Copy the `.env.example` file to `.env` and update the environment variables as needed:
   ```bash
   cp .env.example .env
   ```
5. Migrate the database:
   ```bash
   go run cmd/bun/main.go db init
   go run cmd/bun/main.go db migrate
   ```
6. Run the application:
   ```bash
   go run cmd/api/main.go
   ```
   or if you have specific env file:
   ```bash
   DOTENV_PATH={{YOUR_ENV_FILE_PATH}} go run cmd/api/main.go
   ```

## Running Tests
To run the tests, use the following command:
```bash
go test ./...
```
- Unit tests are written in the domain and usecases layers.