# 📚 Go Todo API — Book Title Search by Author's name
This is a simple RESTful API built in Go that allows you to:
- Create todo items (e.g., book titles)
- Search todos by author name
- View all todos
- Delete todos

It uses:
- **Go’s built-in net/http package**
- **MySQL database**
- **Modular project structure** for clean separation of logic

---

## 🚀 Features

| Method | Endpoint     | Description                     |
|--------|--------------|---------------------------------|
| GET    | `/ping`      | Health check – returns "pong"   |
| POST   | `/`          | Create a new todo (book)        |
| GET    | `/?name=...` | Search todos by author name     |
| GET    | `/all`       | Fetch all todos               |
| DELETE | `/{name}`    | Delete a todo by author name    |

---

## 🧱 Project Structure
go-todo-api/
│
├── main.go # App entry point
├── go.mod / go.sum # Go modules
│
├── controller/ # HTTP handlers
│ ├── route.go
│ ├── ping.go
│ └── crud.go
│
├── model/ # Database logic
│ ├── connect.go
│ ├── create.go
│ └── read.go
│
└── views/ # Shared structs
└── struct.go


---

## Database Setup (MySQL)

1. Create a MySQL database:
   ```sql
   CREATE DATABASE GoDataBase;
2. Create the table
CREATE TABLE TodoList (
    name VARCHAR(100),
    todo TEXT
);

3. Update DB credentials in model/connect.go:
db, err := sql.Open("mysql", "root:yourpassword@tcp(localhost:3306)/GoDataBase")


## Installation

### Clone this repo
git clone https://github.com/your-username/go-todo-api.git

### Navigate to project
cd go-todo-api

### Get dependencies
go mod tidy

### Run the server
go run main.go

2. Build the project:
    ```bash
    go build
    ```
3. Run the application:
    ```bash
    ./go-todo-api
    ```
### Example Request
Open browser or API tool like Postman to test:Juma
→ http://localhost:3002/ping

### Sample POST Request (JSON)
POST /
{
  "name": "Silas Odero",
  "todo": "The Art of Go Programming"
}

## Author
. Silas Odero

. Go Developer | AWS Cloud Practitioner

## License
MIT License – feel free to use, fork, and build on this project!
