package document

import (
	"time"
)

type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

type DocumentChange struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UpdatedBy string `json:"updatedBy"`
}

func NewDocument(t, b, u, id string) *Document {
	now := time.Now()
	return &Document{
		ID:        id,
		Title:     t,
		Body:      b,
		CreatedBy: u,
		UpdatedBy: u,
		Created:   now,
		Updated:   now,
	}
}

func (d *Document) UpdateDocument(dc *DocumentChange) {
	var change bool
	if d.Title != dc.Title {
		d.Title = dc.Title
		change = true
	}

	if d.Body != dc.Body {
		d.Body = dc.Body
		change = true
	}

	if change {
		d.UpdatedBy = dc.UpdatedBy
		d.Updated = time.Now()
	}
}
