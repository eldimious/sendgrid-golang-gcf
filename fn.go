package sendEmail

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eldimious/sendgrid-golang-gcf/config"
	dispatcher "github.com/eldimious/sendgrid-golang-gcf/data/dispatcher"
	email "github.com/eldimious/sendgrid-golang-gcf/domain/emails"
	validator "github.com/eldimious/sendgrid-golang-gcf/router"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	data := &validator.MessageValidator{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	validationErrs := validator.Validate(data)
	if len(validationErrs) > 0 {
		err := map[string]interface{}{"validationError": validationErrs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	from := &email.From{
		Name:  data.FromName,
		Email: data.FromEmail,
	}
	to := &email.To{
		Name:  data.ToName,
		Email: data.ToEmail,
	}
	message := &email.Message{
		Subject:          data.Subject,
		PlainTextContent: data.PlainTextContent,
		HtmlContent:      data.HtmlContent,
	}

	configuration, err := config.NewConfig()
	if err != nil {
		log.Println(err.Error())
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	emailDispatcher := dispatcher.New(configuration.SendGrid)
	emailsSvc := email.NewService(emailDispatcher)
	dispatcherError := emailsSvc.SendEmail(from, to, message)
	if dispatcherError != nil {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dispatcherError.Error()))
		log.Println(dispatcherError.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "applciation/json")
	w.Write([]byte("Your message has been sent successfully!"))
	return
}
