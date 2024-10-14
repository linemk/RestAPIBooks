package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"tz2/pkg"
)

func main() {
	r := chi.NewRouter()
	r.Get("/authors", pkg.AuthorHandlerGet)
	r.Post("/authors", pkg.AuthorsHandlerPost)

	r.Get("/books", pkg.BooksHandlerGet)
	r.Post("/books", pkg.BooksHandlerPost)

	r.Get("/books/{id}", pkg.GetBooksId)
	r.Put("/books/{id}", pkg.PutBooksId)
	r.Delete("/books/{id}", pkg.DeleteBooksId)

	r.Get("/authors/{id}", pkg.GetAuthorById)
	r.Put("/authors/{id}", pkg.PutAuthorById)
	r.Delete("/authors/{id}", pkg.DeleteAuthorById)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
