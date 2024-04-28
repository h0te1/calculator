package main
import (
	"fmt"
	"os"
	"strings"
	"errors"
	"strconv"
	"bufio"
)

func main() {
	var input string
	fmt.Println("Введите выражение арабскими или римскими цифрами")
	fmt.Scanf(input)
	base(strings.ReplaceAll(input, " ", ""))
}

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"V":    5,
	"IV":   4,
	"I":    1,
}
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}
var a, b *int
var operations = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

func checkRange(num *int) {
	return *num > 0 && *num <= 10
}

func base(s string) {
	IsNotMathOperationError := errors.New("Не является математической операцией (слишком много мат. операций)")
	MathEquasionFormError := errors.New("Не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	RangeError := errors.New("Калькулятор умеет работать только с числами от 1 до 10")
	NotationError := errors.New("В примере может использоваться только одна система счисления")
	var operator string
	var StringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operations {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(IsNotMathOperationError)
	case len(operator) < 1:
		panic(MathEquasionFormError)
	}


	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			StringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}
	switch StringsFound {
	case 0:
		if val, ok := operations[operator]; ok == true && checkRange(&numbers[0]) == true && checkRange(&numbers[1]) == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(RangeError)
		}
	case 1:
		panic(NotationError)
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok == true && checkRange(&val) == true {
                romansToInt = append(romansToInt, val)
            } else {
                panic(RangeError)
            }
			if val, ok := operations[operator]; ok == true {
				a, b = &romansToInt[0], &romansToInt[1]
				intToRoman(val())
			}
		}
	}
}

func intToRoman(romanResult int) {
	ZeroRomansError := errors.New("В римских цифрах не может быть нуля")
	NegativeRomansError := errors.New("Римские цифры не могут быть отрицательными")
	var FinalRomanString string
	switch {
	case romanResult == 0:
		panic(ZeroRomansError)
	case romanResult < 0:
		panic(NegativeRomansError)
	case romanResult > 0:
		for romanResult > 0 {
			for _, elem := range convIntToRoman {
				for i := elem; i <= romanResult; {
					for index, value := range roman {
						if value == elem {
							FinalRomanString += index
							romanResult -= elem
						}
					}
				}
			}
		}
		fmt.Println(FinalRomanString)
	}
}