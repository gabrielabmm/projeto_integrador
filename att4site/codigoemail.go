package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var lastReenvio = make(map[string]time.Time)

type PageData struct {
	Email     string
	Codigo    string
	Erro      string
	Sucesso   bool
	Countdown int
	CanResend bool
	Redirect  bool
}

func renderCodigoPage(w http.ResponseWriter, r *http.Request) {
	email := "teste@exemplo.com" // Em produção: extrair da sessão ou querystring
	now := time.Now()

	last := lastReenvio[email]
	restante := 60 - int(now.Sub(last).Seconds())
	if restante < 0 {
		restante = 0
	}

	data := PageData{
		Email:     email,
		CanResend: restante == 0,
		Countdown: restante,
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		codigo := r.FormValue("codigo")
		data.Codigo = codigo

		if action == "verificar" {
			match, _ := regexp.MatchString(`^\d{6}$`, codigo)
			if !match {
				data.Erro = "O código deve ter 6 dígitos numéricos."
			} else if codigo == "123456" {
				http.Redirect(w, r, "/escolher-nova-senha.html", http.StatusSeeOther)
				return
			} else {
				data.Erro = "Código incorreto. Tente novamente."
			}
		} else if action == "reenviar" {
			if restante == 0 {
				lastReenvio[email] = now
				data.Countdown = 60
				data.CanResend = false
				data.Sucesso = true
			} else {
				data.Erro = "Aguarde " + strconv.Itoa(restante) + "s para reenviar."
			}
		}
	}

	tmpl, err := template.ParseFiles("codigoemail.html")
	if err != nil {
		log.Println("Erro ao carregar template:", err)
		http.Error(w, "Erro interno ao renderizar página", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func RegistrarRotasCodigoEmail() {
	http.HandleFunc("/codigoemail", renderCodigoPage)
}
