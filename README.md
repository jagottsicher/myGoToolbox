# myGoToolbox
myGoToolbox is package which contains a collection of useful functions and handy tools for recurring tasks in my Go projects

## thousands.go
Thousands s mainly consists of the function Separate which adds thousands separators to numbers. It takes two arguments: the number and a language-code you want to use. It returns a string containing the number with thousands separators which are used in that region/country. Actually Separate supports English and German. Feels free to open an issue and request more notations.

```
func Separate(N interface{}, lang string) (string, error)
```

Examples:
```
	n := 1023.123456789
	output_en, _ := Separate(n, "en")
	output_de, _ := Separate(n, "de")
	fmt.Println(output_en)
	fmt.Println(output_de)
	// Output:
	// 1,023.123456789
	// 1.023,123456789

```