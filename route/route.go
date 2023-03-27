package route

import (
	"github.com/ainmtsn1999/sql-book-api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/books") // prefix
	{
		api.POST("", server.CreateBook)       // /employees
		api.GET("", server.GetAllBook)        // /employees
		api.GET("/:id", server.GetBookById)   // /employees/:id
		api.PUT("/:id", server.UpdateBook)    // /employees/:id
		api.DELETE("/:id", server.DeleteBook) // /employees/:id
	}
}
