package main

import (
	"net/http"
	"web_start2/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
