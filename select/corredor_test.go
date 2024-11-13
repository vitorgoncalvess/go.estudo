package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		servidorLento := criarServidorComAtraso(20 * time.Millisecond)
		servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida
		resultado, _ := Corredor(URLLenta, URLRapida)

		if resultado != esperado {
			t.Errorf("resultado %q, esperado %q", resultado, esperado)
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidor := criarServidorComAtraso(25 * time.Millisecond)

		defer servidor.Close()

		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}
