package service

import "github.com/ainmtsn1999/sql-book-api/model"

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (count int64, err error)
	DeleteBook(id int64) (count int64, err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	// call repo
	// res, err = s.repo.CreateBook(in)
	// if err != nil {
	// 	return res, err
	// }

	// return res, nil

	return s.repo.CreateBook(in)
}

func (s *Service) GetAllBook() (res []model.Book, err error) {
	return s.repo.GetAllBook()
}

func (s *Service) GetBookById(id int64) (res model.Book, err error) {
	return s.repo.GetBookById(id)
}
func (s *Service) UpdateBook(in model.Book) (count int64, err error) {
	return s.repo.UpdateBook(in)
}
func (s *Service) DeleteBook(id int64) (count int64, err error) {
	return s.repo.DeleteBook(id)
}
