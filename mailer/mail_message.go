package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

type Message struct {
	Sender   string
	Receiver string
	Host     string
	Port     int
	Data     string
}

type MessageBuilder struct {
	template     *template.Template
	templateData map[string]string
	content      string
	header       map[string]string
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		header: map[string]string{},
	}
}

func (b *MessageBuilder) UseTemplate(name string, data map[string]string) *MessageBuilder {
	templatePath := "templates/" + name
	emailTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		emailTemplate, _ = template.New("default").Parse(`
			<!DOCTYPE html>
			<html>
			<body>
				<h1>{{.Title}}</h1>
				<p>{{.Message}}</p>
			</body>
			</html>
		`)
	}
	b.template = emailTemplate
	b.templateData = data
	b.header["Content-Type"] = "text/html"
	return b
}

func (b *MessageBuilder) AddHeader(key, value string) *MessageBuilder {
	b.header[key] = value
	return b
}

func (b *MessageBuilder) SetContent(rawContent string) *MessageBuilder {
	b.content = rawContent
	return b
}

func (b *MessageBuilder) Build() ([]byte, error) {
	if b.content == "" && b.template == nil {
		return nil, errors.New("no body set")
	}

	mimeHeaders := fmt.Sprintf("MIME-version: 1.0;\nContent-Type: %s; charset=\"UTF-8\";\n\n", b.header["Content-Type"])
	from := fmt.Sprintf("From: %s\r", b.header["from"])
	to := fmt.Sprintf("To: %s\r", b.header["to"])
	subject := fmt.Sprintf("Subject: %s\r", b.header["subject"])

	header := from + "\n" + to + "\n" + subject + "\n" + mimeHeaders + "\n\n"

	if b.template != nil {
		var body bytes.Buffer
		body.Write([]byte(header))
		if err := b.template.Execute(&body, b.templateData); err != nil {
			return nil, err
		}
		return body.Bytes(), nil
	}

	return []byte(header + b.content), nil
}
