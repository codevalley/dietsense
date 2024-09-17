**Introduction**

DietSense is an AI-powered nutrition analysis service that allows users to upload images of their meals and receive detailed nutritional information. The service leverages advanced machine learning models, such as OpenAI's GPT-4 and Anthropic's Claude, to analyze food images and provide a comprehensive breakdown of nutritional contents, including calories, macronutrients, and micronutrients.

---

## What Does This Project Do?

DietSense serves as an API service where users can:

- **Upload Food Images**: Users send images of food to the service.
- **Receive Nutritional Analysis**: The service analyzes the images using AI models and returns nutritional data.
- **Integrate with Other Applications**: Provides endpoints that can be integrated into other health and nutrition applications.

The core functionality involves processing an image of food, estimating its nutritional content, and returning the results in a structured JSON format.

---

## Features Available

### 1. **Multiple AI Service Support**

- **Service Types**: Supports different AI services like OpenAI's GPT-4, Anthropic's Claude, and a mock service for testing.
- **Dynamic Service Selection**: Users can specify which AI service to use via a parameter in the API request.

### 2. **Image Analysis Endpoint**

- **Endpoint**: `/api/v1/analyze`
- **Functionality**: Accepts an image file and optional context, processes it using the selected AI service, and returns nutritional information.

### 3. **API Key Generation**

- **Endpoint**: `/api/v1/generate-api-key`
- **Functionality**: Generates API keys for clients to authenticate and authorize access to the service.
- **Access Control**: Protected by IP-based middleware to restrict access to authorized IP addresses.

### 4. **Configuration Management**

- **Flexible Configuration**: Uses a `config.yaml` file for setting up server address, environment, database URL, API keys, and other settings.
- **Context String**: Customizable prompt for AI services to guide the analysis process.

### 5. **Database Support**

- **Database Types**: Supports SQLite and PostgreSQL for storing API keys and potentially user data.
- **ORM**: Utilizes GORM for database interactions.
- **Models**: Defines models for API keys, user configurations, nutrition details, and usage statistics.

### 6. **Middleware**

- **Logging**: Uses Logrus for logging and integrates with Gin's middleware for request logging.
- **IP Restriction**: Middleware to restrict access to certain endpoints based on client IP addresses.

### 7. **Docker Support**

- **Dockerfile**: Provided for containerizing the application.
- **Docker Compose**: Sample `docker-compose.yml` for running the service in a Docker environment.

### 8. **Testing**

- **Unit Tests**: Includes basic tests for the API endpoints using Go's testing package and Testify.

---

## Features That Can Be Added

### 1. **API Key Authentication Middleware**

- **Description**: Implement middleware to authenticate requests using API keys, ensuring that only authorized users can access the service.
- **Benefit**: Enhances security and enables usage tracking per API key.

### 2. **Rate Limiting**

- **Description**: Implement rate limiting to prevent abuse of the API, potentially using Redis or in-memory solutions.
- **Benefit**: Ensures fair usage and protects the service from being overwhelmed by excessive requests.

### 3. **Data Persistence for Nutrition Details**

- **Description**: Save analyzed nutrition details to the database for future reference and user-specific insights.
- **Benefit**: Allows users to track their nutritional intake over time and enhances the value of the service.

### 4. **User Authentication and Management**

- **Description**: Introduce user accounts, authentication, and personalized settings.
- **Benefit**: Enables personalized experiences, such as saving preferences and tracking user-specific data.

### 5. **Enhanced Error Handling and Validation**

- **Description**: Improve error messages and input validation to provide better feedback to clients.
- **Benefit**: Makes the API more robust and user-friendly.

### 6. **Support for Additional AI Services**

- **Description**: Integrate more AI models or services for analysis, such as custom models or other providers.
- **Benefit**: Provides flexibility and potentially better performance or cost options.

### 7. **Front-end Interface**

- **Description**: Develop a web interface for users to interact with the service without needing to integrate the API.
- **Benefit**: Makes the service accessible to non-developers and can serve as a demo or testing tool.

### 8. **Automated Documentation (Swagger)**

- **Description**: Implement API documentation using Swagger or similar tools.
- **Benefit**: Provides clear documentation for developers intending to integrate the API.

---

## How to Invoke the Service

### Prerequisites

- **Environment Setup**: Ensure that you have Docker and Docker Compose installed.
- **Configuration File**: Create a `config.yaml` file based on `config.sample.yaml` with your specific settings and API keys.

### Starting the Service

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/codevalley/dietsense.git
   cd dietsense
   ```

2. **Build the Docker Image**:

   ```bash
   docker build -t dietsense .
   ```

3. **Run the Service with Docker Compose**:

   ```bash
   docker-compose up
   ```

   The service will start on `localhost:8080`.

### Generating an API Key

To generate an API key, you need to access the `/api/v1/generate-api-key` endpoint. This endpoint is protected by IP-based middleware, meaning only requests from allowed IP addresses can access it.

1. **Configure Allowed IPs**:

   - In your `config.yaml`, set the `allowed_ips` field to include your IP address.

     ```yaml
     allowed_ips: "127.0.0.1,::1,your_allowed_ip"
     ```

2. **Make a Request to Generate an API Key**:

   ```bash
   curl -X POST "http://localhost:8080/api/v1/generate-api-key" \
     -H "Content-Type: application/json" \
     -d '{
       "email": "your_email@example.com",
       "rate_limit_per_hour": 100
     }'
   ```

   - **Response**:

     ```json
     {
       "api_key": "generated-api-key"
     }
     ```

### Analyzing a Food Image

Once you have an API key, you can invoke the `/api/v1/analyze` endpoint.

1. **Prepare Your Image**:

   - Have the image file you want to analyze ready on your local machine.

2. **Make the API Request**:

   ```bash
   curl -X POST "http://localhost:8080/api/v1/analyze" \
     -H "Content-Type: multipart/form-data" \
     -H "Authorization: Bearer your-api-key" \
     -F "image=@/path/to/your/food_image.jpg" \
     -F "context=Additional context if any"
   ```

   - **Parameters**:
     - `image`: The image file of the food.
     - `context` (optional): Any additional context or description about the food.
     - `service` (optional): Specify the AI service to use (`openai`, `claude`, `mock`).

3. **Specify AI Service (Optional)**:

   - To use a specific AI service, include the `service` parameter:

     ```bash
     -F "service=openai"
     ```

4. **Response**:

   - The service will return a JSON response containing nutritional information.

     ```json
     {
       "summary": "A healthy salad with mixed greens, tomatoes, and a light dressing.",
       "nutrition": [
         {
           "component": "Calories",
           "value": 150,
           "unit": "kcal",
           "confidence": 0.8
         },
         {
           "component": "Protein",
           "value": 5,
           "unit": "g",
           "confidence": 0.7
         },
         // Additional nutritional components...
       ],
       "service": "openai"
     }
     ```

### Example Request Using cURL

```bash
curl -X POST "http://localhost:8080/api/v1/analyze" \
  -H "Authorization: Bearer your-api-key" \
  -F "image=@/path/to/your/food_image.jpg" \
  -F "context=Grilled chicken with vegetables" \
  -F "service=claude"
```

---

## Additional Notes

- **API Key Usage**: The API key should be kept secure and included in the `Authorization` header of your requests.
- **IP Restrictions**: Ensure your IP is included in the `allowed_ips` in the configuration to access restricted endpoints.
- **Configuration File**: Sensitive information like API keys for OpenAI and Claude should be securely stored and not checked into version control.
- **Testing**: The service includes a mock service (`service=mock`) for testing without incurring costs from AI providers.
- **Database Migrations**: The application automatically migrates database schemas using GORM when it starts.
