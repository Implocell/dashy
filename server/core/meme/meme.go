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

type SerializableMeme struct {
	Url         string `json:"url" firestore:"url,omitempty"`
	DateCreated int    `json:"dateCreated" firestore:"dateCreated,omitempty"`
	Text        string `json:"text" firestore:"text,omitempty"`
	Origin      string `json:"origin" firestore:"origin,omitempty"`
}
