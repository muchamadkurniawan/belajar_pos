package controller

import (
	"belajar_pos/app"
	"belajar_pos/helper"
	"belajar_pos/model/web"
	"belajar_pos/repository"
	"belajar_pos/service"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

type KasirControllerImpl struct {
	KasirService service.KasirService
}

func NewKasirController(kasirService service.KasirService) KasirController {
	return &KasirControllerImpl{
		KasirService: kasirService,
	}
}

func (controller *KasirControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "input_kasir.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)

	var nama string
	fmt.Println(nama)
	fmt.Println("method : ", r.Method)
	fmt.Println("formValue :", r.FormValue("nama"))
	data := web.KasirCreateRequest{}
	if r.Method == "POST" {
		nama = r.Form.Get("nama")
		data.Nama = r.Form.Get("nama")
		data.Alamat = r.Form.Get("alamat")
		kasirService.Create(context.Background(), data)

		fmt.Println("nnnnnnn ", nama)
		fmt.Println(data)
		http.Redirect(w, r, "http://localhost:9000/", http.StatusMovedPermanently)
	}
	tmpl.Execute(w, "")
}

func (controller *KasirControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "update_kasir.html")
	tmpl, err := template.ParseFiles(filepath)
	nip := r.URL.Query().Get("nip")
	id, err := strconv.Atoi(nip)
	if err != nil {
		panic(err)
	}
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	response := kasirService.FindById(context.Background(), id)
	toMap := helper.StructToMap(response)
	fmt.Println(toMap)

	if err != nil {
		panic(err)
	}
	kasir := web.KasirResponse{
		Nip: id,
	}
	fmt.Println(r.Method)
	if r.Method == "POST" {
		fmt.Println(r.Method)
		r.FormValue("nama")
		kasir.Nama = r.Form.Get("nama")
		kasir.Alamat = r.Form.Get("alamat")
		fmt.Println(kasir)
		kasirService.Update(context.Background(), kasir)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	err = tmpl.Execute(w, toMap)
}

func (controller *KasirControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	nip := r.URL.Query().Get("nip")
	id, err := strconv.Atoi(nip)
	if err != nil {
		panic(err)
		fmt.Println("error konversi")
	}
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	kasirService.Delete(context.Background(), id)
	http.Redirect(w, r, "http://localhost:9000/", http.StatusMovedPermanently)
}

func (controller *KasirControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "show.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	kasirResponses := kasirService.FindAll(context.Background())
	datas := helper.StructSliceToMap(kasirResponses)
	err1 := tmpl.Execute(w, datas)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
