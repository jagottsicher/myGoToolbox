# myGoToolbox
myGoToolbox is package which contains a collection of useful functions and handy tools for recurring tasks in my Go projects

## thousands.go
Thousands s mainly consists of the function Separate which adds thousands separators to numbers. It takes two arguments: the number and a language-code you want to use. It returns a string containing the number with thousands separators which are used in that region/country. Actually Separate supports English and German. Feels free to open an issue and request more notations.

```
func Separate(N interface{}, lang string) string
```

Examples:
```
n := 5000.61

fmt.Println("English notation: " + thousands.Separate(n, "en"))

// English notation: 5,000.61

fmt.Println("Deutsche Ausgabe: " + thousands.Separate(n, "de"))

// Deutsche Ausgabe: 5.000,61

```
