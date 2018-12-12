// Package add implements numbers to words converter.
package add

import (
	"fmt"
	"math"
)

type text struct {
	word        string
	groupedText []string
}

type digit struct {
	countGroupNumber int
	groupedNumber    []int
	number           int
	minus            bool
	positiveNumber   float64
}
type number struct {
	txt *text
	num *digit
}

// Convert converts number into the words.
func Convert(i int) string {
	n := create(i)
	return n.convert().String()
}

func create(i int) *number {
	return &number{
		txt: &text{
			groupedText: []string{},
		},
		num: &digit{
			number: i,
			minus:  i < ZERO,
			countGroupNumber: func(i int) int {
				return int((math.Log10(math.Abs(float64(i))) / 3) + 1)
			}(i),
			positiveNumber: func(i int) float64 {
				return math.Abs(float64(i))
			}(i),
		},
	}
}

func (n *number) convert() (*number) {
	// Zero rule
	if n.num.number == ZERO {
		n.txt.word = smallNumbers[ZERO]
	}

	// create three-digit groups 12345 => {12, 345}
	n.createDigitGroup()

	// Convert small groups number to word  {12, 345} => {Twelve, Three Hundred Forty Five}
	n.smallDigitGroupToText()

	// Add mega prefix and create string {Twelve, Three Hundred Forty Five} => "Twelve Thousand Three Hundred Forty Five"
	n.addMegaPrifix()

	return n
}

func (n *number) addMegaPrifix() {
	for i := range n.txt.groupedText {
		var space string
		if i > 0 {space = " "}
		if n.num.groupedNumber[i] != ZERO {
			prefix := n.txt.groupedText[i] + space + bigNumbers[i]
			if len(n.txt.word) != ZERO {
				prefix += and
			}
			n.txt.word = prefix + n.txt.word
		}
	}
}

func (n *number) createDigitGroup() {
	for i := ZERO; i < n.num.countGroupNumber; i++ {
		n.num.groupedNumber = append(n.num.groupedNumber, int(math.Mod(n.num.positiveNumber, THOUSAND)))
		n.num.positiveNumber /= THOUSAND
	}
}

func (n *number) smallDigitGroupToText() {

	for i := range n.num.groupedNumber {
		res := ""
		hundredsIndex := n.num.groupedNumber[i] / HUNDRED
		hundredMod := intMod(n.num.groupedNumber[i], HUNDRED)

		if hundredsIndex != ZERO {
			res = hundredNumber[hundredsIndex]

			if hundredMod != ZERO {
				res += and
			}
		}

		tenIndex := hundredMod / 10
		tenMod := intMod(hundredMod, 10)

		if tenIndex >= 2 {
			res += tensNumbers[tenIndex]

			if tenMod != ZERO {
				res += and + smallNumbers[tenMod]
			}
		} else if hundredMod != ZERO {
			res += smallNumbers[hundredMod]
		}

		n.txt.groupedText = append(n.txt.groupedText, res)
	}

}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func (n *number) String() string {
	if n.num.number < ZERO {
		n.txt.word = minus + n.txt.word
	}
	return fmt.Sprintf("%s", n.txt.word)
}
