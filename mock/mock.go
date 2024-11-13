package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Go!"
const inicioContagem = 3

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	s.pausa(s.duracao)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

const escrita = "escrita"
const pausa = "pausa"

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
