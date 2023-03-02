package main

import (
	"testing"
)

func TestNullSearch(t *testing.T) {
	search := ""
	contents := getArticleContentsBySearch(search)

	if contents != nil {
		t.Fatalf(`getArticleContentsBySearch("") = %q, want match for %v`, contents, nil)
	}
}

func TestNonExistentSearch(t *testing.T) {
	search := "jibberjabberjibberjabber"
	contents := getArticleContentsBySearch(search)

	if contents != nil {
		t.Fatalf(`getArticleContentsBySearch("jibberjabberjibberjabber") = %q, want match for %v`, contents, nil)
	}
}

func TestRegularSearchWithMultipleResults(t *testing.T) {
	search := "biden"
	contents := getArticleContentsBySearch(search)

	contentList := contents["results"].([]map[string]string)

	if contentList[0]["title"] != "Biden At UN To Call Russian War An Affront To Body's Charter one" || contentList[0]["id"] != "9" {
		t.Fatalf(`First article title and id do not match`)
	}

	if contentList[1]["title"] != "Biden Says U.S. Forces Would Defend Taiwan If China Invaded three" || contentList[1]["id"] != "21" {
		t.Fatalf(`Second article title and id do not match`)
	}

	if contentList[2]["title"] != "Biden Says Queen's Death Left 'Giant Hole' For Royal Family four" || contentList[2]["id"] != "30" {
		t.Fatalf(`Third article title and id do not match`)
	}
}

func TestRegularSearchWithSingleResult(t *testing.T) {
	search := "COVID"
	contents := getArticleContentsBySearch(search)

	contentList := contents["results"].([]map[string]string)

	if contentList[0]["title"] != "Over 4 Million Americans Roll Up Sleeves For Omicron-Targeted COVID Boosters one" || contentList[0]["id"] != "0" {
		t.Fatalf(`Article title and id do not match`)
	}
}

func TestCaseSensitivity(t *testing.T) {
	search := "covid"
	contents := getArticleContentsBySearch(search)

	contentList := contents["results"].([]map[string]string)

	if contentList[0]["title"] != "Over 4 Million Americans Roll Up Sleeves For Omicron-Targeted COVID Boosters one" || contentList[0]["id"] != "0" {
		t.Fatalf(`Article title and id do not match`)
	}
}

func TestBlankSearch(t *testing.T) {
	search := "    "
	contents := getArticleContentsBySearch(search)

	if contents != nil {
		t.Fatalf(`Wasn't expecting any articles to be brought up in the search.`)
	}
}

func TestMultipleWordSearch(t *testing.T) {
	search := "Million Americans"

	contents := getArticleContentsBySearch(search)

	contentList := contents["results"].([]map[string]string)

	if contentList[0]["title"] != "Over 4 Million Americans Roll Up Sleeves For Omicron-Targeted COVID Boosters one" || contentList[0]["id"] != "0" {
		t.Fatalf(`Article title and id do not match`)
	}
}
