package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//CarregarTeplates insere os templetes html na variavel templates   
func CarregarTempletes() {
	templates = template.Must(template.ParseGlob("views/*.html"))  
}

//ExecutarTemplate renderiza uma pagina html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}