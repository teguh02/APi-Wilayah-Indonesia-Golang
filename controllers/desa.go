package controllers

import (
	"encoding/json"
	"net/http"
	. "wilayahIndonesia/app/config"
	. "wilayahIndonesia/app/models"
)

func Desa_By_KecamatanId(w http.ResponseWriter, r *http.Request) {
	var desa []Villages
	var results = make(map[string]interface{})
	id_kecamatan := r.URL.Query().Get("id_kecamatan")

	// Cek apakah id_kecamatan ada atau tidak
	if id_kecamatan == "" {
		results["status"] = "id_kecamatan_tidak_boleh_kosong"
		results["data"] = []Villages{}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	} else {
		query := Db().Where("district_id = ?", id_kecamatan).Find(&desa)

		// cek apakah data ditemukan atau tidak
		if query.RowsAffected == 0 {
			results["status"] = "data_tidak_ditemukan"
			results["data"] = desa
		} else {
			// Menyusun hasil respon
			results["status"] = "success"
			results["data"] = desa
		}

		// Menampilkan hasil respon menjadi JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

	}
}

func Desa_By_id(w http.ResponseWriter, r *http.Request) {
	var desa Villages
	var results = make(map[string]interface{})
	id_desa := r.URL.Query().Get("id_desa")

	// Cek apakah id_desa ada atau tidak
	if id_desa == "" {
		results["status"] = "id_desa_tidak_boleh_kosong"
		results["data"] = Villages{}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	} else {
		query := Db().Where("id = ?", id_desa).First(&desa)

		// cek apakah data ditemukan atau tidak
		if query.RowsAffected == 0 {
			results["status"] = "data_tidak_ditemukan"
			results["data"] = desa
		} else {
			// Menyusun hasil respon
			results["status"] = "success"
			results["data"] = desa

		}

		// Menampilkan hasil respon menjadi JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

	}

}