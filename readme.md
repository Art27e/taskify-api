## A Simple RESTful To-Do API in Go

Project demonstrates key features as:
1. **Basic CRUD Operations**: The API allows for the creation, reading, updating, and deletion of tasks.
2. **Database Integration**: You use SQLite with GORM for object-relational mapping, which is a good choice for a simple to-do list app.
3. **Clear API Endpoints**: You have well-defined endpoints for viewing tasks (all, active, completed), creating tasks, updating task status, and deleting tasks.
4. **Error Handling**: The API checks for errors, such as missing task IDs or invalid requests, and responds with appropriate status codes.