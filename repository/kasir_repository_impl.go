package repository

import (
	"belajar_pos/model/domain"
	"context"
	"database/sql"
	"errors"
)

type KasirRepositoryImpl struct{}

func NewKasirRepository() KasirRepository {
	return &KasirRepositoryImpl{}
}
func (KasirRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, kasir domain.Kasir) domain.Kasir {
	SQL := "insert into kasir(nama,alamat) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, kasir.Nama, kasir.Alamat)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	kasir.Nip = int(id)
	return kasir
}

func (KasirRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Kasir {
	SQL := "select nip,nama,alamat from kasir"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var kasirs []domain.Kasir

	for rows.Next() {
		kasir := domain.Kasir{}
		err := rows.Scan(&kasir.Nip, &kasir.Nama, &kasir.Alamat)
		if err != nil {
			panic(err)
		}
		kasirs = append(kasirs, kasir)
	}
	return kasirs
}

func (KasirRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, nip int) (domain.Kasir, error) {
	SQL := "select nip,nama,alamat from kasir where nip=?"
	rows, err := tx.QueryContext(ctx, SQL, nip)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	kasir := domain.Kasir{}
	if rows.Next() {
		err = rows.Scan(&kasir.Nip, &kasir.Nama, &kasir.Alamat)
		if err != nil {
			panic(err)
		}
		return kasir, nil
	} else {
		return kasir, errors.New("Data Tidak Ditemukan")
	}
}

func (KasirRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, kasir domain.Kasir) {
	SQL := "delete from kasir where nip = ?"
	_, err := tx.ExecContext(ctx, SQL, kasir.Nip)
	if err != nil {
		panic(err)
	}
}

func (KasirRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, kasir domain.Kasir) domain.Kasir {
	SQL := "update kasir set nama = ?, alamat = ? where nip = ?"
	_, err := tx.ExecContext(context.Background(), SQL, kasir.Nama, kasir.Alamat, kasir.Nip)
	if err != nil {
		return domain.Kasir{}
	}
	return kasir
}
