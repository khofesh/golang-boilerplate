package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const body = `{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}`

func BenchmarkUnmarshal(b *testing.B) {
	w := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(body))
		if err != nil {
			b.Fatal(err)
		}

		createMovieHandlerUnmarshal(w, r)
	}
}

func BenchmarkDecoder(b *testing.B) {
	w := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(body))
		if err != nil {
			b.Fatal(err)
		}

		createMovieHandlerDecoder(w, r)
	}
}
