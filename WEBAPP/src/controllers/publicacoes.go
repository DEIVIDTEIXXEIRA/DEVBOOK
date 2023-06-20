package controllers

import (
	"fmt"
	"net/http"
)

// CriarPublicacao chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criando Publicação....")
}
