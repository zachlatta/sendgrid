package main

import (
	"net/http"

	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))

	m.Post(`/new_email`, HandleNewEmail)
}

func HandleNewEmail(r *http.Request) int {
	return http.StatusOK
}
