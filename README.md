# go flat api sample
GoでREST API Server をフラットパッケージで実装しました。

## how to run

```
$ git pull https://github.com/Le0tk0k/go-flat-api-sample
$ cd go-flat-api-sample
$ docker-compose up -d --build
```

## api

**endpoint**

```
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
