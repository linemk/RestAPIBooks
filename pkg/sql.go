package pkg

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "modernc.org/sqlite"
)

// возвращаем всех авторов
func SelectAllAuthors(query string) ([]Author, error) {
	db, err := sql.Open("sqlite", "sqldata/books.db")
	if err != nil {
		return []Author{}, fmt.Errorf("Cannot open db: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return []Author{}, fmt.Errorf("Cannot execute query: %v", err)
	}
	defer rows.Close()
	var authors []Author

	for rows.Next() {
		var author Author
		err = rows.Scan(&author.Id, &author.Name, &author.BirthYear)
		if err != nil {
			return []Author{}, fmt.Errorf("Cannot scan row: %v", err)
		}
		authors = append(authors, author)
	}
	err = rows.Err()
	if err != nil {
		return []Author{}, fmt.Errorf("Cannot scan rows: %v", err)
	}
	return authors, nil
}

// публикуем автора
func PostAuthors(authors []Author) error {
	db, err := sql.Open("sqlite", "sqldata/books.db")
	if err != nil {
		return fmt.Errorf("Cannot open db: %v", err)
	}
	defer db.Close()

	for _, author := range authors {
		_, err := db.Exec("INSERT INTO author (fio, birth_date) VALUES (:fio, :birth_date)",
			sql.Named("fio", author.Name), sql.Named("birth_date", author.BirthYear))
		if err != nil {
			return fmt.Errorf("Cannot execute query: %v", err)
		}
	}
	return nil
}

// выбор всех книш
func SelectAllBooks(query string) ([]Book, error) {
	db, err := sql.Open("sqlite", "sqldata/books.db")
	if err != nil {
		return []Book{}, fmt.Errorf("Cannot open db: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return []Book{}, fmt.Errorf("Cannot execute query: %v", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.PublishedDate)
		if err != nil {
			return []Book{}, fmt.Errorf("Cannot scan row: %v", err)
		}
		books = append(books, book)
	}
	err = rows.Err()
	if err != nil {
		return []Book{}, fmt.Errorf("Cannot scan rows: %v", err)
	}
	return books, nil
}

// вставка книг
func PostBooks(newBook []Book) error {
	db, err := sql.Open("sqlite", "sqldata/books.db")
	if err != nil {
		return fmt.Errorf("Cannot open db: %v", err)
	}
	defer db.Close()

	for _, book := range newBook {
		_, err := db.Exec("INSERT INTO book (title, author_id, published_date) VALUES (:title, :author_id, :published_date)",
			sql.Named("title", book.Title), sql.Named("author_id", book.AuthorId), sql.Named("published_date", book.PublishedDate))
		if err != nil {
			return fmt.Errorf("Cannot execute query or wrong author: %v", err)
		}
	}
	return nil
}

// выбор всех книш
func SelectOneBooks(id string) (string, error) {
	db, err := sql.Open("sqlite", "sqldata/books.db")
	if err != nil {
		return "", fmt.Errorf("Cannot open db: %v", err)
	}
	defer db.Close()
	idForRow, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("Cannot convert to int: %v", err)
	}
	row := db.QueryRow("SELECT title FROM book WHERE id= :id", sql.Named("id", idForRow))

	var book Book
	err = row.Scan(&book.Title)
	if err != nil {
		return "", fmt.Errorf("Cannot scan row: %v", err)
	}

	return book.Title, nil
}
