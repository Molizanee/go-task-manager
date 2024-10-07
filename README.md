# RESTful API For Task Manager App in Golang

## Stack used:
- Golang
- Librarys:
  - Chi
  - Pgx
- PostgreSQL

## Endpoints:
- `/tasks` | GET
  - Return all tasks
- `/task/{id}` | GET
  - Return a task by id
- `/task` | POST
  - Create a task with this schema
  ```json
  {
  "title": "Task Title",
  "description": "Task Description",
  "owner": "Task Owner",
  "status": "Task Status"
  }
  ```  
