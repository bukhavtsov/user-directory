# User directory

### Technologies

Frontend: Vue.js, Bootstrap

Backend: Golang, gorm, gorilla/mux, gorm-paginator, PostgreSQL

DevOps: Makefile, Docker, Docker-compose

### Description

Full Stack Application for storing users.

|                Path                   |    Method     |                   Description               |
| --------------------------------------|---------------| --------------------------------------------|
| `/users/{id}`                         |    `DELETE`   |   delete user                               |
| `/users`                              |    `POST`     |   create new user                           |
| `/users/uploadicon/{id}`              |    `PUT`      |   upload `user_icon` for user with `id`     |
| `/users/{id}`                         |    `PUT`      |   update user                               |
| `/users/pagination/{page}/{limit}`    |    `GET`      |   get all users with pagination             |
| `/users/{id}`                         |    `GET`      |   get user by `id`                          |
| `/users/find/{first_name}/{last_name}`|    `GET`      |   get user by `first_name` and `last_name`  |

### Usage

1. Run server on port `8080`

```bash
go run ./cmd/main.go
```

2. Open URL  `http://localhost:8080`

### Makefile commands

Allow all targets to be self documenting

```bash
make help
```