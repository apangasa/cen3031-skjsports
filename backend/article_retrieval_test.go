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

func TestArticleWithTextAndImage(t *testing.T) {
	id := "2"
	contents := getArticleContentsByID(id)

	expectedTitle := "23 Of The Funniest Tweets About Cats And Dogs This Week (Sept. 17-23) one"

	if contents["title"] != expectedTitle {
		t.Fatalf(`getArticleContentsByID("2")["title"] = %q, want match for %q`, contents["title"], expectedTitle)
	}

	contentList := contents["content"].([]map[string]string)

	if len(contentList) != 3 {
		t.Fatalf(`len(getArticleContentsByID("2")["content"]) = %q, want match for %q`, len(contentList), 3)
	}

	if contentList[0]["contentType"] != "text" {
		t.Fatalf(`getArticleContentsByID("2")["content"][0]["contentType"] = %q, want match for %q`, contentList[0]["contentType"], "text")
	}

	expectedText := "Lowry didn't play Saturday at Milwaukee because of what the Heat called left knee soreness. The Heat will re-evaluate him next week, when Miami has two more games before the All-Star break. He missed three games in January due to left"

	if contentList[0]["text"] != expectedText {
		t.Fatalf(`getArticleContentsByID("2")["content"][0]["text"] = %q, want match for %q`, contentList[0]["text"], expectedText)
	}

	if contentList[1]["contentType"] != "img" {
		t.Fatalf(`getArticleContentsByID("2")["content"][1]["contentType"] = %q, want match for %q`, contentList[1]["contentType"], "img")
	}

	expectedId := "3dh982"

	if contentList[1]["id"] != expectedId {
		t.Fatalf(`getArticleContentsByID("2")["content"][1]["img"] = %q, want match for %q`, contentList[1]["img"], expectedId)
	}

	if contentList[2]["contentType"] != "text" {
		t.Fatalf(`getArticleContentsByID("2")["content"][2]["contentType"] = %q, want match for %q`, contentList[2]["contentType"], "text")
	}

	expectedText = " knee discomfort and has struggled to find his shot since returning, averaging 5.6 points on 25% shooting and 4.4 assists in his last five games. Gabe Vincent got the start in Saturday's loss at Milwaukee, finishing with seven points."

	if contentList[2]["text"] != expectedText {
		t.Fatalf(`getArticleContentsByID("2")["content"][2]["text"] = %q, want match for %q`, contentList[2]["text"], expectedText)
	}

}
