package model

// Message - structure of email message, that
// will be sent using email
type Message struct {
	Receiver string
	Body     string
	Subject  string
}
