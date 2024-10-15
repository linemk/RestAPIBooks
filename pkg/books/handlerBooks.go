package books

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"tz2/pkg"
	"tz2/pkg/authors"
)

// запрос на поиск книг
func BooksHandlerGet(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM book"
	books, err := pkg.SelectAllBooks(query)
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
	var NewBook []pkg.Book
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
	// проверка на id
	for _, book := range NewBook {
		err = authors.SelectAuthorMaxId(strconv.Itoa(book.AuthorId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = pkg.PostBooks(NewBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// получаем по ID
func GetBooksId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := pkg.SelectOneBooks(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(book))
}

// обновляем по id
func PutBooksId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book pkg.Book
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(buf.Bytes(), &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = authors.SelectAuthorMaxId(strconv.Itoa(book.AuthorId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pkg.UpdateBookById(id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteBooksId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := pkg.DeleteByBookId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
