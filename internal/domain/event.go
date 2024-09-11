package domain

import "time"

type EventMessage struct {
	UUID      string    `json:"uid"`
	Context   string    `json:"context"`
	EventType string    `json:"event_type"`
	Client    string    `json:"client"`
	CreatedAt time.Time `json:"created_at"`
}
