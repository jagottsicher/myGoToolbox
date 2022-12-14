// myGoToolbox is package which contains a collection of useful functions and handy tools for recurring tasks in your Go projects
package myGoToolbox

import (
	"fmt"
	"strconv"
	"strings"
)

// Separate takes a numerical value and a language code and returns a string with thousands separators as commonly used in this lang/region. Actually supports English, German and French (intl. recommended) notation. Invalid or empty language codes let Separate return English notation as default/fallback. Scientific notation for very small or very big numbers is taken care of.
func Separate(N interface{}, lang ...string) (string, error) {

	// handles all numeric types, also unsigneds (and byte and rune as they're aliases)
	switch N.(type) {
	case int:
	case uint:
	case int16:
	case int32:
	case int64:
	case int8:
	case uint16:
	case uint32:
	case uint64:
	case uint8:
	case float32:
	case float64:
	default:
		return "", fmt.Errorf("Not a valid number!")
	}

	// convert to a string
	n := fmt.Sprintf("%v", N)

	// dealing with scientific notation
	if strings.Contains(n, "e") {
		nAsFlaot64, _ := strconv.ParseFloat(n, 64)
		n = strconv.FormatFloat(nAsFlaot64, 'f', -1, 64)
	}

	// if lang omitted English is default
	if len(lang) < 1 {
		lang[0] = "en"
	}

	switch lang[0] {
	case "de":

		n = strings.ReplaceAll(n, ",", ".")

		dec := ""

		if strings.Index(n, ".") != -1 {
			dec = n[strings.Index(n, ".")+1 : len(n)]
			n = n[0:strings.Index(n, ".")]
		}

		for i := 0; i <= len(n); i = i + 4 {
			a := n[0 : len(n)-i]
			b := n[len(n)-i : len(n)]
			n = a + "." + b
		}

		if n[0:1] == "." {
			n = n[1:len(n)]
		}

		if n[len(n)-1:len(n)] == "." {
			n = n[0 : len(n)-1]
		}
		n = strings.Trim(n, " ")

		if dec != "" {
			n = n + "," + dec
		}

		return n, nil

	case "fr":
		// TO-DO space as separator some more sophisticated
		n = strings.ReplaceAll(n, ",", ".")

		dec := ""

		if strings.Index(n, ".") != -1 {
			dec = n[strings.Index(n, ".")+1 : len(n)]
			n = n[0:strings.Index(n, ".")]
		}

		for i := 0; i <= len(n); i = i + 4 {
			a := n[0 : len(n)-i]
			b := n[len(n)-i : len(n)]
			n = a + "." + b
		}

		if n[0:1] == "." {
			n = n[1:len(n)]
		}

		if n[len(n)-1:len(n)] == "." {
			n = n[0 : len(n)-1]
		}

		if dec != "" {
			n = n + "," + dec
		}
		// like German but exchange all "." to whitespaces. some lazy, I know :)
		n = strings.ReplaceAll(n, ".", " ")

		return n, nil

	case "en":

		n = strings.ReplaceAll(n, ",", "")

		dec := ""

		if strings.Index(n, ".") != -1 {
			dec = n[strings.Index(n, ".")+1 : len(n)]
			n = n[0:strings.Index(n, ".")]

		}

		for i := 0; i <= len(n); i = i + 4 {
			a := n[0 : len(n)-i]
			b := n[len(n)-i : len(n)]
			n = a + "," + b
		}

		if n[0:1] == "," {
			n = n[1:len(n)]
		}

		if n[len(n)-1:len(n)] == "," {
			n = n[0 : len(n)-1]
		}

		if dec != "" {
			n = n + "." + dec
		}

		return n, nil
	}
	return n, nil
}
