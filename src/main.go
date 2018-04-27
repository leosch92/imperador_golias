package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	got, err := ParseFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	var t Tree

	if len(os.Args) > 2 {
		switch os.Args[2] {

		case "-result":
			fmt.Println("-------- RESULTADO DO PARSER --------")
			fmt.Println(got)
		case "-bplc":
			t = *got.(*Tree)
			fmt.Println("-------- ÁRVORE BPLC GERADA  --------")
			printTree(&t)
		case "-verbose":
			t = *got.(*Tree)
			fmt.Println("\n\n------- RESOLUÇÃO SMC --------")
			resolverSMC(iniciaSMC(), t, true)
		default:
			fmt.Println("Comando não reconhecido.")

		}
	} else {
		t := *got.(*Tree)
		fmt.Println("-------- EXECUÇÃO DO PROGRAMA -------")
		resolverSMC(iniciaSMC(), t, false)
	}

	os.Exit(0)

}
