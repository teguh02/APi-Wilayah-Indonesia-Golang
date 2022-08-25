package models

// model untuk tabel districts (kecamatan)
type Districts struct {
	Id string `json:"id"`
	Regency_id int32 `json:"regency_id"`
	Name string `json:"name"`
}