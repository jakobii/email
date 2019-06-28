# Email

A simple package that makes sending emails in GO easy and straight forward.

```bash
go get github.com/jakobii/email
```

## Example

```go
package main

import (
	"github.com/jakobii/email"
)

func main() {

	// reusable authentication
	auth := email.Auth{
		Server:   "smtp.gmail.com",
		Port:     587,
		Username: "user@gmail.com",
		Password: "p@$$w0rd",
	}

	// a single email message
	msg := email.Message{
		From:    "user@gmail.com",
		To:      []string{"user2@gmail.com", "user3@gmail.com"},
		Subject: "test passed!",

		// note that the Body's ContentType defaults to email.HTML
		//BodyType: email.Plain or email.HTML
		Body:    []byte("<h1 style=\"color:green;\">This Test Passed!!!</h1>"),
	}

	// send email
	err := auth.Send(msg)
	if err != nil {
		panic(err)
	}
}
```