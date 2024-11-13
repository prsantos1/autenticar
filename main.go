package main

import (
	"autenticar/server"
	"log"
)

func main() {
	log.Println("Servidor iniciado na porta 808")
	if err := server.Start(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
