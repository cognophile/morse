package dictionary

func Get() map[string]string {
	dict := make(map[string]string)

	dict["0"] = "-----"
	dict["1"] = ".----"
	dict["2"] = "..---"
	dict["3"] = "...--"
	dict["4"] = "....-"
	dict["5"] = "....."
	dict["6"] = "-...."
	dict["7"] = "--..."
	dict["8"] = "---.."
	dict["9"] = "----."
	dict["A"] = ".-"
	dict["B"] = "-..."
	dict["C"] = "-.-."
	dict["D"] = "-.."
	dict["E"] = "."
	dict["F"] = "..-."
	dict["G"] = "--."
	dict["H"] = "...."
	dict["I"] = ".."
	dict["J"] = ".---"
	dict["K"] = "-.-"
	dict["L"] = ".-.."
	dict["M"] = "--"
	dict["N"] = "-."
	dict["O"] = "---"
	dict["P"] = ".--."
	dict["Q"] = "--.-"
	dict["R"] = ".-."
	dict["S"] = "..."
	dict["T"] = "-"
	dict["U"] = "..-"
	dict["V"] = "...-"
	dict["W"] = ".--"
	dict["X"] = "-..-"
	dict["Y"] = "-.--"
	dict["Z"] = "--.."

	return dict
}

func Invert(original map[string]string) map[string]string {
	inverted := make(map[string]string)
	for k, v := range original {
		inverted[v] = k
	}

	return inverted
}
