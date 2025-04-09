package main

import "fmt"

func main() {
	var firstVar string = "first variable"

	fmt.Println(firstVar)

	var (
		var1 string = "var"
		var2 string = "var2"
	)

	fmt.Println(var1, var2)

	var3, var4 := "var3", "var4"
	fmt.Println(var3, var4)

	var3, var4 = var4, var3

	fmt.Println(var3, var4)
}
