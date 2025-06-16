package main

import (
	"net/http"
)

// Define o cookie de consentimento com validade de 1 ano
func setCookieConsent(w http.ResponseWriter, consent string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "cookieConsent",
		Value:  consent,
		Path:   "/",
		MaxAge: 86400 * 365, // 1 ano em segundos
	})
}

// Handler para receber a decisão de consentimento
func handleCookieConsent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	consent := r.URL.Query().Get("value")
	if consent != "accepted" && consent != "declined" {
		http.Error(w, "Valor de consentimento inválido", http.StatusBadRequest)
		return
	}

	setCookieConsent(w, consent)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Preferência de cookies registrada"))
}

// Função para registrar a rota no seu main.go
func RegistrarRotaCookieConsent() {
	http.HandleFunc("/set-cookie-consent", handleCookieConsent)
}
