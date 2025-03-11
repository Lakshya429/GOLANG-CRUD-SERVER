# Movie API in Golang

This is a simple RESTful API built using Golang and the `gorilla/mux` router. It manages a list of movies, allowing users to create, read, update, and delete movie records.

## Features
- Retrieve all movies (`GET /movies`)
- Retrieve a single movie by ID (`GET /movies/{id}`)
- Add a new movie (`POST /movies`)
- Update an existing movie (`PUT /movies/{id}`)
- Delete a movie (`DELETE /movies/{id}`)

## Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/dl/)
- `gorilla/mux` package (`go get -u github.com/gorilla/mux`)

## Installation & Running the API
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/movie-api-go.git
   cd movie-api-go
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the application:
   ```sh
   go run main.go
   ```
4. The server will start on `http://localhost:8080`

## API Endpoints

### 1. Get All Movies
```http
GET /movies
```
Response:
```json
[
  {
    "id": "1",
    "title": "Legendary Lakshya From Dhapara",
    "director": {
      "FristName": "Sawam",
      "LastName": "Hum"
    },
    "isbn": "123"
  }
]
```

### 2. Get Movie by ID
```http
GET /movies/{id}
```
Response:
```json
{
  "id": "1",
  "title": "Legendary Lakshya From Dhapara",
  "director": {
    "FristName": "Sawam",
    "LastName": "Hum"
  },
  "isbn": "123"
}
```

### 3. Create a Movie
```http
POST /movies
```
Request Body:
```json
{
  "title": "New Movie",
  "director": {
    "FristName": "John",
    "LastName": "Doe"
  },
  "isbn": "456"
}
```
Response:
```json
[ ... updated list of movies ... ]
```

### 4. Update a Movie
```http
PUT /movies/{id}
```
Request Body:
```json
{
  "title": "Updated Movie",
  "director": {
    "FristName": "Jane",
    "LastName": "Doe"
  },
  "isbn": "789"
}
```
Response:
```json
{
  "id": "1",
  "title": "Updated Movie",
  "director": {
    "FristName": "Jane",
    "LastName": "Doe"
  },
  "isbn": "789"
}
```

### 5. Delete a Movie
```http
DELETE /movies/{id}
```
Response:
```json
{ "message": "Movie deleted successfully" }
```

## Notes
- The API uses `math/rand` to generate random IDs for new movies.
- Error handling should be improved for production use.

## Author
Your Name

## License
This project is licensed under the MIT License.

