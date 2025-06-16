package main

import (
	"log"
	"net/http"
)

func main() {
	RegistrarRotasCodigoEmail()
	RegistrarRotasAgendar()
	RegistrarRotasAPI()
	RegistrarRotaCookieConsent()

	http.HandleFunc("/api/sac", handleSacForm)

	// Arquivos estáticos e páginas HTML
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
