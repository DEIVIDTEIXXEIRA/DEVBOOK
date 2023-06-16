package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ApiUrl representa a Url para comunicação com a api
var (
	APIURL  = ""
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

func Carregar() {
	var erro error
	
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}