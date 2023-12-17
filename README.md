
# Task Scheduler App with GoFr Framework

This repository contains the source code for a CRUD (Create, Read, Update, Delete) API built in the Go language using the GoFr Framework. The API is designed for managing tasks in a task-scheduling application, providing endpoints for task creation, retrieval, updating, and deletion.

## Features

- **Task Creation:** Create new tasks with specified attributes.
- **Task Retrieval:** Retrieve task information by ID.
- **Task Update:** Update task attributes, including title, description, priority, and due date.
- **Task Deletion:** Soft-delete tasks using the is_shadowed field.
- **DB Integration:** Utilizes SQL as the database with the specified schema hosted on a Docker image.

## Prerequisites

Ensure you have the following prerequisites before setting up the project:

- Go installed on your machine
- [GoFr Framework](https://gofr.dev/docs) (you can install it with `go get gofr.dev`)
- SQL Database (MySQL) with appropriate configurations
- Docker MySQL Image (freely available Docker image)
- Postman for testing API endpoints

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/ImRi8/task-Scheduler-App
   ```

2. Navigate to the project directory:

   ```bash
   cd task-Scheduler-App
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Set up the database:

   Run the MySQL server and create a database locally using the following Docker command:

   ```
   docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=taskDB -p 3306:3306 -d mysql
   ```
   ```
   docker exec -it gofr-mysql mysql -uroot -proot taskDB -e "CREATE TABLE task (Id INT AUTO_INCREMENT PRIMARY KEY,is_shadowed BOOLEAN,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,title VARCHAR(255),description VARCHAR(255),priority BIGINT,due_date TIMESTAMP);"
   ```
**If the above command does not work then, Create a DataBase named taskDB and create a table with below schema (Follow the sqlCommandFile.txt)**

5. Enter your MySQL credentials in `config/.env`:

   ```env
   APP_NAME=task-scheduler
   HTTP_PORT=8000
   APP_VERSION=1

   DB_HOST=localhost
   DB_USER=root
   DB_PASSWORD=root
   DB_NAME=taskDB
   DB_PORT=3306
   DB_DIALECT=mysql
   ```

6. Run the application:

   ```bash
   go run main.go
   ```

## API Endpoints

### Check Health

- **Endpoint:** `/task/health`
- **Method:** `GET`

### Create Task

- **Endpoint:** `/task/createTask`
- **Method:** `POST`
- **Request Payload:**

  ```json
  {
   "title": "Task Title",
   "description": "Task Description",
   "priority": 3,
   "dueDate": "2023-12-31T23:59:59Z"
  }
  ```

### Update Task

- **Endpoint:** `/task/updateTask`
- **Method:** `POST`
- **Request Payload:**

  ```json
  {
     "id" : 5,
     "title": "new title",
     "description" : "new update"
  }
  ```

### Retrieve Task

- **Endpoint:** `/tasks/getTaskById?id={id}`
- **Method:** `GET`

### Delete Task

- **Endpoint:** `/tasks/deleteTaskById?id={id}`
- **Method:** `GET`

## Data Schema

- `id` (int): Auto-updating identifier.
- `is_shadowed` (bool): Soft deletion flag.
- `title` (string): Task title.
- `description` (string): Task description.
- `priority` (bigint): Task priority (1 to 5).
- `due_date` (time): Task due date (cannot be before the current timestamp).


## Contact

Rishabh Kumar Singh - singh.rishabh701@gmail.com
