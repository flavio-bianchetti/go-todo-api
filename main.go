package main

import (
	"net/http"

	"github.com/flavio-bianchetti/go-todo-api/controllers"
)

func main() {
	mux : controllers.Register()
	http.ListenAndServe(":3000", mux)
}
