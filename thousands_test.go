package myGoToolbox

import (
	"fmt"
	"testing"
)

func ExampleSeparate() {

	n := 1232323.123456789
	output_en, _ := Separate(n, "en")
	output_de, _ := Separate(n, "de")
	output_fr, _ := Separate(n, "fr")

	fmt.Println(output_en)
	fmt.Println(output_de)
	fmt.Println(output_fr)
	// Output:
	// 1,232,323.123456789
	// 1.232.323,123456789
	// 1 232 323,123456789
}

// Table tests for en, de, and fr and malicious/missing language codes
func TestSeparateTT(t *testing.T) {

	type testData struct {
		N      float64
		lang   string
		result string
	}

	tests := []testData{
		testData{1023.123456789, "en", "1,023.123456789"},
		testData{1023123456789, "de", "1.023.123.456.789"},
		testData{1023123456789, "en", "1,023,123,456,789"},
		testData{0.000006316159832682479, "de", "0,000006316159832682479"},
		testData{102312345678.9, "fr", "102 312 345 678,9"},
		testData{1023123456789, "fr", "1 023 123 456 789"},
		testData{102312345678.9, "bla", "102312345678.9"},
		testData{1023123456789, "bla", "1023123456789"},
		testData{102312345678.9, "", "102312345678.9"},
	}

	for _, v := range tests {
		wrongOutput, _ := Separate(v.N, v.lang, v.result)
		// fmt.Println(wrongOutput)
		if wrongOutput != v.result {
			t.Error("Expected", v.result, "got", wrongOutput)
		}
	}
}
