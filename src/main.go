package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cognophile.com/morse/decoder"
	"cognophile.com/morse/encoder"
)

func main() {
	fmt.Print("Enter message: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	message := strings.TrimSuffix(input, "\n")

	encoded := encoder.Encode(message)
	fmt.Println("Encoded: " + strings.Replace(encoded, "/", "", -1))

	decoded := decoder.Decode(encoded)
	fmt.Println("Decoded: " + decoded)
}
