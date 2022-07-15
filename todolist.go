package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
) // Adicionando pacotes necessarios para aplicação

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Healt is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
} // a Função ira resonder {"alive": true} toda vez que ela for chamada

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todolist API server")
	route := mux.NewRouter()
	route.HandleFunc("/healthz", Healthz).Methods("GET") //criando Rota para chamar a função Healthz
	http.ListenAndServe(":8000", route)                  // chamando a rota e passando a porta que será ouvida
}

// stop in
// https://www.fadhil-blog.dev/blog/golang-todolist/#step-2-connect-to-mysql-db
