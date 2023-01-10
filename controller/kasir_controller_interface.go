package controller

import (
	"belajar_pos/service"
	"net/http"
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
	//TODO implement me
	panic("implement me")
}

func (controller *KasirControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *KasirControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *KasirControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
