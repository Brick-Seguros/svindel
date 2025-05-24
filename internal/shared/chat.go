package shared

import "time"

// --- User Message ---

type UserMessage struct {
	ID          string // Unique ID for tracking
	Content     string // User input text
	CreatedAt   time.Time
	ReceivedAt  time.Time
	ProcessedAt time.Time
}

// --- AI Response ---

type AIMessageType string

const (
	AIMessageTypeText             AIMessageType = "TEXT"
	AIMessageTypeReportShortcut   AIMessageType = "REPORT_SHORTCUT"
	AIMessageTypeAgentTrigger     AIMessageType = "AGENT_TRIGGER"
	AIMessageTypeResourceSelector AIMessageType = "RESOURCE_SELECTOR"
)

type AIMessage struct {
	Type      AIMessageType
	Text      string
	Shortcut  Shortcut
	Resources []Resource
	// - TEXT: slice of texts
	// - REPORT_SHORTCUT: slice with Report IDs
	// - AGENT_TRIGGER: slice with Agent IDs
	// - RESOURCE_SELECTOR: slice with Resource IDs
}

type Shortcut struct {
	ID        string
	Title     string
	Document  string
	CreatedAt string
}

type Resource struct {
	ID         string
	Title      string
	HelperText string
}

type AIResponse struct {
	ID          string
	Messages    []AIMessage
	Liked       *bool // nil = not rated, true/false = like/dislike
	CreatedAt   time.Time
	ReceivedAt  time.Time
	ProcessedAt time.Time
}

// --- Chat Session ---

type Chat struct {
	ID           string
	Title        string
	UserMessages []UserMessage
	AIResponses  []AIResponse
	CreatedAt    time.Time
}
