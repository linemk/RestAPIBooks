package main

import (
	"log"
	"net/http"
	"tz2/pkg/authors"
	"tz2/pkg/books"

	"github.com/go-chi/chi"
	_ "modernc.org/sqlite"
)

func main() {
	r := chi.NewRouter()
	// получение всего списка авторов
	r.Get("/authors", authors.AuthorHandlerGet)
	// публикация нового автора/авторов
	r.Post("/authors", authors.AuthorsHandlerPost)
	// получение автора по id
	r.Get("/authors/{id}", authors.GetAuthorById)
	// обновление автора по id
	r.Put("/authors/{id}", authors.PutAuthorById)
	// удаление автора по id
	r.Delete("/authors/{id}", authors.DeleteAuthorById)

	// получение всего списка книг
	r.Get("/books", books.BooksHandlerGet)
	// публикация всего списка книг
	r.Post("/books", books.BooksHandlerPost)
	// получение книги по id
	r.Get("/books/{id}", books.GetBooksId)
	// обновление книги по id
	r.Put("/books/{id}", books.PutBooksId)
	// удаление книги по id
	r.Delete("/books/{id}", books.DeleteBooksId)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
