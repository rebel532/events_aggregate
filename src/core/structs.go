package core

type Event struct {
	UserId    int    `json:"userId"`
	EventType string `json:"eventType"`
	Timestamp int64  `json:"timestamp"`
}

type Report struct {
	UserId      int
	Date        string
	EventCounts map[string]int
}
