package myGoToolbox

import (
	"fmt"
	"testing"
)

func ExampleEscapeStringRegex() {

	s := `How much $ for a 🦄?`

	fmt.Println(EscapeStringRegex(s))
	// Output:
	// How much \$ for a 🦄\?
}

// Table tests for en, de, and fr and malicious/missing language codes
func TestEscapeStringRegexTT(t *testing.T) {

	type testData struct {
		s      string
		result string
	}

	tests := []testData{
		testData{`How much $ for a 🦄?`, `How much \$ for a 🦄\?`},
		testData{`^(@|\s)*"anything":?`, `\^\(@\|\\s\)\*"anything":\?`},
		testData{`Escaping symbols like: .+*?()|[]{}^$`, `Escaping symbols like: \.\+\*\?\(\)\|\[\]\{\}\^\$`},
	}

	for _, v := range tests {
		wrongOutput := EscapeStringRegex(v.s)
		// fmt.Println(wrongOutput)
		if wrongOutput != v.result {
			t.Error("Expected", v.result, "got", wrongOutput)
		}
	}
}
