## A Simple RESTful To-Do API in Go

Project demonstrates key features as:
1. **Basic CRUD Operations**: The API allows for the creation, reading, updating, and deletion of tasks.
2. **Database Integration**: SQLite with GORM for object-relational mapping
3. **Clear API Endpoints**: Well-defined endpoints for viewing tasks (all, active, completed), creating tasks, updating task status, and deleting tasks.
4. **Error Handling**: The API checks for errors, such as missing task IDs or invalid requests, and responds with appropriate status codes.

Links:
**All tasks**
> http://localhost:8080/tasks
**Completed tasks**
> http://localhost:8080/tasks/completed
**Active tasks**
> http://localhost:8080/tasks/active
**Task by ID number**
> http://localhost:8080/tasks/<ID Number>

Instructions:
**Add a new task**
use body.json file
Powershell:
> Invoke-RestMethod -Uri "http://localhost:8080/tasks" -Method Post -ContentType "application/json" -InFile "body.json"
**Mark as completed/uncompleted**
Powershell:
> Invoke-WebRequest -Uri "http://localhost:8080/update?id=<ID Number>" -Method PATCH
**Delete from the list**
Powershell:
> Invoke-WebRequest -Uri "http://localhost:8080/tasks/del/<ID Number>" -Method DELETE      


