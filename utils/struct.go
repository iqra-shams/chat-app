package utils
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Receiver string `json:"to"`
}