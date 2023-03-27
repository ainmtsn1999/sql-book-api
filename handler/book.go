package handler

import (
	"strconv"

	"github.com/ainmtsn1999/sql-book-api/helper"
	"github.com/ainmtsn1999/sql-book-api/model"
	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, gin.H{"New Book Data:": res})
}

func (h HttpServer) GetAllBook(c *gin.Context) {
	res, err := h.app.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.Ok(c, gin.H{"Book Datas:": res})
}

func (h HttpServer) GetBookById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetBookById(int64(idInt))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, gin.H{"Book Data:": res})
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in.Id = idInt

	res, err := h.app.UpdateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, gin.H{"Updated Data Amount:": res})
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.DeleteBook(int64(idInt))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, gin.H{"Deleted Data Amount:": res})
}
