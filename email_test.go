package email

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSend(t *testing.T) {

	// json mirror
	settings := struct {
		Server    string
		Port      int
		Username  string
		Password  string
		From      string
		To        []string
		Cc        []string
		Bcc       []string
		Subject   string
		Body      string
		PlainText string
	}{}

	// get json settings
	// rename email_test.sample.json -> email_test.json
	jsonFile, err := os.Open("email_test.json")
	if err != nil {
		fmt.Println("rename email_test.sample.json to email_test.json")
		t.Error(err)
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(jsonBytes, &settings)
	if err != nil {
		t.Error(err)
	}

	// reusable auth
	auth := Auth{
		Server:   settings.Server,
		Port:     settings.Port,
		Username: settings.Username,
		Password: settings.Password,
	}

	// one off message
	message := Message{
		From:     settings.From,
		To:       settings.To,
		Cc:       settings.Cc,
		Bcc:      settings.Bcc,
		Subject:  settings.Subject,
		Body:     []byte(settings.Body),
		BodyType: HTML,
	}

	// send email
	err = auth.Send(message)

	if err != nil {
		t.Error(err)
	}

}
