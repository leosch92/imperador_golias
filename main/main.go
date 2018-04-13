package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: calculator 'EXPR'")
	}
	got, err := ParseReader("", strings.NewReader(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	t := got.(Tree)

	fmt.Println("------- CÓDIGO BPLC GERADO ---------")
	printBplc(&t)

	fmt.Println("\n\n------- RESOLUÇÃO SMC --------")
	resolverSMC(iniciaSMC(), t)

}
