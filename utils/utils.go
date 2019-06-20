package utils

import "log"

func LogErrorOrSuccess(err error, ErrorMessage, SuccessMessage string) {

	if err != nil {
		log.Println(ErrorMessage)
		log.Fatalln(err)
	} else if SuccessMessage != "" {
		log.Println(SuccessMessage)

	}
}
