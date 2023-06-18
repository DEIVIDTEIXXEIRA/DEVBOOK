package modelos

import "time"

//Publicacao representa uma publicação feita pelo usuario 
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}