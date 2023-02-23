package meme

import (
	"time"
)

type Origin = string

const (
	Generated Origin = "generated"
	Uploaded  Origin = "uploaded"
)

type Meme struct {
	Url         string
	DateCreated time.Time
	Text        string
	Origin      Origin
}

func NewMeme(url, text string, origin Origin) Meme {
	return Meme{
		Url:         url,
		DateCreated: time.Now(),
		Text:        text,
		Origin:      origin,
	}
}

func (m Meme) AsSerializable() SerializableMeme {
	return SerializableMeme{
		Url:         m.Url,
		DateCreated: int(m.DateCreated.Unix()),
		Text:        m.Text,
		Origin:      m.Origin,
	}
}

type SerializableMeme struct {
	Url         string `json:"url" firestore:"url,omitempty"`
	DateCreated int    `json:"dateCreated" firestore:"dateCreated,omitempty"`
	Text        string `json:"text" firestore:"text,omitempty"`
	Origin      string `json:"origin" firestore:"origin,omitempty"`
}
