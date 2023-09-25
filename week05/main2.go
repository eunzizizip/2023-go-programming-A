package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Input score :")
	reader := bufio.NewReader(os.Stdin)
	inputNumber := reader.ReadString(\n)
	fmt.Println(inputNumber)
}