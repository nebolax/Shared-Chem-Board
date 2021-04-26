package all_boards

type ChatHistory []ChatMessage

type TimeStamp struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type ChatContent struct {
	Text string `json:"text"`
}

type ChatMessage struct {
	ID        int         `json:"id"`
	SenderID  int         `json:"senderid"`
	TimeStamp TimeStamp   `json:"timestamp"`
	Content   ChatContent `json:"content"`
}

type Chat struct {
	ID      int
	History ChatHistory
}
