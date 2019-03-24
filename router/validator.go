package messageValidator

import (
	"net/url"
)

type MessageValidator struct {
	FromName         string `binding:"required" json:"fromName"`
	FromEmail        string `binding:"required" json:"fromEmail"`
	ToName           string `binding:"required" json:"toName"`
	ToEmail          string `binding:"required" json:"toEmail"`
	Subject          string `binding:"required" json:"subject"`
	PlainTextContent string `binding:"required" json:"text"`
	HtmlContent      string `json:"htmlContent"`
}

func Validate(a *MessageValidator) url.Values {
	errs := url.Values{}

	// check if the title empty
	if a.FromName == "" {
		errs.Add("fromName", "FromName not provided. Make sure you have a 'fromName' property in your request body.")
	}
	if a.FromEmail == "" {
		errs.Add("fromEmail", "FromEmail not provided. Make sure you have a 'fromEmail' property in your request body.")
	}
	if a.ToName == "" {
		errs.Add("toName", "ToName not provided. Make sure you have a 'toName' property in your request body.")
	}
	if a.ToEmail == "" {
		errs.Add("toEmail", "ToEmail not provided. Make sure you have a 'toEmail' property in your request body.")
	}
	if a.Subject == "" {
		errs.Add("subject", "subject not provided. Make sure you have a 'subject' property in your request body.")
	}
	if a.PlainTextContent == "" {
		errs.Add("text", "text not provided. Make sure you have a 'text' property in your request body.")
	}

	return errs
}
