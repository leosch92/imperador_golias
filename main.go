package main

import (
	"fmt"
	"log"
)

func main() {
	/*if len(os.Args) != 2 {
		log.Fatal("Usage: calculator 'EXPR'")
	}*/
	got, err := ParseFile("program.imp") //strings.NewReader(os.Args[1]))
	if err != nil {
		log.Fatal(err)
		fmt.Println("oi")
	}

	fmt.Println(got)

	/*t := got.(Tree)

	fmt.Println("------- CÓDIGO BPLC GERADO ---------")
	printBplc(&t)

	fmt.Println("\n\n------- RESOLUÇÃO SMC --------")
	resolverSMC(iniciaSMC(), t)*/

}
