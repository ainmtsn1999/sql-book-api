package app

import (
	"fmt"
	"os"

	"github.com/ainmtsn1999/sql-book-api/config"
	"github.com/ainmtsn1999/sql-book-api/handler"
	"github.com/ainmtsn1999/sql-book-api/repository"
	"github.com/ainmtsn1999/sql-book-api/route"
	"github.com/ainmtsn1999/sql-book-api/service"
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
