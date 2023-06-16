package modelos

//DadosDeautenticacao contem o id e o token do usuário autenticado
type DadosDeautenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
