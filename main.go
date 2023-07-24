package main

import (
	"fmt"

	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
