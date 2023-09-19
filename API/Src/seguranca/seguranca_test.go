package seguranca

import (
	"testing"
)

func TestVerificaSenha(t *testing.T) {
	senha := "Golang@2023"
	senhaComHash, err := Hash(senha)

	if err != nil {
		t.Fatalf("Erro ao gerar o hash: %v", err)
	}

	t.Run("Teste utilizando senha correta", func(t *testing.T) {
		err = VerificarSenha(string(senhaComHash), senha)
		if err != nil {
			t.Fatalf("Erro ao verificar a senha correta: %v", err)
		}
	})

	t.Run("Teste utilizando senha incorreta", func(t *testing.T) {
		err = VerificarSenha(string(senhaComHash), "SenhaIncorreta")
		if err == nil {
			t.Fatal("Esperava um erro ao verificar uma senha incorreta, mas não ocorreu erro.")
		}
	})

}
