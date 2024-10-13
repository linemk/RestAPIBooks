package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

// запрос на поиск книг
func BooksHandlerGet(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM book"
	books, err := SelectAllBooks(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var booksName []string
	for _, book := range books {
		booksName = append(booksName, book.Title)
	}

	booksNameInJson, err := json.Marshal(booksName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "json; charset=utf-8")
	_, _ = w.Write(booksNameInJson)
}

// запрос на добавление книги
func BooksHandlerPost(w http.ResponseWriter, r *http.Request) {
	var NewBook []Book
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(buf.Bytes(), &NewBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authors, err := SelectAllAuthors("SELECT * FROM author")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// проверка на id автора
	var MaxAuthorId int
	for _, author := range authors {
		if author.Id > MaxAuthorId {
			MaxAuthorId = author.Id
		}
	}

	for _, book := range NewBook {
		if book.AuthorId > MaxAuthorId {
			http.Error(w, fmt.Sprintf("Wrong author"), http.StatusInternalServerError)
			return
		}
	}

	err = PostBooks(NewBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// получаем по ID
func GetBooksId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := SelectOneBooks(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(book))
}
