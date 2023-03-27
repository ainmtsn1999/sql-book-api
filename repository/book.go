package repository

import (
	"fmt"

	"github.com/ainmtsn1999/sql-book-api/model"

	"github.com/ainmtsn1999/sql-book-api/repository/query"
)

// interface employee
type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (count int64, err error)
	DeleteBook(id int64) (count int64, err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.AddBook,
		in.Title,
		in.Author,
		in.Desc,
	).Scan(
		&res.Id,
		&res.Title,
		&res.Author,
		&res.Desc,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	rows, err := r.db.Query(query.GetAllBook)

	if err != nil {
		return res, nil
	}

	defer rows.Close()

	for rows.Next() {
		var book = model.Book{}

		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			return res, nil
		}

		res = append(res, book)
	}

	return res, nil
}

func (r Repo) GetBookById(id int64) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.GetBookById,
		id,
	).Scan(
		&res.Id,
		&res.Title,
		&res.Author,
		&res.Desc,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}
func (r Repo) UpdateBook(in model.Book) (count int64, err error) {
	res, err := r.db.Exec(query.UpdateBook, in.Id, in.Title, in.Author, in.Desc)
	if err != nil {
		return 0, err
	}

	count, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	fmt.Println("Updated data amount:", count)

	return count, err

}

func (r Repo) DeleteBook(id int64) (count int64, err error) {
	res, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return 0, err
	}

	count, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	fmt.Println("Deleted data amount:", count)

	return count, err
}
