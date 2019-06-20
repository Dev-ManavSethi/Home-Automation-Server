package main

import (
	"net/http"
	"os"

	"github.com/Dev-ManavSethi/Home-Automation-Server/utils"

	"github.com/joho/godotenv"

	"github.com/Dev-ManavSethi/Home-Automation-Server/controllers"
)

type Student struct {
	Name  string
	ID    int32
	Class string
}

func init() {

	err := godotenv.Load(".env")
	utils.LogErrorOrSuccess(err, "Error setting env variables from /.env", "Sucessfully set env variables from /.env")

}

func main() {

	http.HandleFunc("/", controllers.Home)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	utils.LogErrorOrSuccess(err, "Error starting HTTP server at port: "+os.Getenv("PORT"), "HTTP server listening at port: "+os.Getenv("PORT"))

}
