package modelos

import "time"

//usuario representa uma pessoa utilizando a aplicação
type Usuario struct {
	ID         uint64       `json:"id"`
	Nome       string       `json:"nome"`
	Email      string       `json:"email"`
	Nick       string       `json:"nick"`
	CriadoEm   time.Time    `json:"criadoEm"`
	Seguidores []Usuario    `json:"seguidores"`
	Seguindo   []Usuario    `json:"seguindo"`
	Publicacao []Publicacao `json:"publicacoes"`
}
