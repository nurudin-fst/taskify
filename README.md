# How to run
run 
```make docker-dev-up```

# Structur Endpoint

## Register
```
Methode: POST
{{host}}/register
Content-Type: application/json
Payload:
{
    "email": "...",
    "name": "...",
    "password": "..."
}
```

## Login
```
Methode: POST
{{host}}/login
Content-Type: application/json
Payload:
{
    "email": "...",
    "password": "..."
}
```

## Insert Project
```
Methode: POST
{{host}}/projects
Content-Type: application/json
Authorization: Bearer {{jwt_token}}
Payload:
{
    "name": "...",
    "description": "..."
}
```

## Insert Task
```
Methode: POST
{{host}}/projects/:id/tasks
Content-Type: application/json
Authorization: Bearer {{jwt_token}}
Payload:
{
    "title": "...",
    "description": "...",
    "deadline": "2025-07-17T01:01:03Z"
}
```

## Project List
```
Methode: GET
{{host}}/projects
Authorization: Bearer {{jwt_token}}
```

## Task List
```
Methode: GET
{{host}}/projects/:id/tasks
Authorization: Bearer {{jwt_token}}
```

## Task Update
```
Methode: PUT
{{host}}/task/:id
Content-Type: application/json
Authorization: Bearer {{jwt_token}}
Payload:
{
    "title": "...",
    "status": "...",
    "description": "...",
    "deadline": "2026-04-09T02:03:11Z"
}
```
## Project Update
```
Methode: PUT
{{host}}/project/:id
Authorization: Bearer {{jwt_token}}
Payload:
{
    "Name": "...",
    "Description": "..."
}
```
## Delete Task
```
Methode: DELETE
{{host}}/task/:id
Authorization: Bearer {{jwt_token}}
```
