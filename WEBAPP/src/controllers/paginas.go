package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDelogin renderiza tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarPaginaDeCadastroDeUsuario carrega a pagina de cadastro de usuario 
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

//CarregarPaginaPrincipal carrega a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "home.html", nil)
}