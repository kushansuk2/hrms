# Human Resource Management System (HRMS)

This is a mini project developed in Golang using the gofr framework. It was created as part of an assignment for Zopsmart. HRMS is a system that manages employees with the following endpoints:

1. Get all employees
2. Create an employee
3. Update an employee
4. Delete an employee

## Prerequisites

-   Golang
-   Gofr framework

## Configuration

To configure the application, you need to setup the `.env` file. Follow the steps below:

1. Create a new file in the configs directory of the project and name it `.env`.
2. Open the `example.env` file and according to this update your crediantials in `.env`.

## Running the Application

Follow these steps to run the application:

1. Open a terminal in the root directory of the project.
2. Run the following command to start the application:

```
go mod tidy
go run main.go
```

The application should now be running and ready to accept requests at `localhost:8080`.

## API Documentation

For more details about the APIs and how to use them, check out the [API documentation](https://documenter.getpostman.com/view/23397038/2s9Ykn9hq5).
