package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// получение списка всех авторов
func AuthorHandlerGet(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM author"
	authors, err := SelectAllAuthors(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authorsInJson, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "json; charset=utf-8")
	_, _ = w.Write(authorsInJson)
}

// добавление автора
func AuthorsHandlerPost(w http.ResponseWriter, r *http.Request) {
	var NewAuthor []Author
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &NewAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = PostAuthors(NewAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
