// ./app/queries/book_query.go
package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/snickers31/go-with-fiber/app/models"
)

type DBQueries struct {
	*sqlx.DB
}

func (q *DBQueries) GetBooks() ([]models.Book, error) {
	books := []models.Book{}

	query := `SELECT * FROM book`

	err := q.Get(&books, query)
	if err != nil {
		return books, err
	}

	return books, nil
}

func (q *DBQueries) GetBook(id uuid.UUID) (models.Book, error) {
	book := models.Book{}

	query := `SELECT * FROM book WHERE id = $1`

	err := q.Get(&book, query, id)
	if err != nil {
		return book, err
	}

	return book, nil

}

func (q *DBQueries) CreateBook(b *models.Book) (models.Book, error) {
	book := models.Book{}

	query := `INSERT INTO book VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserId, b.Title, b.Author, b.BookStatus, b.BookAttrs)

	if err != nil {
		return book, err
	}

	book.ID = b.ID
	book.CreatedAt = b.CreatedAt
	book.UpdatedAt = b.UpdatedAt
	book.UserId = b.UserId
	book.Title = b.Title
	book.Author = b.Author
	book.BookStatus = b.BookStatus
	book.BookAttrs = b.BookAttrs

	return book, nil
}

func (q *DBQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	query := `UPDATE book SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	_, err := q.Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)

	if err != nil {
		return err
	}

	return nil
}

func (q *DBQueries) DeleteBook(id uuid.UUID) error {
	query := `DELETE FROM book WHERE id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
