package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `json:"deleted_at"`
}

func (b Book) Validation() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required),
		validation.Field(&b.Author, validation.Required),
		validation.Field(&b.Desc, validation.Required))
}
