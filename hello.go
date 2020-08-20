package main

import "fmt"
import "os"
import "net/http" 

func main() {
	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			initmonitoring()
		case 2:
			fmt.Println("Showing Logs...")
		case 0:
			fmt.Println("Exiting the program...")
			os.Exit(-1)
		default:
			fmt.Println("Command not found")
		}
	}
}

func showIntroduction() {
	name := "Leonardo"
	version := 1.1

	fmt.Println("Hello ", name)
	fmt.Println("Esta é a versão", version)
}
func showMenu() {
	fmt.Println("1- Init monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var commandReaded int
	fmt.Scan(&commandReaded)

	fmt.Println("You choose the option:", commandReaded)

	return commandReaded
}

func initmonitoring() {
	fmt.Println("Monitoring...")
	site := "https://google.com.br"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("The website", site, "was successfully loaded!")
	} else {
		fmt.Println("The website", site, "loading failed!")
	}
}
