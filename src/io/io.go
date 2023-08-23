package io

import (
	"bufio"
	"fmt"
	"os"
)

func Info(message string) {
	fmt.Println(message)
}

func Error(message string, err error) {
	fmt.Printf("ERROR: %v - %v", message, err)
}

func Read() (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return input, nil
}
