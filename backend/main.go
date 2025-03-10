package main

import "fmt"

type Test struct {
	Id  int
	Bit uint8
}

func main() {
	var db []Test
	for i := 0; i < 5; i++ {
		db = append(db, Test{Id: i, Bit: 1 << i})
	}
	var a uint8 = 1 << 0
	var b uint8 = 1 << 1

	var qwe uint8

	qwe |= a

	fmt.Println(qwe&b != 0)
}
