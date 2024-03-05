
# Task-Manager

Task Management System API: A RESTful API for managing tasks, assigning them to users, and updating their status. Built with Go and utilizing microservices architecture for scalability and flexibility
## Features

- User registration
- User authentication
- Add tasks to the system
- Retrieve tasks assigned to a specific user
- Retrieve tasks created by a specific user
- Update the status of tasks assigned to a user
- Update the status of tasks created by a user



## End Points

- `POST /postTask`: Add a task to the system.
- `GET /getTaskForUser/{username}`: Retrieve tasks assigned to a specific user.
- `GET /getTaskForCreatedUser/{username}`: Retrieve tasks created by a specific user.
- `PUT /updateStatus/{username}/{title}/{status}`: Update the status of a task assigned to a user.
- `PUT /updateStatusCreatedUser/{username}/{title}/{status}`: Update the status of a task created by a user.

## Tech Stack

**Client:** React, Typescript

**Server:** Go, Microservices

**Database:** MySql


## Usage

To use this API, you can send HTTP requests to the specified routes using tools like cURL, Postman, or any HTTP client library.


