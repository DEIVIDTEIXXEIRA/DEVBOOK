package rotas

import (
	"api/Src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)


// Rota ira representar todas as rotas da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasDePublicacoes...)


	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.Uri,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}
	return r
}
