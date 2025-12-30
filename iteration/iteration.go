package iteration

import "strings"

func Repeat(character string, chars int) string {
	var repeated strings.Builder

	for i := 0; i < chars; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func FullName(firstName string, lastName string) (fullName string) {
	names := []string{firstName, lastName}
	fullName = strings.Join(names, " ")
	return
}

func main() {

}
