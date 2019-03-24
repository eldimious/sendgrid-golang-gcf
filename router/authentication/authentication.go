package authentication

import (
	"errors"
	"net/http"

	"github.com/eldimious/sendgrid-golang-gcf/config"
)

func Authenticate(r *http.Request, config *config.FaaS) error {
	apiKey := r.Header.Get("api-key")
	if apiKey == "" || apiKey != config.APIKey {
		err := errors.New("API key not provided. Make sure you have a 'api-key' as header.")
		return err
	}

	return nil
}
