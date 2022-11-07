package go_odd_or_even

import "fmt"

func OddOrEven(i int, msg string) string {
	if i%2 == 0 {
		return fmt.Sprintf("[%s] %d adalah bilangan genap", msg, i)
	} else {
		return fmt.Sprintf("[%s] %d adalah bilangan ganjil", msg, i)
	}
}
