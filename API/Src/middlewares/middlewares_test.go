// middlewares_test.go

package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogger(t *testing.T) {
	// Crie uma função de teste simulada para usar como próxima função
	var proximaFuncaoChamada bool
	proximaFuncaoSimulada := func(w http.ResponseWriter, r *http.Request) {
		proximaFuncaoChamada = true
	}

	// Crie uma solicitação de teste simulada
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)

	// Crie um gravador de resposta de teste simulado
	w := httptest.NewRecorder()

	// Chame a função Logger com a função de teste simulada
	handler := Logger(http.HandlerFunc(proximaFuncaoSimulada))
	handler.ServeHTTP(w, req)

	// Verifique se a próxima função foi chamada
	if !proximaFuncaoChamada {
		t.Error("A próxima função não foi chamada.")
	}
}

func TestAutenticar(t *testing.T) {
	// Crie uma função de teste simulada para usar como próxima função
	var proximaFuncaoChamada bool
	proximaFuncaoSimulada := func(w http.ResponseWriter, r *http.Request) {
		proximaFuncaoChamada = true
	}

	// Crie uma solicitação de teste simulada
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)

	// Adicione um token de autenticação simulado aos cabeçalhos da solicitação
	req.Header.Add("Authorization", "Bearer meu-token-simulado")

	// Crie um gravador de resposta de teste simulado
	w := httptest.NewRecorder()

	// Chame a função Autenticar com a função de teste simulada
	handler := Autenticar(http.HandlerFunc(proximaFuncaoSimulada))
	handler.ServeHTTP(w, req)

	// Verifique se a próxima função foi chamada
	if !proximaFuncaoChamada {
		t.Error("A próxima função não foi chamada.")
	}
}
