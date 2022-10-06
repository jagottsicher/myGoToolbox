package myGoToolbox

import (
	"fmt"
	"testing"
)

func ExampleSeparate() {
	n := 1023.123456789
	output_en, _ := Separate(n, "en")
	output_de, _ := Separate(n, "de")
	fmt.Println(output_en)
	fmt.Println(output_de)
	// Output:
	// 1,023.123456789
	// 1.023,123456789
}

func TestSeparateTT(t *testing.T) {

	type testData struct {
		N      float64
		lang   string
		result string
	}

	tests := []testData{
		testData{1023.123456789, "en", "1,023.123456789"},
		//testData{1023123456789, "de", "1.023.123.456.789"},
		//testData{1023123456789, "en", "1,023,123,456,789"},
		//testData{10.000006316159832682479, "de", "10,000006316159832682479"},
	}

	for _, v := range tests {
		wrongOutput, _ := Separate(v.N, v.lang, v.result)
		fmt.Println(wrongOutput)
		if wrongOutput != v.result {
			t.Error("Expected", v.result, "got", wrongOutput)
		}
	}
}
