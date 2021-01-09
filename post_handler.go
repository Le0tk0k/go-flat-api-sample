package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *app) getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	p := post{ID: id}
	if err := p.findByID(a.db); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Post not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJson(w, http.StatusOK, p)
}

func (a *app) getPosts(w http.ResponseWriter, _ *http.Request) {
	posts, err := findAllPosts(a.db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, posts)
}

func (a *app) createPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	authorID, err := strconv.Atoi(params["authorID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author ID")
	}

	var p post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.store(a.db, authorID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, map[string]string{"result": "Successfully created"})
}

func (a *app) updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var p post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p.ID = id
	if err := p.update(a.db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, p)
}

func (a *app) deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, "Invalid post ID")
	}

	p := post{ID: id}
	if err := p.delete(a.db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "Successfully deleted"})
}
