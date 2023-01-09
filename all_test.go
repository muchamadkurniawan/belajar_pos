package belajar_pos

import (
	"belajar_pos/app"
	"belajar_pos/helper"
	"belajar_pos/model/domain"
	"belajar_pos/model/web"
	"belajar_pos/repository"
	"belajar_pos/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"path"
	"testing"
)

func TestGetAllRepositoryConsole(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	repo := repository.NewKasirRepository()
	var kasirs []domain.Kasir
	kasirs = repo.GetAll(context.Background(), tx)
	fmt.Println(kasirs)
}

func TestFindAllServiceKasirConsole(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	kasirResponses := kasirService.FindAll(context.Background())
	fmt.Println(kasirResponses)
}

func TestFindAllServieKasirConsoleToMap(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	kasirResponses := kasirService.FindAll(context.Background())
	var datas []map[string]interface{}
	for index, _ := range kasirResponses {
		kasire := kasirResponses[index]
		var myMap map[string]interface{}
		data, _ := json.Marshal(kasire)
		json.Unmarshal(data, &myMap)
		datas = append(datas, myMap)
		fmt.Println(myMap)
		fmt.Println(myMap["nip"])
		fmt.Println(myMap["nama"])
		fmt.Println(myMap["alamat"])
		fmt.Println("--------------")
	}
	fmt.Println(datas)
	for index, _ := range datas {
		fmt.Println(index)
	}
}

func TestFindAllServieKasirConsoleToWeb(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func TestRepositoryInsert(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	repo := repository.NewKasirRepository()
	kasir1 := domain.Kasir{
		Nama:   "namatesting1",
		Alamat: "alamat1",
	}
	save := repo.Save(context.Background(), tx, kasir1)
	if save.Nip != -1 {
		tx.Commit()
	}
	tx.Rollback()
}

func TestServiceSaveConsole(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	kasir := web.KasirCreateRequest{
		Nama:   "testingService2",
		Alamat: "alamtservice2",
	}
	kasirService.Create(context.Background(), kasir)
}

func TestServiceSaveWeb(t *testing.T) {
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
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

	})

	fmt.Println("server started at localhost:9000")
	http.HandleFunc("/", show)
	http.ListenAndServe(":9000", nil)
}

func show(w http.ResponseWriter, r *http.Request) {
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

func TestServiceFindByIdConsole(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	kasirRepo := repository.NewKasirRepository()
	kasirService := service.NewKasirService(kasirRepo, db, validate)
	response := kasirService.FindById(context.Background(), 1)
	fmt.Println(response)
}
