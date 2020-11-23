package mail

import (
	"bytes"
	"flycode.go/abacusf/app/tools/debug"
	"html/template"
	"net/smtp"
	"strconv"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func (r *Request) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Request) sendMail(address string, port int, userMail string, userPassword string) bool {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	f := address+":"+strconv.Itoa(port)
	err := smtp.SendMail(
		f,
		smtp.PlainAuth(
			"",
			userMail,
			userPassword,
			address),
		userMail,
		r.to,
		[]byte(body))

	if err != nil {
		debug.Error("Failed to send the email : %v\n", err)
		return false
	}
	return true
}

func (r *Request) Send(templateName string, address string, port int, userMail string, userPassword string, items interface{}) {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		debug.Error("Failed to send the email : %v\n", err)
	}
	if ok := r.sendMail(address, port, userMail, userPassword); ok {
		debug.Log("Email has been sent to %s\n", r.to)
	} else {
		debug.Error("Failed to send the email to %s\n", r.to)
	}
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

