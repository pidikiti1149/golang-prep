package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var s string

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"

	fmt.Println(s)

	s = ` 
	<html>
		<body>"Hy"</body>
	</html
	`
	fmt.Println(s)

	name := "阿蒂塔"

	fmt.Println(len(name)) //len calculates bytes of the string you need to use utf8 package to calculate actual chracters

	fmt.Println(utf8.RuneCountInString(name))

	msg := os.Args[1]
	l := len(msg)

	repeat := strings.Repeat("!", l) + msg + strings.Repeat("!", l)

	repeat = strings.ToUpper(repeat)

	fmt.Println(repeat)
}
