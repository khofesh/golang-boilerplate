package yamltohtml_test

import (
	"fmt"
	"go-test/yamltohtml"
	"os"
	"testing"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestMain(m *testing.M) {
	fmt.Println("hello")

	ret := m.Run()
	fmt.Println("tests have executed")
	os.Exit(ret)
}

func TestYamltoHTML(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "Test Case 1",
			path:     "testdata/test_01.yml",
			expected: "<html><head><title>Test Case 1 page</title></head><body>some content</body></html>",
		},
		{
			desc:     "Test Case 2",
			path:     "testdata/test_02.yml",
			expected: "<html><head><title>Test Case 2 page</title></head><body>some content</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := yamltohtml.YamltoHTML(test.path)
			if err != nil {
				t.Fail()
			}

			t.Log(result)

			if result != test.expected {
				t.Fail()
			}
		})
	}
}
