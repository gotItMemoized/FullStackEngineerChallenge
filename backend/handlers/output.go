package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// this is technically inefficient, but allows for fast iterations and we can still get very fast responses (this may become a problem with much larger data)
func WriteToOutput(w http.ResponseWriter, object interface{}) {
	output, err := json.Marshal(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(output)
	if err != nil {
		log.Printf("error while writing output: %+v\n", err)
	}
}
