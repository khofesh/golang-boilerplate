package developer_test

import (
	"go-test/developer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []developer.Developer{
		{Name: "Elliot"},
		{Name: "Elliot"},
		{Name: "David"},
		{Name: "Alexander"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	t.Log(developer.FilterUnique(input))

	assert.Equal(t, expected, developer.FilterUnique(input))
}

func TestNegativeFilterUnique(t *testing.T) {
	input := []developer.Developer{
		{Name: "Elliot"},
		{Name: "Elliot"},
		{Name: "David"},
		{Name: "Alexander"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.NotEqual(t, expected, developer.FilterUnique(input))
}
