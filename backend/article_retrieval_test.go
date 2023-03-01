package main

import (
	"testing"
)

func TestEmptyArticleID(t *testing.T) {
	id := ""
	contents := getArticleContentsByID(id)
	if contents != nil {
		t.Fatalf(`getArticleContentsByID("") = %q, want match for %v`, contents, nil)
	}
}
