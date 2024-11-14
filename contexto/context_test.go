package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store foi cancelado")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWritter struct {
	written bool
}

func (s *SpyResponseWritter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWritter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("não implementado")
}

func (s *SpyResponseWritter) WriteHeader(statusCode int) {
	s.written = true
}

func TestContexto(t *testing.T) {
	data := "olá, mundo"
	t.Run("retorna dados da store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("resultado %q, esperado %q", response.Body.String(), data)
		}
	})

	t.Run("avisa a store para cancelar o trabalho se a requsição for cancelada", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWritter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("uma resposta não deveria ter sido escrita")
		}
	})
}
