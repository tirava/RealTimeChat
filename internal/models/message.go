package models

// MessageType base type.
type MessageType string

// Constants.
const (
	MTPing    MessageType = "ping"
	MTPong    MessageType = "pong"
	MTMessage MessageType = "message"
)

// Message base struct.
type Message struct {
	Type MessageType `json:"type"`
	Data string      `json:"data,omitempty"`
}
