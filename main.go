package main

import (
	"net/http"
	"os"

	"github.com/Dev-ManavSethi/HomeAutomtionServer/controllers"
)

type Student struct {
	Name  string
	ID    int32
	Class string
}

func init() {

}

func main() {

	http.HandleFunc("/", controllers.Home)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
