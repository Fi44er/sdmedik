package main

import (
	"fmt"
	"regexp"
)

func main() {
	tru := "123456789.12345678901234567890"
	reg := `^\d{9}\.\d{20}$`
	ok, _ := regexp.MatchString(reg, tru)
	fmt.Println(ok)
}
