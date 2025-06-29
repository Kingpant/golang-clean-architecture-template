# Golang Clean Architecture Template

This is a template for building applications in Go using the Clean Architecture principles. It provides a structured way to organize your code, making it easier to maintain and test.

## Features

-   **Domain Layer**: Contains the core business logic and domain entities.
-   **Use Cases Layer**: Contains application-specific business rules and orchestrates the flow of data between the domain and the interface layers.
-   **Interface Layer**: Contains the interfaces for external systems, such as HTTP handlers, database repositories, and external APIs.
-   **Infrastructure Layer**: Contains the implementation details for the interfaces defined in the Interface Layer, such as database connections and external service clients.

## Tools

-   **Fiber**: A fast HTTP web framework for Go.
-   **Bun**: An ORM for Go that provides a simple and efficient way to interact with databases.
-   **GoDotEnv**: A library for loading environment variables from `.env` files.
