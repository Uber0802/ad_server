# ad_server

## Design Choices

### Framework and Libraries

- **Gorilla Mux:** Chosen for its  URL routing and dispatching, making it ideal for RESTful services.
- **GORM:** used for its compatibility with MySQL and features like AutoMigrate for easy schema updates.
- **Jsoniter:** used to replace the standard library's JSON encoding and decoding with faster processing.

### Database

- **MySQL:** Selected for its widespread use and support for JSON data types.

### Caching

- **In-Memory Caching with `sync.Map`:** implemented in-memory caching prevents repetitive processing for frequent, identical requests.

## Code Comments

Below is an overview of the critical components of the application, along with explanations for significant choices made in the codebase.

### main.go

The entry point of the application, setting up the HTTP server and routing.

### package/config/app.go

Establishes the database connection.

### package/controllers/controller.go

Contains handler functions for the API endpoints. It includes optimizations like JSON parsing with `jsoniter` for efficiency and a caching mechanism to speed up repeated condition checks.

### package/models/model.go

Defines the `Ad` and `Conditions` structs, mirroring the database schema.

### package/routes/route.go

Registers the API endpoints with the router, linking URL paths to their corresponding handler functions in `controllers`.

### package/utils/utils.go

Provides utility functions, such as `ParseBody`, which streamlines the parsing of JSON request bodies.

## Enhancements for Speed

To ensure the API serves requests rapidly, we focused on optimizing database interactions and processing speed. This includes using connection pooling, indexing database columns, caching frequent computations, and utilizing `jsoniter` for fast JSON processing.
