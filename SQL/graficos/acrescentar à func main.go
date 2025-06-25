func main() {
  http.HandleFunc("/primeiro-exame", primeiroExameHandler)
  http.HandleFunc("/raça-cor", racaCorHandler)
  http.HandleFunc("/resultados", resultadosHandler)
  http.HandleFunc("/idade", idadeHandler)
  http.HandleFunc("/escolaridade", escolaridadeHandler)
  http.HandleFunc("/outras-infomações", outrasInfosHandler)
}
