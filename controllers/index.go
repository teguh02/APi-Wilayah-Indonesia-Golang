package controllers

import (
	"net/http"
	"html/template"
	"path"
	."wilayahIndonesia/app/config"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	var data = make(map[string]interface{})
	template_file, error := template.ParseFiles(path.Join("views", "index.html"))

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	

	data["title"] = App("name")
	data["description"] = App("description")
	data["version"] = App("version")
	data["author"] = App("author")
	data["github"] = App("github")

	template_file.Execute(w, data)
}