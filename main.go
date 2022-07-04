package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flavio-bianchetti/go-todo-api/controllers"
	models "github.com/flavio-bianchetti/go-todo-api/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controllers.Register()
	db := models.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
