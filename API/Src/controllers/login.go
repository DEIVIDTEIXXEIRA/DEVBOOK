package controllers

import (
	"api/Src/autenticacao"
	"api/Src/banco"
	"api/Src/modelos"
	"api/Src/repositorios"
	"api/Src/respostas"
	"api/Src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Login é responsável por autenticar um usuario na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepsitorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.Id)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.Id, 10)
	
	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{ID: usuarioID, Token: token})

}
