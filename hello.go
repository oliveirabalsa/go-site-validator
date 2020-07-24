package main

import "fmt"
import "os"
import "net/http"

func main() {
	showIntroduction()
	showMenu()

	command := readCommand()

	switch command {
	case 1:
		initmonitoring()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Comando não encontrado")
		os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Leonardo"
	version := 1.1

	fmt.Println("Hello ", name)
	fmt.Println("Esta é a versão", version)
}
func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var commandReaded int
	fmt.Scan(&commandReaded)

	fmt.Println("You choose the option:", commandReaded)

	return commandReaded
}

func initmonitoring() {
	fmt.Println("Monitorando...")
	site := "https://google.com.br"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("O Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O Site", site, "Falhou em seu carregamento!")
	}
	// fmt.Println(response)
}
