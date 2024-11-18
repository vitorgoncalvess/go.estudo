package main

import (
	"database/sql"
	"fmt"
)

type ArmazenamentoMySQL struct {
	conn *sql.DB
}

func (a *ArmazenamentoMySQL) getJogador(nome string) Jogador {
	var jogador Jogador

	err := a.conn.QueryRow("SELECT * FROM jogador WHERE nome = ?", nome).Scan(&jogador.id, &jogador.nome, &jogador.gols)

	if err != nil {
		return Jogador{}
	}

	return jogador
}

func (a *ArmazenamentoMySQL) ObterPontuaçãoJogador(nome string) int {
	jogador := a.getJogador(nome)

	return jogador.gols
}

func (a *ArmazenamentoMySQL) CriarJogador(nome string) bool {
	if jogador := a.getJogador(nome); jogador.nome == nome {
		return false
	}

	insert, err := a.conn.Query("INSERT INTO jogador (nome, gols) VALUES (?, 0)", nome)

	if err != nil {
		return false
	}

	defer insert.Close()

	return true
}

func (a *ArmazenamentoMySQL) RealizarGolJogador(nome string) bool {
	jogador := a.getJogador(nome)
	if jogador.nome == "" {
		return false
	}

	insert, err := a.conn.Query("UPDATE jogador SET gol = ? WHERE nome = ?", jogador.gols+1, nome)

	if err != nil {
		return false
	}

	defer insert.Close()

	return true
}

func (a *ArmazenamentoMySQL) Close() {
	a.conn.Close()
}

func NovoArmazenamentoMySQL(user, password, database string) *ArmazenamentoMySQL {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, database))

	if err != nil {
		panic(err.Error())
	}

	return &ArmazenamentoMySQL{conn: db}
}
