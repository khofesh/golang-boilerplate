package yamltohtml_test

import (
	"go-test/yamltohtml"
	"testing"
)

type TestCase struct {
	desc     string
	path     string
	expected string
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
