package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputString string

	// scanln is scanning only the fist word in the string. Even though the document says it reads till \n
	// instead used and tested bufio Scanner.
	// fmt.Scanln(&inputString)
	// fmt.Print(inputString)

	//https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	inputString = scanner.Text()

	fmt.Println("Hello World !")
	fmt.Println(inputString)

}
