package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ArmazenamentoMemoria struct {
	jogadores map[string]Jogador
}

func (a *ArmazenamentoMemoria) ObterPontuaçãoJogador(nome string) int {
	return a.jogadores[nome].gols
}

func (a *ArmazenamentoMemoria) RealizarGolJogador(nome string) bool {
	jogador := a.jogadores[nome]
	jogador.gols += 1
	return true
}

func (a *ArmazenamentoMemoria) CriarJogador(nome string) bool {
	a.jogadores[nome] = Jogador{1, nome, 0}
	return true
}

func (a *ArmazenamentoMemoria) Close() {}

func TestJogador(t *testing.T) {
	armazenamento := ArmazenamentoMemoria{
		jogadores: map[string]Jogador{
			"vitor":    {1, "vitor", 0},
			"leonardo": {2, "leo", 3},
		},
	}

	server := ServidorJogador{&armazenamento}

	t.Run("retornar 0 quando jogador não existente", func(t *testing.T) {
		nome := "toma da silva"

		req := NewGetPlayerRequest(nome)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		if res.Body.String() != "0" {
			t.Errorf("esperado 0, recebido %v", res.Body.String())
		}
	})

	t.Run("retornar 3 quando jogador leo", func(t *testing.T) {
		nome := "leo"

		req := NewGetPlayerRequest(nome)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		if res.Body.String() != "3" {
			t.Errorf("esperado 3, recebido %v", res.Body.String())
		}
	})
}

func NewGetPlayerRequest(nome string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}
