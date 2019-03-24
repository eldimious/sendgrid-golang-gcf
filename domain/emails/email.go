package email

//struct contains information about a email.
type Message struct {
	PlainTextContent string
	HtmlContent      string
	Subject          string
}

type Sender struct {
	Name  string
	Email string
}

type Receiver struct {
	Name  string
	Email string
}
