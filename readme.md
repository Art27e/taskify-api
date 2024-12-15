## A Simple RESTful To-Do API in Go

Project demonstrates key features as:
1. **Basic CRUD Operations**: The API allows for the creation, reading, updating, and deletion of tasks.
2. **Database Integration**: SQLite with GORM for object-relational mapping
3. **Clear API Endpoints**: Well-defined endpoints for viewing tasks (all, active, completed), creating tasks, updating task status, and deleting tasks.
4. **Error Handling**: The API checks for errors, such as missing task IDs or invalid requests, and responds with appropriate status codes.

Links: <br>
**All tasks** <br>
> http://localhost:8080/tasks
**Completed tasks** <br>
> http://localhost:8080/tasks/completed
**Active tasks** <br>
> http://localhost:8080/tasks/active
**Task by ID number** <br>
> http://localhost:8080/tasks/<ID Number>

Instructions: <br>
**Add a new task** <br>
use body.json file <br>
Powershell: <br>
> Invoke-RestMethod -Uri "http://localhost:8080/tasks" -Method Post -ContentType "application/json" -InFile "body.json"
**Mark as completed/uncompleted** <br>
Powershell: <br>
> Invoke-WebRequest -Uri "http://localhost:8080/update?id=<ID Number>" -Method PATCH
**Delete from the list** <br>
Powershell: <br>
> Invoke-WebRequest -Uri "http://localhost:8080/tasks/del/<ID Number>" -Method DELETE      


