package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
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

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, err := SelectOneAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authors, err := SelectAllAuthors("SELECT * FROM author")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var maxAuthorsId int
	for _, v := range authors {
		if v.Id > maxAuthorsId {
			maxAuthorsId = v.Id
		}
	}
	if author.Id > maxAuthorsId {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(fmt.Sprintf("%v\n%v", author.Name, author.BirthYear)))
}

func PutAuthorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var author Author
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
	err = UpdateAuthorById(id, author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := DeleteByAuthorId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
