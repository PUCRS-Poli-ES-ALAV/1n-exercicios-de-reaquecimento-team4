package alphabet_sequency_9

func Run() {
	result := generate(3)
	for _, str := range result {
		println(str)
	}
}

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate(n int) []string {
	if n <= 0 {
		return []string{}
	}

	if n == 1 {
		return []string{string(alphabet[0])}
	}

	var result []string
	for i := 0; i < n; i++ {
		subCombinations := generate(n - 1)
		for _, sub := range subCombinations {
			result = append(result, string(alphabet[i])+sub)
		}
	}

	return result
}
