package autenticacao

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestToken(t *testing.T) {
	var usuarioId uint64
	usuarioId = 10

	Token, err := CriarToken(uint64(usuarioId))
	if err != nil {
		t.Fatalf("Erro ao criar o token '%v'.", err)
	}
	

	t.Run("Extraindo Token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/exemplo", nil)
		req.Header.Set("Authorization", "Bearer "+Token)

		esperado := Token
		resultado := extrairToken(req)

		if esperado != resultado {
			t.Errorf("Erro: esperado '%s', resultado '%s'.", esperado, resultado)
		}
	})

	t.Run("Token Inválido", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/exemplo", nil)
		err := ValidarToken(req)
		if err == nil {
			t.Error("Esperava um erro para token inválido, mas nenhum erro ocorreu.")
		}
	})

	t.Run("Token Ausente", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/exemplo", nil)
		token := extrairToken(req)
		if token != "" {
			t.Error("Esperava uma string vazia para token ausente, mas recebeu algo diferente.")
		}
	})

	t.Run("Extrair usuarioId", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/exemplo", nil)
		req.Header.Set("Authorization", "Bearer "+Token)

		esperado := usuarioId
		resultado, err := ExtrairUsuarioID(req) 
		if err != nil {
			t.Fatalf("Erro: '%v'.", err)
		}

		if esperado != resultado {
			t.Errorf("Erro: esperado '%d', resultado '%d'.", esperado, resultado)
		}
	})

	t.Run("Verifica se o Token esta valido", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/exemplo", nil)
		req.Header.Set("Authorization", "Bearer "+Token)

		err := ValidarToken(req)
		if err != nil {
			t.Fatalf("Erro: '%v'.", err)
		}
	})


}
