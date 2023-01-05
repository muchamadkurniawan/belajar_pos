package repository

import (
	"belajar_pos/model/domain"
	"context"
	"database/sql"
)

type KasirRepository interface {
	Save(ctx context.Context, tx *sql.Tx, kasir domain.Kasir) domain.Kasir
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Kasir
	GetById(ctx context.Context, tx *sql.Tx, nip int) (domain.Kasir, error)
	Delete(ctx context.Context, tx *sql.Tx, kasir domain.Kasir)
	Update(ctx context.Context, tx *sql.Tx, kasir domain.Kasir) domain.Kasir
}
