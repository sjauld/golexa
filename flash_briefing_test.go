package golexa

import (
	"encoding/json"
	"net/url"
	"testing"
)

func testBriefing() *FlashBriefing {
	return NewFlashBriefing("test", "Beer is yum")
}

func TestNewFlashBriefing(t *testing.T) {
	fb := testBriefing()
	if fb.Title != "test" {
		t.Errorf("Expected title to be test, got %v", fb.Title)
	}
}

func TestAddAudioContent(t *testing.T) {
	urlString := "https://beer.com/mp3.wav"
	fb := testBriefing()

	u, _ := url.Parse(urlString)
	fb.AddAudioContent(u)

	if fb.AudioContent != urlString {
		t.Errorf("Expected fb.DisplayURL to be %v, got %v", urlString, fb.AudioContent)
	}
}

func TestAddDisplayURL(t *testing.T) {
	urlString := "https://beer.com/"
	fb := testBriefing()

	u, _ := url.Parse(urlString)
	fb.AddDisplayURL(u)

	if fb.DisplayURL != urlString {
		t.Errorf("Expected fb.DisplayURL to be %v, got %v", urlString, fb.DisplayURL)
	}
}

func TestMarshalToJSON(t *testing.T) {
	_, err := json.Marshal(testBriefing())
	if err != nil {
		t.Error(err)
	}
}

func TestString(t *testing.T) {
	fb := testBriefing()
	expect := `{"uid":"`
	if fb.String()[0:8] != expect {
		t.Errorf("Expected %s, got %s", expect, fb.String()[0:8])
	}
}
