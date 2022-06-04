package calculator_test

import (
	"go-test/calculator"
	"testing"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsAmstrong(t *testing.T) {
	t.Run("test for all 3 digits armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			{name: "Testing value: 153", value: 153, expected: true},
			{name: "Testing value: 370", value: 370, expected: true},
			{name: "Testing value: 371", value: 371, expected: true},
			{name: "Testing value: 407", value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := calculator.CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})
}

func TestNegativeCalculateIsAmstrong(t *testing.T) {
	t.Run("should fail for case 350", func(t *testing.T) {
		testCase := TestCase{
			value:    350,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should fail for case 300", func(t *testing.T) {
		testCase := TestCase{
			value:    300,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
