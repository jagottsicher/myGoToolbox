# myGoToolbox [![Actions Status](https://github.com/jagottsicher/myGoToolbox/workflows/Go/badge.svg)](https://github.com/jagottsicher/myGoToolbox/actions) [![GoDoc](https://godoc.org/github.com/jagottsicher/myGoToolbox?status.svg)](https://godoc.org/github.com/jagottsicher/myGoToolbox)
myGoToolbox is package which contains a collection of useful functions and handy tools for recurring tasks in your Go projects. myGoToolbox is aiming at as clear and clean Go code as possible and preferable to make use of built-in packages only and to import as reasonably maintained and simple packages only to reduce the impact of broken dependencies to a minimum.

## Install
```sh
go get github.com/jagottsicher/myGoToolbox
```
## Features
#### thousands.go
thousands.go mainly consists of the function `Separate()` which adds thousands separators to numbers. It takes two arguments: the number and a language-code you want to use. It returns a string with thousands separators as commonly used in this lang/region. Actually `Separate()` supports English, German and French (intl. recommended) notation. More information about [decimal separators](https://en.wikipedia.org/wiki/Decimal_separator). 

English is default/fallback for wrong or missing language code. Feels free to open an issue and request more or other notations.

```
func Separate(N interface{}, lang string) (string, error)
```

Examples:
```
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

```
#### terminal.go
`Clearscreen()` empties the terminal screen independently from the underlying operation system by calling the OS's builtin functions works smootly even on both Windows' cmd.exe and powershell.
