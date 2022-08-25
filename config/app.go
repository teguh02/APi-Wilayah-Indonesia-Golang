package config

import "os"

func App(key string) string {
	var app = map[string]string{}

	app["name"] = "Api Wilayah Seluruh Indonesia"
	app["version"] = "1.0.0"
	app["description"] = "API Wilayah Seluruh Indonesia dibuat menggunakan Bahasa Pemrograman golang"
	app["author"] = "Teguh Rijanandi"
	app["github"] = "https://github.com/teguh02/API-Wilayah-Indonesia-Golang"

	// Get Port from heroku or localhost
	app["port"] = ":" + os.Getenv("PORT")

	// return the configuration
	return app[key]
}
