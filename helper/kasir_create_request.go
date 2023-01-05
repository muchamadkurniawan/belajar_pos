package helper

type KasirCreateRequest struct {
	Nama   string `validate:"requiered, min = 1, max =100" json:"nama"`
	Alamat string `validate:"requiered, min = 1, max =100" json:"alamat"`
}
