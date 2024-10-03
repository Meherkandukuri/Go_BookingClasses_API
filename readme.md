# Studio Classes API

This is a RESTful API built with Go that allows users to manage and book studio classes. The API provides endpoints for creating classes and bookings, with proper validation and error handling. 

## Features

- Create studio classes
- Book a class
- Input validation for requests
- Structured routing with Chi
- JSON-based responses for API interaction

## Project Structure

- **handlers**: Contains the HTTP handlers that manage the endpoints.
- **routes**: Defines the routes for the API.
- **helpers**: Utility functions for tasks such as JSON decoding, response writing, and validation.
- **models**: Contains the data models representing classes and bookings.

## Endpoints

| Method | Endpoint      | Description                     |
|--------|---------------|---------------------------------|
| POST   | /classes      | Create a new class              |
| POST   | /bookings     | Create a new booking            |

## Getting Started

### Prerequisites

Make sure you have Go installed. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/MeherKandukuri/studioClasses_API.git
   cd studioClasses_API

2. Install dependencies:

    ```bash
    git mod tidy

3. Run the application:

    ```bash
    git run main.go

## API Usage

### Create a Class

#### Endpoint: POST /Classes

#### Request Body:
```json
{
  "className": "Yoga",
  "startDate": "2024-10-01",
  "endDate": "2024-10-07",
  "capacity": 15
}
```

#### Response Body:
```json
{
  "message": "Created Yoga classes between 2024-10-01 and 2024-10-07 with Capacity: 15"
}
```


### Create a Booking

#### Endpoint: POST /bookings

#### Request Body:
```json
{
 "name": "Meher",
 "date": "2024-10-02"
}
```

#### Response Body:
```json
{
  "message": "Meher has been enrolled for class on 2024-10-02"
}
```
## Contribution Guidelines

We welcome contributions to improve the project! If you're interested in contributing, please follow the guidelines below:

1. Fork the Repository
Start by **forking** the repository to your own GitHub account. You can do this by clicking the "Fork" button at the top right of the repository page.

2. Clone the Repository
Once you've forked the repository, clone it to your local machine:

```bash
    git clone https://github.com/your-username/BookingClasses_API.git
```

3. Create a New Branch
Before starting any changes, create a new branch for your feature or bug fix.

4. Make sure your code follows go's  Standard conventions

5. Write Tests for your changes

6. Create a pull request and describe the chagnes you made with reasoning.





