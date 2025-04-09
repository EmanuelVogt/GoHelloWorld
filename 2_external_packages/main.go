package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main file")
	auxiliar.Write()

	error := checkmail.ValidateFormat("examplemailcom")
	fmt.Println(error)

}
