
# Technical Architecture Document

## Overview

**DietSense** is a Go-based API service designed to analyze food images and provide nutritional information using AI services like OpenAI's GPT-4 and Anthropic's Claude. The application is built with modularity and scalability in mind, following clean architecture principles.

## Architecture Diagram

```
+---------------------------+
|       Clients (API)       |
| (Mobile Apps, Web Apps)   |
+------------+--------------+
             |
             v
+---------------------------+
|       API Gateway         |
|   (Gin Web Framework)     |
+------------+--------------+
             |
             v
+---------------------------+
|        Handlers           |
| (AnalyzeHandler, etc.)    |
+------------+--------------+
             |
             v
+---------------------------+
|         Services          |
| (Service Factory, AI      |
|  Service Implementations) |
+------------+--------------+
             |
             v
+---------------------------+
|       Repositories        |
| (Database Access Layer)   |
+------------+--------------+
             |
             v
+---------------------------+
|         Database          |
|  (SQLite/PostgreSQL)      |
+---------------------------+
```

## Components

### 1. **Clients**

- **Description**: External entities that interact with the API, such as mobile apps or web applications.
- **Interaction**: Send HTTP requests to the API endpoints.

### 2. **API Gateway (Gin Web Framework)**

- **Description**: The main entry point of the application, handling incoming HTTP requests.
- **Components**:
  - **Routes**: Defined in `api/routes.go`.
  - **Middleware**: Logging, authentication, and IP restriction.

### 3. **Handlers**

- **Description**: Functions that process requests and generate responses.
- **Key Handlers**:
  - **AnalyzeHandler**: Processes image analysis requests.
  - **APIKeyHandler**: Manages API key generation.

### 4. **Services**

- **Description**: Business logic layer that handles the core functionality.
- **Components**:
  - **Service Factory**: Creates instances of AI services based on configuration or request parameters.
  - **AI Service Implementations**: `OpenAIService`, `ClaudeService`, `MockService`.
- **Interaction**: Handlers invoke services to perform operations.

### 5. **Repositories**

- **Description**: Data access layer responsible for interacting with the database.
- **Components**:
  - **Database Interface**: Defines methods for data operations.
  - **Implementations**: PostgreSQL and SQLite implementations.
- **Interaction**: Services use repositories to persist or retrieve data.

### 6. **Database**

- **Description**: Stores persistent data such as API keys, user configurations, and analysis results.
- **Supported Databases**:
  - **SQLite**: Default for development and testing.
  - **PostgreSQL**: Recommended for production environments.

## Data Flow

1. **Client Request**:

   - A client sends a request to the API endpoint (e.g., `/api/v1/analyze`), including any required data such as an image file and API key.

2. **API Gateway**:

   - The Gin framework routes the request to the appropriate handler.
   - Middleware processes the request for logging and authentication.

3. **Handler Processing**:

   - The handler extracts data from the request.
   - It determines which AI service to use based on request parameters or configuration.

4. **Service Invocation**:

   - The handler calls the service factory to get an instance of the requested AI service.
   - The service processes the image and context, interacting with external AI APIs.

5. **AI Service Interaction**:

   - The AI service sends the image and context to the external AI API (e.g., OpenAI or Anthropic).
   - It receives the response and parses it into a usable format.

6. **Data Persistence (Optional)**:

   - The service may save analysis results to the database via the repository layer.
   - Usage statistics and API key usage may be updated.

7. **Response Generation**:

   - The handler constructs the response based on the service output.
   - It sends the response back to the client.

## Configuration Management

- **Viper**: Used for configuration management.
- **Configuration File**: `config.yaml` contains settings for server address, environment, database URL, API keys, etc.
- **Environment Variables**: Can override configuration settings.

## Logging

- **Logrus**: Used for logging throughout the application.
- **Middleware**: Custom logging middleware integrates with Gin to log requests and responses.

## Security

- **API Key Authentication**: Planned to be implemented to secure endpoints.
- **IP Restriction Middleware**: Currently restricts access to certain endpoints based on client IP.
- **Sensitive Data Handling**: API keys and credentials are managed securely and not exposed.

## Error Handling

- **Centralized Error Handling**: Handlers return consistent error responses.
- **Logging of Errors**: Errors are logged for monitoring and debugging purposes.

## Extensibility

- **Service Factory Pattern**: Makes it easy to add new AI services.
- **Modular Architecture**: Components are decoupled, allowing for independent development and testing.

## Future Enhancements

- **User Authentication**: Introduction of user accounts and authentication mechanisms.
- **Rate Limiting**: Implementation of rate limiting per API key.
- **Front-end Development**: Building a web interface for user interaction.

---
