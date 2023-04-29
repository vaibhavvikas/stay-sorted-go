# stay-sorted

Stay-Sorted is a web application that allows users to create and manage property listings, make reservations, and search for available properties. The application provides user authentication, property management, and reservation handling.

## Features

### Completed
- User authentication: users can sign up, log in, and log out of the application
- Property management: users can create, update, and delete their property listings

### TODO:
- Property search: users can search for available properties based on location and amenities
- Reservation handling: users can make reservations for available properties and view their reservations
- Admin panel: an admin user can manage all the listings and reservations in the application

## Technologies
Stay-Sorted is built using the following technologies:
- Go: the programming language used to build the backend of the application
- Gin: a web framework for Go that provides routing, middleware, and handlers
- GORM: an Object-Relational Mapping (ORM) library for Go that provides a high-level interface for interacting with databases
- PostgreSQL: a relational database used to store the application data

## Project Features

- Proper project structure following Golang Standards
- Use of middlewares for proper authentication, cors and request handling
- Use of proper error messages and error handling using the middlewares
- Separating the json Request and Response objects from the entities object to show only the required results to users
- Code is divided into controllers with each controller handling validators, service, and transformers to sepearate the business logic from code logic
- Using .env file to handle various environment secrets like database credentials, JWT secret keys
- Using utils for common functions

## Installation
To install and run Stay-Sorted on your local machine, follow these steps:

- Clone the repository: git clone https://github.com/vaibhavvikas/stay-sorted-go.git
- Install Go, PostgreSQL, and any other dependencies needed by the application
- Create a new PostgreSQL database for the application
- Update the database connection information in .env file
- Start the application: go run .
- The application should now be running at http://localhost:5000.

## API Endpoints
Stay-Sorted provides the following API endpoints:

### Completed
- /api/users/signup: creates a new user account
- /api/users/login: logs in an existing user and returns a JWT token
- /api/users/:user_pid: returns details for the specific user
- /api/houses/create: create a listing for a specific property
- /api/houses/:house_pid: returns details for a specific property (requires used to be logged in)
- /api/houses: returns a list of all houses.

### TODO:
- /api/houses/search: searches for properties based on location and amenities
- Reservations related api

## License
Stay-Sorted is licensed under the MIT License. See the LICENSE file for more information.
