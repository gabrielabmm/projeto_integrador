package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type SacForm struct {
	Email    string `json:"email"`
	Mensagem string `json:"mensagem"`
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(email)
}

func handleSacForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var form SacForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}

	if !isValidEmail(form.Email) {
		http.Error(w, "Email inválido", http.StatusBadRequest)
		return
	}

	if form.Mensagem == "" {
		http.Error(w, "Mensagem não pode estar vazia", http.StatusBadRequest)
		return
	}

	// Aqui você poderia salvar os dados, enviar email etc.
	log.Printf("SAC recebido: Email=%s | Mensagem=%s\n", form.Email, form.Mensagem)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mensagem recebida com sucesso"))
}
