package encoder

import (
	"strings"

	"cognophile.com/morse/dictionary"
)

func encode(message string) string {
	encoded := ""
	dict := dictionary.Get()
	message = strings.ToUpper(message)

	for _, el := range strings.Split(message, " ") {
		for _, v := range el {
			encoded += dict[string(v)] + "/"
		}

		encoded += " "
	}

	return encoded
}

func Encode(message string) string {
	encoded := encode(message)
	return strings.TrimSuffix(encoded, " ")
}
