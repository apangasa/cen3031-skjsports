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

func TestArticleWithJustText(t *testing.T) {
	id := "1"
	contents := getArticleContentsByID(id)

	expectedTitle := "American Airlines Flyer Charged, Banned For Life After Punching Flight Attendant On Video one"

	if contents["title"] != expectedTitle {
		t.Fatalf(`getArticleContentsByID("1")["title"] = %q, want match for %q`, contents["title"], expectedTitle)
	}

	contentList := contents["content"].([]map[string]string)

	if len(contentList) != 1 {
		t.Fatalf(`len(getArticleContentsByID("1")["content"]) = %q, want match for %q`, len(contentList), 1)
	}

	if contentList[0]["contentType"] != "text" {
		t.Fatalf(`getArticleContentsByID("1")["content"][0]["contentType"] = %q, want match for %q`, contentList[0]["contentType"], "text")
	}

	expectedText := "Lowry didn't play Saturday at Milwaukee because of what the Heat called left knee soreness. The Heat will re-evaluate him next week, when Miami has two more games before the All-Star break. He missed three games in January due to left knee discomfort and has struggled to find his shot since returning, averaging 5.6 points on 25% shooting and 4.4 assists in his last five games. Gabe Vincent got the start in Saturday's loss at Milwaukee, finishing with seven points."

	if contentList[0]["text"] != expectedText {
		t.Fatalf(`getArticleContentsByID("1")["content"][0]["text"] = %q, want match for %q`, contentList[0]["text"], expectedText)
	}

}
