package main


import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/SoheilRezaei/Go-API-/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go-API Service...")

	fmt.Println(`
	.d8888b.                d88888888888b.8888888 
	d88P  Y88b              d88888888   Y88b 888   
	888    888             d88P888888    888 888   
	888        .d88b.     d88P 888888   d88P 888   
	888  88888d88""88b   d88P  8888888888P"  888   
	888    888888  888  d88P   888888        888   
	Y88b  d88PY88..88P d8888888888888        888   
	 "Y8888P88 "Y88P" d88P     888888      8888888 
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}