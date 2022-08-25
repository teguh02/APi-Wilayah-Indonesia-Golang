package models

// model untuk tabel kabupaten 
type Regencies struct {
	Id string `json:"id"`
	Province_id int32 `json:"province_id"`
	Name string `json:"name"`
}