package email

import (
	"bytes"
	"errors"
	"net/smtp"
	"strconv"
	"strings"
)

// ContentType defines the type of body the email will have.
type ContentType string

// maybe there are more email content types out there.. idk..
const (
	HTML  ContentType = "text/html"
	Plain ContentType = "text/plain"
)

func (t ContentType) String() string {
	return string(t)
}

// Message represents a single email message
type Message struct {
	From     string
	To       []string
	Cc       []string
	Bcc      []string
	Subject  string
	Body     string
	BodyType ContentType
}

// Auth provide basic smtp parameter
type Auth struct {
	Server   string
	Port     int
	Username string
	Password string
}

// Send method is a more realistic approach to sending emails
// authentication is packages seperately in a struct to be reused.
func (a *Auth) Send(m Message) error {
	err := Send(a.Server, a.Port, a.Username, a.Password, m.From, m.To, m.Cc, m.Bcc, m.Subject, m.Body, m.BodyType)
	if err != nil {
		return err
	}
	return nil
}

// Send is a simple function that will send emails
// the parameter list is large but its a one stop solution
// to sending an email.
func Send(Server string, Port int, Username string, Password string, From string, To []string, Cc []string, Bcc []string, Subject string, Body string, Content ContentType) error {

	//err checks
	if len(To) < 1 {
		return errors.New("'To' was left blank")
	}

	Authentication := smtp.PlainAuth("", Username, Password, Server)
	Address := Server + ":" + strconv.FormatInt(int64(Port), 10)

	// message body defines From, To, Subject.
	// apparently this is how email works?
	var body bytes.Buffer
	body.WriteString("From: " + From + "\r\n")
	body.WriteString("To: " + strings.Join(To, ",") + "\r\n")

	if len(Cc) < 1 {
		body.WriteString("Cc: " + strings.Join(Cc, ",") + "\r\n")
	}
	if len(Bcc) < 1 {
		body.WriteString("Bcc: " + strings.Join(Bcc, ",") + "\r\n")
	}

	body.WriteString("Subject: " + Subject + "\r\n")

	if Content.String() == "" {
		Content = HTML
	}

	body.WriteString("MIME-version: 1.0;\nContent-Type: " + Content.String() + "; charset=\"UTF-8\";\n\n")
	body.WriteString("\r\n" + Body + "\r\n")

	err := smtp.SendMail(Address, Authentication, From, To, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
