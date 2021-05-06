package main

import (
	"fmt"

	"github.com/fgdemussy/luca/ledger"
)

func main() {
	fmt.Println("START")
	a := ledger.Account{Name: "Hola"}
	fmt.Println(a.Name)
}
