package controllers

import (
	"fmt"
	"net/http"
)

func Home(HTTPresponseWriter http.ResponseWriter, HTTPrequest *http.Request) {

	fmt.Fprintln(HTTPresponseWriter, "Welcome to Home Automation Server!")

}

func App(HTTPresponseWriter http.ResponseWriter, HTTPrequest *http.Request) {

	Commands, ok := HTTPrequest.URL.Query()["command"]

	if !ok || len(Commands[0]) < 1 {

		return
	}

	switch Commands[0] {

	case "CreateClient":

		break

	case "EditClient":

		break

	}

}
