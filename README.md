# go flat api sample
GoでREST API Server をフラットパッケージで実装しました。

## how to run

```
$ git pull https://github.com/Le0tk0k/go-flat-api-sample
$ cd go-flat-api-sample
$ docker-compose up -d --build
```

## db scheme

```sql
CREATE DATABASE IF NOT EXISTS sample;
USE sample;

CREATE TABLE IF NOT EXISTS users (
    id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name        VARCHAR(256) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS posts (
    id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    title       VARCHAR(256) NOT NULL,
    body        TEXT NOT NULL,
    author_id   INT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_author_id
    FOREIGN KEY (author_id)
    REFERENCES users (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

## api

**endpoint**

```go
a.router.HandleFunc("/users/{id}", a.getUser).Methods("GET")
a.router.HandleFunc("/users", a.getUsers).Methods("GET")
a.router.HandleFunc("/users", a.createUser).Methods("POST")
a.router.HandleFunc("/users/{id}", a.updateUser).Methods("PUT")
a.router.HandleFunc("/users/{id}", a.deleteUser).Methods("DELETE")

a.router.HandleFunc("/posts/{id}", a.getPost).Methods("GET")
a.router.HandleFunc("/posts", a.getPosts).Methods("GET")
a.router.HandleFunc("/posts", a.createPost).Methods("POST")
a.router.HandleFunc("/posts/{id}", a.updatePost).Methods("PUT")
a.router.HandleFunc("/posts/{id}", a.deletePost).Methods("DELETE")
```

## reference

- https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql  
- https://github.com/mingrammer/go-todo-rest-api-example
