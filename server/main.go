package main

import (
	"log"
	"net/http"

	"github.com/go.estudo/server/routes"

	_ "github.com/go-sql-driver/mysql"
)

type Jogador struct {
	id   int
	nome string
	gols int
}

type ArmazenamentoJogador interface {
	ObterPontuaçãoJogador(nome string) int
	RealizarGolJogador(nome string) bool
	CriarJogador(nome string) bool
	Close()
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	network := Network()

	network.Route("/jogadores/:id", routes.GetPlayerByName)
}

func (s *ServidorJogador) Close() {
	s.armazenamento.Close()
}

func main() {
	server := ServidorJogador{NovoArmazenamentoMySQL("root", "password", "gofc")}

	defer server.Close()

	if err := http.ListenAndServe(":8080", http.HandlerFunc(server.ServeHTTP)); err != nil {
		log.Fatalf("não foi possivel escutar na porta 8080 %v", err)
	}
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Route struct {
	route   string
	handler Handler
}

type Net struct {
	routes []Route
}

func (n *Net) Route(path string, handler Handler) {
	n.routes = append(n.routes, Route{path, handler})
}

func Network() Net {
	return Net{}
}
