package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthCheckDTO struct {
	Message string
	Time    time.Time
}

func (h HealthCheckDTO) ToString() string {
	b, err := json.MarshalIndent(h, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (hh *HttpHandler) CheckServerHealth(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckDTO{
		Message: "OK",
		Time:    time.Now(),
	}

	b, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}
