# First-Go-Project  

A simple backend service for managing notes (CRUD API) written in Go.  
This project demonstrates fundamental principles of clean architecture, including separation of concerns between entities, services, repositories, and DTOs.

## Features:  

- Create, read, update, and delete notes
- Input validation
- In-memory data storage
- Basic events logging
- Unit tests for service and repository layers

## Technologies

- Go 1.24
- Echo web framework
- Go standard library (net/http, sync, log, etc.)
- Testify assert package


## API Endpoints

### Notes

| Method | Path           | Description             |
|--------|----------------|-------------------------|
| GET    | /notes         | Get all notes           |
| GET    | /notes/:id     | Get note by ID          |
| POST   | /notes         | Create a new note       |
| PUT    | /notes/:id     | Update an existing note |
| DELETE | /notes/:id     | Delete a note by ID     |

## How to run

Run the application from the project folder using:
```bash
go run main.go
```  
Or if you have `make` installed use:
```bash
make run
```

