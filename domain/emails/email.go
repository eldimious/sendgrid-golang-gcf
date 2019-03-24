package email

//struct contains information about a email.
type Message struct {
	PlainTextContent string
	HtmlContent      string
	Subject          string
}

type From struct {
	Name  string
	Email string
}

type To struct {
	Name  string
	Email string
}
