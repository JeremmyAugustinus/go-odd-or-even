package go_odd_or_even

import "fmt"

func OddOrEven(i int) string {
	if i%2 == 0 {
		return fmt.Sprintf("[OddOrEven] %d adalah bilangan genap", i)
	} else {
		return fmt.Sprintf("[OddOrEven] %d adalah bilangan ganjil", i)
	}
}
