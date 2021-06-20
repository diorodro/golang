package main

import "fmt"

func main() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("Monitoramento...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs")
	} else if comando == 0 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Não foi possível completar a operação")
	}
}
