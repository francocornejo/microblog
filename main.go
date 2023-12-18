package main

import (
	"fmt"
	"log"
	"microblog/api"
	"microblog/handler"
	"microblog/repository"
	"microblog/service"
	"net/http"

	_ "microblog/docs"

	"github.com/gin-gonic/gin"
)

// @title Ejercicio UALA
// @version 1.0
// @description API realizada para enviar o recibir mensajes y seguir a otros usuarios.

// @contact.name Cornejo Franco
// @contact.email cornejo.francodavid@gmail.com

// @host localhost:8080
// @BasePath /microblog
func main() {
	router := gin.New()

	port := ":8080" // o el puerto que desees
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	api.SetupRouter(router, ConfigHandler())

	// Inicia el servidor
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(fmt.Sprintf("Error en el servidor: %s\n", err))
	}

}

func ConfigHandler() handler.MicroBlogHandler {
	sqlRepository := repository.NewSQLRepository()
	service := service.NewService(sqlRepository)
	return handler.NewMicroblogHandler(service)
}
