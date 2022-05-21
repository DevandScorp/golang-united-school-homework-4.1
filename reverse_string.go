package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	for _, symbol := range(input) {
		output = string(symbol) + output
	}
	return output
}
