package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

func DataBaseGet() *sql.DB {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("Переменная окружения DB_PATH не установлена")
	}

	db, err := sql.Open("sqlite", "sqldata/books.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// возвращаем всех авторов
func SelectAllAuthors(query string) ([]Author, error) {
	db := DataBaseGet()
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
	db := DataBaseGet()
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

// выбор всех книг
func SelectAllBooks(query string) ([]Book, error) {
	db := DataBaseGet()
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
	db := DataBaseGet()
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

// выбор одной книги
func SelectOneBooks(id string) (string, error) {
	db := DataBaseGet()
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

func SelectOneAuthor(id string) (Author, error) {
	db := DataBaseGet()
	defer db.Close()

	var author Author

	idForRow, err := strconv.Atoi(id)
	if err != nil {
		return Author{}, fmt.Errorf("Cannot convert to int: %v", err)
	}
	row := db.QueryRow("SELECT * FROM author WHERE id= :id", sql.Named("id", idForRow))
	err = row.Scan(&author.Id, &author.Name, &author.BirthYear)
	if err != nil {
		return Author{}, fmt.Errorf("Cannot scan row: %v", err)
	}

	return author, nil
}

func UpdateBookById(id string, book Book) error {
	db := DataBaseGet()
	defer db.Close()

	_, err := db.Exec("UPDATE book SET title = :title, author_id = :author_id, published_date = :published_date WHERE id= :id",
		sql.Named("title", book.Title),
		sql.Named("author_id", book.AuthorId),
		sql.Named("published_date", book.PublishedDate),
		sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("Cannot execute query: %v", err)
	}
	return nil
}

func DeleteByBookId(id string) error {
	db := DataBaseGet()
	defer db.Close()

	_, err := db.Exec("DELETE FROM book WHERE id= :id", sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("Cannot execute query: %v", err)
	}
	return nil
}

func UpdateAuthorById(id string, author Author) error {
	db := DataBaseGet()
	defer db.Close()

	_, err := db.Exec("UPDATE author SET fio = :fio, birth_date = :birth_date WHERE id= :id",
		sql.Named("fio", author.Name),
		sql.Named("birth_date", author.BirthYear),
		sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("Cannot execute query: %v", err)
	}
	return nil
}

func DeleteByAuthorId(id string) error {
	db := DataBaseGet()
	defer db.Close()

	_, err := db.Exec("DELETE FROM author WHERE id= :id", sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("Cannot execute query: %v", err)
	}
	return nil
}
