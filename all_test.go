package belajar_pos

import (
	"belajar_pos/app"
	"belajar_pos/helper"
	"belajar_pos/model/domain"
	"belajar_pos/model/repository"
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
		//var datas []map[string]interface{}
		//for index, _ := range kasirResponses {
		//	kasire := kasirResponses[index]
		//	var myMap map[string]interface{}
		//	data, _ := json.Marshal(kasire)
		//	json.Unmarshal(data, &myMap)
		//	datas = append(datas, myMap)
		//}
		datas := helper.StructSliceToMap(kasirResponses)
		err1 := tmpl.Execute(w, datas)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
