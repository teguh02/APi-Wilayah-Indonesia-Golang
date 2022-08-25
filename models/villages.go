package models

// model untuk tabel villages (desa)
type Villages struct {
	Id string `json:"id"`
	District_id int32 `json:"district_id"`
	Name string `json:"name"`
}