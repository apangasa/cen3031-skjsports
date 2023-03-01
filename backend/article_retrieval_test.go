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

func TestNonExistentArticleID(t *testing.T) {
	id := "anIdThatDoesNotExist"
	contents := getArticleContentsByID(id)
	if contents != nil {
		t.Fatalf(`getArticleContentsByID("anIdThatDoesNotExist") = %q, want match for %v`, contents, nil)
	}
}
