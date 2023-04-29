# stay-sorted

Stay-Sorted is a web application that allows users to create and manage property listings, make reservations, and search for available properties. The application provides user authentication, property management, and reservation handling.

## Features
- User authentication: users can sign up, log in, and log out of the application
- Property management: users can create, update, and delete their property listings
- Property search: users can search for available properties based on location and amenities
- Reservation handling: users can make reservations for available properties and view their reservations
- Admin panel: an admin user can manage all the listings and reservations in the application

## Technologies
Stay-Sorted is built using the following technologies:
- Go: the programming language used to build the backend of the application
- Gin: a web framework for Go that provides routing, middleware, and handlers
- GORM: an Object-Relational Mapping (ORM) library for Go that provides a high-level interface for interacting with databases
- PostgreSQL: a relational database used to store the application data
- HTML/CSS/JavaScript: the languages used to build the frontend of the application

## Installation
To install and run Stay-Sorted on your local machine, follow these steps:

- Clone the repository: git clone https://github.com/vaibhavvikas/stay-sorted-go.git
- Install Go, PostgreSQL, and any other dependencies needed by the application
- Create a new PostgreSQL database for the application
- Copy the .env.example file to .env and update the database connection information
- Start the application: go run .
- The application should now be running at http://localhost:5000.

## API Endpoints
Stay-Sorted provides the following API endpoints:

- /api/users/signup: creates a new user account
- /api/users/login: logs in an existing user and returns a JWT token
- /api/users/:id: returns details for the specific user
- /api/houses/:id: returns details for a specific property
- /api/properties/search: searches for properties based on location and amenities
- /api/reservations: returns a list of all reservations made by the logged-in user
- /api/reservations/:id: returns details for a specific reservation made by the logged-in user

## License
Stay-Sorted is licensed under the MIT License. See the LICENSE file for more information.
