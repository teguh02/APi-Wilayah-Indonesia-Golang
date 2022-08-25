package main

import (
	"fmt"
	"net/http"
	."wilayahIndonesia/app/controllers"
	."wilayahIndonesia/app/config"
)

func main() {
	defer http.ListenAndServe(App("port"), nil)
	fmt.Println("Server is running on port ", App("port"))

	// routing
	http.HandleFunc("/", Index)
	http.Handle("/assets/",
        http.StripPrefix("/assets/",
            http.FileServer(http.Dir("assets"))))
	
	http.HandleFunc("/provinsi_all", Provinsi_ListAll)
	http.HandleFunc("/provinsi_by_id", Provinsi_ById)
	http.HandleFunc("/kabupaten_by_province_id", Kabupaten_ByProvinceId)
	http.HandleFunc("/kabupaten_by_id", Kabupaten_ById)
	http.HandleFunc("/kecamatan_by_regency_id", Kecamatan_By_Regency_id)
	http.HandleFunc("/kecamatan_by_id", Kecamatan_By_id)	
	http.HandleFunc("/desa_by_district_id", Desa_By_KecamatanId)	
	http.HandleFunc("/desa_by_id", Desa_By_id)	
	
}