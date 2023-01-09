package web

type KasirCreateRequest struct {
	Nama   string `validate:"required, min=1, max=100:"json:"nama"`
	Alamat string `validate:"required, min=1, max=100:"json:"alamat"`
}
