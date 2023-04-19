package main

import (
	"testing"
)

func TestWebScrapePlayerStats(t *testing.T) {
	playerName := "Lionel Messi"
	stats := webScrapePlayerStats(playerName)

	// Check that stats is not empty
	if len(stats) == 0 {
		t.Error("Expected stats to be non-empty")
	}

	// Check that stats contains expected keys
	expectedKeys := []string{"Nation", "Matches Played", "Total Clean Sheets", "Save Percentage"}
	for _, key := range expectedKeys {
		_, ok := stats[key]
		if !ok {
			t.Errorf("Expected stats to contain key %q", key)
		}
	}
}

func TestGetPositions(t *testing.T) {
	// Test with a player who has a known position
	fbrefURL := "https://fbref.com/en/search/search.fcgi?search=Lionel+Messi"
	expectedPositions := "FW"

	positions := getPositions(fbrefURL)

	if positions != expectedPositions {
		t.Errorf("Expected positions to be %q but got %q", expectedPositions, positions)
	}

	// Test with a player who has no position listed
	fbrefURL = "https://fbref.com/en/search/search.fcgi?search=Nonexistent+Player"
	expectedPositions = ""

	positions = getPositions(fbrefURL)

	if positions != expectedPositions {
		t.Errorf("Expected positions to be %q but got %q", expectedPositions, positions)
	}
}

func TestGetGoalkeeperStats(t *testing.T) {
	fbrefURL := "https://fbref.com/en/players/2a2e00b8/Jordan-Pickford"
	expectedStats := map[string]string{
		"Nation":             "England",
		"Matches Played":     "156",
		"Total Clean Sheets": "49",
		"Save Percentage":    "65.7",
	}

	stats := getGoalkeeperStats(fbrefURL, make(map[string]string))

	for key, value := range expectedStats {
		if stats[key] != value {
			t.Errorf("Expected %s to be %s but got %s", key, value, stats[key])
		}
	}
}

func TestGetDefenderStats(t *testing.T) {
	fbrefURL := "https://fbref.com/en/players/8b2c9269/Aaron-Wan-Bissaka"
	expectedStats := map[string]string{
		"Nation":         "England",
		"Matches Played": "111",
	}

	stats := getDefenderStats(fbrefURL, make(map[string]string))

	for key, value := range expectedStats {
		if stats[key] != value {
			t.Errorf("Expected %s to be %s but got %s", key, value, stats[key])
		}
	}
}
