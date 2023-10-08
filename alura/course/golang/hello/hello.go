package main

import (
	"fmt"
)

func main() {
	nome := "Marcelo"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	// Necessário import reflect
	// fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	// fmt.Scan(&comando)

	fmt.Println("O endereço da minha variável é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço esse comando")
	}
}
