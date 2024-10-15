package authors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"tz2/pkg"
)

// получение списка всех авторов
func AuthorHandlerGet(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM author"
	authors, err := pkg.SelectAllAuthors(query)
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
	var NewAuthor []pkg.Author
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

	err = pkg.PostAuthors(NewAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// получение автора по id
func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, err := pkg.SelectOneAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(fmt.Sprintf("%v\n%v", author.Name, author.BirthYear)))
}

// обновление автора по id
func PutAuthorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var author pkg.Author
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(buf.Bytes(), &author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pkg.UpdateAuthorById(id, author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// удаление автора по id
func DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := pkg.DeleteByAuthorId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
