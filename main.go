package main

import (
	"belajar_pos/app"
	"belajar_pos/controller"
	"belajar_pos/repository"
	"belajar_pos/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	kasirController := controller.NewKasirController(kasirService)
	http.HandleFunc("/", kasirController.FindAll)
	http.HandleFunc("/insert", kasirController.Create)
	http.HandleFunc("/deleteKasir", kasirController.Delete)
	http.HandleFunc("/showKasir", kasirController.Update)
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
