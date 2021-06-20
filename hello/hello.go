package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()

	/*if comando == 1 {
		fmt.Println("Monitoramento Iniciado...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs:")
	} else if comando == 0 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Não foi possível completar a operação")
	}*/

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitoramento Iniciado...")
	case 2:
		fmt.Println("Exibindo Logs:")
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Não foi possível completar a operação")
	}

}

func exibeIntroducao() {
	fmt.Println("Seja Bem-Vindo!")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
