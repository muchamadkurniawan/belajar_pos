package service

import (
	"belajar_pos/helper"
	"belajar_pos/model/domain"
	"belajar_pos/model/web"
	"belajar_pos/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type KasirServiceImpl struct {
	KasirRepository repository.KasirRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewKasirService(kasirRepository repository.KasirRepository, DB *sql.DB, validate *validator.Validate) KasirService {
	return &KasirServiceImpl{
		KasirRepository: kasirRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *KasirServiceImpl) Create(ctx context.Context, request web.KasirCreateRequest) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(nil)
	}
	defer tx.Commit()
	kasir := domain.Kasir{
		Nama:   request.Nama,
		Alamat: request.Alamat,
	}
	service.KasirRepository.Save(ctx, tx, kasir)
	//if save.Nip != -1 {
	//	defer tx.Commit()
	//}
	//defer tx.Rollback()
}

func (service *KasirServiceImpl) FindAll(ctx context.Context) []web.KasirResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(nil)
		tx.Rollback()
	}
	defer tx.Commit()

	kasirs := service.KasirRepository.GetAll(ctx, tx)
	kasirResponses := helper.ToKasirResponses(kasirs)
	return kasirResponses
}

func (service *KasirServiceImpl) FindById(ctx context.Context, nip int) web.KasirResponse {
	//TODO implement me
	panic("implement me")
}

func (service *KasirServiceImpl) Update(ctx context.Context, response web.KasirResponse) {
	//TODO implement me
	panic("implement me")
}

func (service *KasirServiceImpl) Delete(ctx context.Context, nip int) {
	//TODO implement me
	panic("implement me")
}
