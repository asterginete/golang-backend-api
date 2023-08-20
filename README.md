# Go-Gin Book API

A simple RESTful API for managing books, built with Go and the Gin framework.

## Getting Started

### Prerequisites

- Go (Download and install from [here](https://golang.org/dl/))
- Gin framework (`go get -u github.com/gin-gonic/gin`)

### Running the API

Navigate to the project directory and run:

```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### CRUD Operations

1. **Create a Book**
   - **Endpoint**: `/books`
   - **Method**: `POST`
   - **Body**:
     ```json
     {
       "title": "Book Title",
       "author": "Author Name",
       "genre": "Genre",
       "publication_date": "YYYY-MM-DD"
     }
     ```

2. **List All Books**
   - **Endpoint**: `/books`
   - **Method**: `GET`

3. **Get a Single Book**
   - **Endpoint**: `/books/:id`
   - **Method**: `GET`

4. **Update a Book**
   - **Endpoint**: `/books/:id`
   - **Method**: `PUT`
   - **Body**:
     ```json
     {
       "title": "Updated Title",
       "author": "Updated Author",
       "genre": "Updated Genre",
       "publication_date": "Updated YYYY-MM-DD"
     }
     ```

5. **Delete a Book**
   - **Endpoint**: `/books/:id`
   - **Method**: `DELETE`

### Additional Endpoints

6. **Get Books by Author**
   - **Endpoint**: `/books/author/:name`
   - **Method**: `GET`

7. **Search Books by Title or Author**
   - **Endpoint**: `/books/search/:query`
   - **Method**: `GET`

8. **Filter Books by Genre**
   - **Endpoint**: `/books/genre/:genre`
   - **Method**: `GET`

9. **Sort Books by Publication Date**
   - **Endpoint**: `/books/sort/date`
   - **Method**: `GET`

10. **Sort Books by Title**
    - **Endpoint**: `/books/sort/title`
    - **Method**: `GET`

11. **Filter Books by Publication Year**
    - **Endpoint**: `/books/year/:year`
    - **Method**: `GET`

## Testing

You can test the API using tools like `curl`, Postman, or directly in the browser (for GET requests).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.
