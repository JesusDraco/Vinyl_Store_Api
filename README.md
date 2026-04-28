# Vinyl Store API (Golang + Gin)

## Project Description

This project consists of a REST API developed in Golang using the Gin Framework.  
The API simulates a Vinyl Store inventory system similar to Merchbar.

Users must log in using **Basic Authentication** to receive an access token.  
This token is required to access protected endpoints.

The API provides functionality for authentication, album management, and system monitoring.

---

## Technologies Used

- Golang
- Gin Framework
- JSON
- Token-based Authentication
- Postman / cURL for testing

---

## How to Run the Project

### 1. Clone the repository

```bash
git clone https://github.com/JesusDraco/Vinyl_Store_Api.git
cd vinyl-store-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run .
```

### 4. Server URL

```
http://localhost:8080
```

---

## Test Users

| Username | Password |
|----------|----------|
| admin    | 1234     |
| user     | pass     |

---

## Authentication Flow

1. User logs in using **Basic Auth**.
2. Server returns a token.
3. User must include the token in all protected requests:

```
Authorization: Bearer <ACCESS_TOKEN>
```

---

## API Endpoints

### Login

```
GET /login
```

**Authentication:** Basic Auth

```bash
curl -u admin:1234 http://localhost:8080/login
```

**Response:**

```json
{
  "message": "Hi admin, welcome to the Store System",
  "token": "generated_token"
}
```

---

### Logout

```
GET /logout
```

**Authentication:** Bearer Token

```bash
curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/logout
```

---

### Get All Albums

```
GET /albums
```

**Authentication:** Bearer Token

---

### Get Album by ID

```
GET /albums/:id
```

**Authentication:** Bearer Token

---

### Create Album

```
POST /post-album
```

**Authentication:** Bearer Token  
**Content-Type:** application/json

**Body example:**

```json
{
  "id": "4",
  "title": "Benny Golson New York Scene",
  "artist": "Benny Golson",
  "price": 49.99
}
```

---

### Status

```
GET /status
```

**Authentication:** Bearer Token

---

## Using Postman

### 1. Login

- **Method:** GET
- **URL:** `http://localhost:8080/login`
- **Authorization:**
  - Type: Basic Auth
  - Username: `admin`
  - Password: `1234`

Copy the token from the response.

### 2. Access Protected Endpoints

For all protected endpoints:
- Go to **Authorization**
- Select: **Bearer Token**
- Paste the token

### 3. Get All Albums

- **Method:** GET
- **URL:** `http://localhost:8080/albums`

### 4. Get Album by ID

- **Method:** GET
- **URL:** `http://localhost:8080/albums/1`

### 5. Create Album

- **Method:** POST
- **URL:** `http://localhost:8080/post-album`
- **Authorization:** Bearer Token
- **Body → raw → JSON:**

```json
{
  "id": "4",
  "title": "New Album",
  "artist": "Artist",
  "price": 49.99
}
```

### 6. Logout

- **Method:** GET
- **URL:** `http://localhost:8080/logout`

### 7. Status

- **Method:** GET
- **URL:** `http://localhost:8080/status`

---

## Error Handling

The API includes error handling for:

- Missing Authorization header
- Invalid or expired token
- Invalid username or password
- Album not found (404)
- Invalid JSON format
- Empty fields
- Price <= 0
- Duplicate album IDs

---

## Project Features

- Server starts without issues
- Multiple users can login simultaneously
- Users can login to the API
- Users can logout from the API
- Users can get all albums
- Users can get album by ID
- Users can create new albums
- Users can get system status
- Duplicate album IDs are prevented
- Error handling implemented for all endpoints

---

## Future Work

- Add database persistence (MySQL or PostgreSQL)
- Add update and delete album endpoints
- Implement user roles (admin / client)
- Add inventory stock management
- Add logging and monitoring

---

## Author

`Jesus Abel Gutierrez Calvillo` 
`Jose Salcedo Uribe`
