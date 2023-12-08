package main

import (
	"net/http"
	"webAPI_lesson/repository"
	"webAPI_lesson/routes"
	"webAPI_lesson/utils"
)

func main() {

	repository.RedisDB()
	utils.LoadTempleate("template/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
