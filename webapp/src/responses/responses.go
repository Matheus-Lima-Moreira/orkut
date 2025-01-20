package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type Err struct {
	Err string `json:"err"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func HandleErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var err Err
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
