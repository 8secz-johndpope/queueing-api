package api

import (
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		return true
	}
	return false
}
