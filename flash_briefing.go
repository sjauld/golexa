package alexa

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
)

// FlashBriefing represents the object as described here
// https://developer.amazon.com/docs/flashbriefing/flash-briefing-skill-api-feed-reference.html#feed-format-details
type FlashBriefing struct {
	Identifier   uuid.UUID `json:"uid"`
	Date         time.Time `json:"updateDate"`
	Title        string    `json:"titleText"`
	AudioContent string    `json:"streamUrl"`
	TextContent  string    `json:"mainText"`
	DisplayURL   string    `json:"redirection-URL"`
}

// NewFlashBriefing returns a pointer to a FlashBriefing
func NewFlashBriefing(title, textContent string) *FlashBriefing {
	fb := FlashBriefing{
		Identifier:  uuid.New(),
		Date:        time.Now(),
		Title:       title,
		TextContent: textContent,
	}

	return &fb
}

// AddAudioContent safely adds a URL pointing at your audio content to your flash briefing
func (f *FlashBriefing) AddAudioContent(u *url.URL) {
	f.AudioContent = u.String()
}

// AddDisplayURL safely adds a URL pointing at your display site to your flash briefing
func (f *FlashBriefing) AddDisplayURL(u *url.URL) {
	f.DisplayURL = u.String()
}

func (f *FlashBriefing) String() string {
	data, _ := json.Marshal(f)
	return fmt.Sprintf("%s", data)
}
