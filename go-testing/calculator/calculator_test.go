package calculator_test

import (
	"go-test/calculator"
	"testing"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsAmstrong(t *testing.T) {

	t.Run("should return true", func(t *testing.T) {
		testCase := TestCase{
			value:    371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should return true for 370", func(t *testing.T) {
		testCase := TestCase{
			value:    370,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
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
