# email
Send Emails in GO 
simple package that makes sending emails in go easy.

```go
package main

import (
    "log"
    "github.com/jakobii/email"
)

func main () {
    // reusable auth
    conn := &SMTPAuth{
        Server:   "smtp.gmail.com",
    	Port:     587,
    	Username: "user@gmail.com",
    	Password: "p@$$w0rd",
    }


    from := "user@gmail.com"
    to := []string{"user2@gmail.com","user3@gmail.com"}
    subject = "test passed!"
    body = "<h1 style=\"color:green;\">This Test Passed!!!</h1>"


    // send email
    err = conn.Send(from, to, subject, body, email.HTML)
    if err != nil {
        log.Fatal(err)
    }
}
```