
# DietSense

## Introduction

DietSense is an AI-powered nutrition analysis service that leverages advanced machine learning techniques and state-of-the-art language models to provide detailed insights into the nutritional content of foods. Users can upload images of their meals, and DietSense will analyze these images to identify the food items and provide a comprehensive breakdown of their nutritional contents, including calories, macronutrients, and micronutrients.

## Key Features

- **Image Processing**: Upload meal photos and receive detailed nutritional data.
- **Nutritional Breakdown**: Get insights into calories, macronutrients (proteins, carbs, fats), and micronutrients (vitamins, minerals).
- **Easy Integration**: Offers an API that can be integrated with other health and nutrition applications.

## Tech Stack

- **Programming Language**: Go (Golang)
- **Web Framework**: Gin
- **Database**: SQLite (default), with an option for PostgreSQL in production environments
- **ORM**: GORM
- **Authentication**: API Key-based
- **Rate Limiting**: Implemented using go-redis-rate
- **Logging**: Logrus
- **Configuration Management**: Viper
- **API Documentation**: Swagger
- **Testing**: Go Testing Package and Testify
- **Containerization**: Docker
- **Cloud Deployment**: Azure

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

```bash
docker
docker-compose
```

### Installing

A step by step series of examples that tell you how to get a development environment running:

1. Clone the repository

```bash
git clone https://github.com/codevalley/dietsense.git
cd dietsense
```

2. Build the Docker image

```bash
docker build -t dietsense .
```

3. Run the application using Docker Compose

```bash
docker-compose up
```

This command will start the service on `localhost:8080`. You can access the API endpoints using the specified port.

### Using the API

To analyze a food image for nutritional content, send a POST request to the endpoint:

```bash
curl -X POST "http://localhost:8080/analyze" -F "image=@path_to_your_food_image"
```

## Contribution

Please read [CONTRIBUTING.md](#) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/yourusername/dietsense/tags).

## Authors

* **Narayan Babu** - *Initial work* - [@codevalley](https://github.com/codevalley)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* To my wife and kids for allowing me spend my weekends on useless stuff.
