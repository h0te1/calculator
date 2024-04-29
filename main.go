package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
)

var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

const (
	NotMathError = "Не является математической операцией. В примере должно быть два числа и один оператор"
	TwoOperandsError = "В примере не должно быть двух операторов"
	DifferentNumbersError = "В примере оба числа должны быть одной системы счисления" 
	NegativeRomanNumberError = "В римской системе нет чисел меньше нуля"
	ZeroRomanNumberError  = "В римской системе нет числа 0."
	MaxNumberValueError = "Калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно"
)

func errorCheck(num int) bool {
	if num < 1 || num > 10 {
        panic(MaxNumberValueError)
    } else {
		return true
	}
}

func base(s string) {
	NotMathError := errors.New("не является математической операцией. В примере должно быть два числа и один оператор")
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(TwoOperandsError)
	case len(operator) < 1:
		panic(NotMathError)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 0:
		if val, ok := operators[operator]; ok && errorCheck(numbers[0]) && errorCheck(numbers[1]) {
			a, b = &numbers[0], &numbers[1]
			fmt.Println("Ответ:",val())
		}
	case 1:
		panic(DifferentNumbersError)
	case 2:
		for _, elem := range romans {
			romanInArabic := romanToArabic(elem)
			if errorCheck(romanInArabic) {
				romansToInt = append(romansToInt, romanInArabic)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			answer := val()
			switch {
			case answer == 0:
				panic(ZeroRomanNumberError)
			case answer < 0:
				panic(NegativeRomanNumberError)
			}
			fmt.Println("Ответ:", arabicToRoman(val()))
		}
	}
}

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	arabic := 0
	prevValue := 0

	for _, char := range roman {
		value := romanNumerals[char]
		if value > prevValue {
			arabic += value - 2*prevValue
		} else {
			arabic += value
		}
		prevValue = value
	}

	return arabic
}

func arabicToRoman(arabic int) string {
	romanNumerals := map[int]string{100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	extraHelp := [9]int {100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := ""
	for arabic > 0 {
		for _, elem := range extraHelp {
			for i := elem; i <= arabic; {
				for equivalent, numerals := range romanNumerals {
					if equivalent == elem {
						roman += numerals
						arabic -= equivalent
					}
				}
			}
		}
	}
	return roman
}

func main() {
	fmt.Println("Введите математическое выражение:")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		base(strings.ToUpper(strings.TrimSpace(s)))
	}
}