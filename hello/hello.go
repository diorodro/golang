package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Seja Bem-Vindo!")

	for {
		exibeIntroducao()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs:")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não foi possível completar a operação")
			os.Exit(-1)
		}

	}
}

func exibeIntroducao() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento Iniciado...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso : Status", resp.Status)
	} else {
		fmt.Println("Site:", site, "está com problemas : Status", resp.Status)
	}
}
