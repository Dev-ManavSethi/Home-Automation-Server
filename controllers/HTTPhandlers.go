package controllers

import (
	"fmt"
	"net/http"
)

func Home(HTTPresponseWriter http.ResponseWriter, HTTPrequest *http.Request) {

	fmt.Fprintln(HTTPresponseWriter, "Welcome to Home Automation Server!")

}
