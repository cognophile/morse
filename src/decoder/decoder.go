package decoder

import (
	"strings"

	"cognophile.com/morse/dictionary"
)

func Decode(message string) string {
	decoded := ""
	dict := dictionary.Invert(dictionary.Get())

	for _, el := range strings.Split(message, " ") {
		for _, v := range strings.Split(el, "/") {
			decoded += dict[string(v)]
		}

		decoded += " "
	}

	return decoded
}
