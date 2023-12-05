package main

import (
	"fmt"

	"github.com/hello/ejercicios"
)

func main() {

	ola, cola := ejercicios.ConviertoaNumero("500")

	fmt.Println(ola)
	fmt.Println(cola)
	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {

	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/

}
