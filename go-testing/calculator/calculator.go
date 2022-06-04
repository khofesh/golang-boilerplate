package calculator

import "math"

// CalculateIsArmstrong takes in a 3 digit number 'n'
// and returns true if it is an Armstrong number
// Armstrong number example 371 == 3^3 + 7^3 + 1^3
func CalculateIsArmstrong(n int) bool {
	a := n / 100
	b := n % 100 / 10
	c := n % 10
	return n == int(math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3))
}
