package main

import (
	"fmt"
	"math/cmplx"
	"strings"
)

var (
	module uint64     = 10000007
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Asin(-5 + 12i)
) // как объявлять глобальные переменные

const (
	tmpRune         rune = '▒'
	someInteresting int  = 228
) // Как объявлять глобальные константы

type longLONG uint64    // именованный тип данных (удобно когда функция должна принимать только этот тип данных)
type longerNum = uint64 // псевдоним

func ModifuSpaces(s, mode string) string {
	var newString string
	switch mode {
	case "dash":
		newString = strings.ReplaceAll(s, " ", "-")
		break
	case "underscore":
		newString = strings.ReplaceAll(s, " ", "_")
	default:
		newString = strings.ReplaceAll(s, " ", "*")
	}
	return newString
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", module, module)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("GIGACHAT " + ModifuSpaces("GIGACHAT", "dash") + "!")
}
