# Go Spring CRUD API with File-Based Database

A RESTful API built with Go that provides CRUD operations for articles using persistent file storage as the database.
I did step 1 and 2 in the assignment

## Features

- **Create** new articles (POST)
- **Read** all articles or single article (GET)
- **Update** existing articles (PUT)
- **Delete** articles (DELETE)
- **Persistent file storage** using Go's gob encoding
- **Thread-safe** operations with mutex locks
- **Automatic data persistence** - no manual save needed
- JSON responses
- Proper error handling
- Timestamps for created/updated

## Database

This API uses a **file-based database** system that:

- Stores data in `articles.gob` file using Go's gob encoding
- Automatically loads existing data on startup
- Persists changes immediately to disk
- Creates sample data if no existing data is found
- Is thread-safe for concurrent operations

**Benefits:**

- ✅ No external database installation required
- ✅ Data persists between server restarts
- ✅ Simple setup and deployment
- ✅ Perfect for development and small applications
- ✅ No authentication or connection issues

## Prerequisites

1. **Go** (already installed)
2. **No database installation needed!** ✨

## API Endpoints

| Method | Endpoint         | Description              |
| ------ | ---------------- | ------------------------ |
| GET    | `/`              | Welcome message          |
| GET    | `/articles`      | Get all articles         |
| GET    | `/articles/{id}` | Get single article by ID |
| POST   | `/articles`      | Create new article       |
| PUT    | `/articles/{id}` | Update article by ID     |
| DELETE | `/articles/{id}` | Delete article by ID     |

## Running the Application

1. Run the application:
   ```powershell
   & go run main.go
   ```
2. The server will start on `http://localhost:8080`
3. Sample articles are automatically created on first run
4. Data is automatically saved to `articles.gob` file

## API Usage Examples

### Create a new article (POST)

```powershell
$body = @{
    title = "My First Article"
    desc = "This is a sample article"
    content = "This is the content of my first article."
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/articles" -Method POST -Body $body -ContentType "application/json"
```

### Get all articles (GET)

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/articles" -Method GET
```

### Get single article (GET)

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/articles/1" -Method GET
```

### Update an article (PUT)

```powershell
$body = @{
    title = "Updated Article Title"
    desc = "Updated description"
    content = "Updated content."
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/articles/1" -Method PUT -Body $body -ContentType "application/json"
```

### Delete an article (DELETE)

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/articles/1" -Method DELETE
```

## Response Format

All responses follow this JSON structure:

```json
{
  "message": "Operation result message",
  "data": {}, // Present for successful GET/POST/PUT operations
  "error": "" // Present only when there's an error
}
```

## Article Model

```json
{
  "id": 1,
  "title": "string",
  "desc": "string",
  "content": "string",
  "created": "2025-10-05T21:23:34.123456+03:00",
  "updated": "2025-10-05T21:23:34.123456+03:00"
}
```

## File Structure

```
go-spring/
├── main.go          # Main application code
├── articles.gob     # Database file (auto-created)
├── go.mod          # Go module file
├── go.sum          # Dependencies
└── README.md       # This file
```

## How It Works

1. **Startup**: App checks for existing `articles.gob` file
2. **Load Data**: If file exists, loads articles; otherwise creates sample data
3. **CRUD Operations**: All operations are thread-safe with mutex locks
4. **Auto-Save**: Changes are automatically saved to file after each operation
5. **Persistence**: Data survives server restarts

## Dependencies

- `github.com/gorilla/mux` - HTTP router and URL matcher
- Go standard library for file operations and JSON handling

## Development Benefits

✅ **No Database Setup** - Just run and go!  
✅ **Persistent Storage** - Data survives restarts  
✅ **Simple Deployment** - Single binary + data file  
✅ **Perfect for Learning** - Focus on API logic, not database config  
✅ **Thread-Safe** - Handles concurrent requests safely

## Production Considerations

For production use, consider upgrading to:

- **PostgreSQL** for larger datasets and complex queries
- **MariaDB/MySQL** for relational data requirements
- **MongoDB** for document-based storage
- **Redis** for caching and session storage

The current file-based approach is perfect for:

- Development and testing
- Small applications (< 10,000 records)
- Prototyping and demos
- Applications with simple data requirements

## CRUD Operations Summary

✅ **CREATE** - POST /articles - Add new articles  
✅ **READ** - GET /articles, GET /articles/{id} - Retrieve articles  
✅ **UPDATE** - PUT /articles/{id} - Modify existing articles  
✅ **DELETE** - DELETE /articles/{id} - Remove articles

🎉 **Your CRUD API is ready to use!**
