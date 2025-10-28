package repositories

import (
	"database/sql"
	"serveractualizado/server/cmd/models"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteBookRepository struct {
	db *sql.DB
}

func NewSqliteBookRepository(databasePath string) (*SqliteBookRepository, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}
	return &SqliteBookRepository{db: db}, nil
}

func (r *SqliteBookRepository) Initialize() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL, 
	author TEXT NOT NULL,
	year INTEGER NOT NULL
	);`
	_, err := r.db.Exec(createTableQuery)
	return err
}

func (r *SqliteBookRepository) GetAll() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)

	}
	return books, nil
}

func (r *SqliteBookRepository) GetByID(id int) (*models.Book, error) {
	var book models.Book
	row := r.db.QueryRow("SELECT id, title, author, year FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró el libro con ese ID
			return &models.Book{}, nil
		}
		return &models.Book{}, err
	}
	return &book, nil
}

func (r *SqliteBookRepository) Create(book *models.Book) (models.Book, error) {
	result, err := r.db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", book.Title, book.Author, book.Year)
	if err != nil {
		return models.Book{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Book{}, err
	}
	book.ID = int(id)
	return *book, nil
}

func (r *SqliteBookRepository) Update(id int, book models.Book) (models.Book, error) {
	_, err := r.db.Exec(
		"UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?",
		book.Title, book.Author, book.Year, id,
	)
	if err != nil {
		return models.Book{}, err
	}
	book.ID = id
	return book, nil
}

func (r *SqliteBookRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
