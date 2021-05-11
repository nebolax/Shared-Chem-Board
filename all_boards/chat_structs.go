package all_boards

type ChatHistory []ChatMessage

type TimeStamp struct {
	Year   uint64 `json:"year"`
	Month  uint64 `json:"month"`
	Day    uint64 `json:"day"`
	Hour   uint64 `json:"hour"`
	Minute uint64 `json:"minute"`
}

type ChatContent struct {
	Text string `json:"text"`
}

type ChatMessage struct {
	ID        uint64
	SenderID  uint64
	TimeStamp TimeStamp
	Content   ChatContent
}
