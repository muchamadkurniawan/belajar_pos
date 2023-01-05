package helper

import (
	"belajar_pos/model/domain"
	"belajar_pos/model/web"
)

func ToKasirResponses(kasirs []domain.Kasir) []web.KasirResponse {
	var kasirResponses []web.KasirResponse
	for _, kasir := range kasirs {
		kasirResponses = append(kasirResponses, ToKasirResponse(kasir))
	}
	return kasirResponses
}

func ToKasirResponse(kasir domain.Kasir) web.KasirResponse {
	return web.KasirResponse{
		Nip:    kasir.Nip,
		Nama:   kasir.Nama,
		Alamat: kasir.Alamat,
	}
}
