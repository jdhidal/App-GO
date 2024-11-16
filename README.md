# Project in Go 🌎

This is a basic program in Go that displays a "Hello World" message on the web. The project uses Docker for containerization and has also been deployed on Heroku using the `test` branch.

## Prerequisites

Before you begin, make sure you have **Go** installed on your system. You can download it from the official website:

[Download Go](https://golang.org/dl/)

## Project Cloning

To get started, clone this repository to your local machine using the following command:

```bash
git clone https://github.com/jdhidal/App-GO.git
```

## Run Locally

To run the project locally on your machine:

1. **Navigate to the project directory**:
```bash
cd App-GO
```

2. **Start the application**:
```bash
go run main.go
```
This will make the application available at http://localhost:8080

## Docker Hub

The image for this project is also available on Docker Hub, allowing you to run it without needing to build it locally. You can get it by running:

```bash
docker pull jdhidalgo673/app-go:latest
```

```bash
docker run -p 8080:8080 jdhidalgo673/app-go:latest
```
This will make the application available at http://localhost:8080

## Deployment to Heroku

This project is deployed on Heroku, so you can access the application directly at the following link:

[Visit the Heroku website](https://app-go-cad6f4cb6f2f.herokuapp.com/)


Thank you for exploring this Hello World project in Go! 😊