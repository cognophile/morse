package main

import (
	"strings"

	"cognophile.com/morse/decoder"
	"cognophile.com/morse/encoder"
	io "cognophile.com/morse/io"
)

func showGreeting() {
	io.Info("\n---------- Morse ----------\n")
}

func showInformation(separate bool) {
	if separate {
		io.Info("\n---------- Menu ----------\n")
	}
	io.Info("Use '/' to delimit characters, and spaces to delimit words")
	io.Info("Select execution mode...")
	io.Info("  1. Encode")
	io.Info("  2. Decode\n")
	io.Info("Type 'q' to quit\n")
}

func readUserInput() string {
	input, err := io.Read()
	if input == "" && err != nil {
		io.Error("An error occured while reading the input", err)
		return ""
	}

	return strings.ToLower(strings.TrimSuffix(input, "\n"))
}

func print(str string) {
	io.Info(str + "\n")
}

func main() {
	mode := ""
	showGreeting()
	showInformation(false)

	for mode != "q" {
		mode = readUserInput()

		if mode == "1" || mode == "encode" {
			io.Info("\n----------------------------\n")
			io.Info("Type your message...\n")
			message := readUserInput()
			output := encoder.Encode(message)
			print(output)
			showInformation(true)
		}

		if mode == "2" || mode == "decode" {
			io.Info("\n----------------------------\n")
			io.Info("Type encoded message...\n")
			message := readUserInput()
			output := decoder.Decode(message)
			print(output)
			showInformation(true)
		}
	}
}
