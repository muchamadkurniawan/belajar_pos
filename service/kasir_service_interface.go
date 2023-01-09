package service

import (
	"belajar_pos/model/web"
	"context"
)

type KasirService interface {
	Create(ctx context.Context, request web.KasirCreateRequest)
	FindAll(ctx context.Context) []web.KasirResponse
	FindById(ctx context.Context, nip int) web.KasirResponse
	Update(ctx context.Context, response web.KasirResponse)
	Delete(ctx context.Context, nip int)
}
