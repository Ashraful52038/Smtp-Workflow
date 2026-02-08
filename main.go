package main

import (
	"fmt"
	"log"
	"smtptest/mailer"
	"strings"
)

func main() {
	fmt.Println("Sending mails to mailpit")

	//mailpit SMTP configuaration
	config := struct {
		Host   string
		Port   int
		Sender string
	}{
		Host:   "127.0.0.1",
		Port:   1025,
		Sender: "test@mailpit.com",
	}

	receivers := "user@mailpit.com"

	// Plain text email
	plainText := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: Plain Text Test\r\nContent-Type: text/plain\r\n\r\n"+
			"This is a plain text email sent to Mailpit.\n"+
			"You can view it at: http://localhost:8025",
		config.Sender, receivers,
	)

	err := mailer.SendByGoSMTPPackage(&mailer.Message{
		Host:     config.Host,
		Port:     config.Port,
		Sender:   config.Sender,
		Receiver: receivers,
		Data:     plainText,
	})

	if err != nil {
		log.Printf("Plain Text Error: %v", err)
	} else {
		log.Printf("Plain Text email send to mailpit")
	}

	//Html Mail
	builder := mailer.NewMessageBuilder()

	htmlBody := `<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial; padding: 20px; }
			.email { border: 2px solid #4CAF50; padding: 20px; border-radius: 10px; }
			.header { color: #4CAF50; }
		</style>
	</head>
	<body>
		<div class="email">
			<h1 class="header">📨 HTML Email in Mailpit</h1>
			<p>This email was sent from <strong>Go program</strong> to <strong>Mailpit SMTP</strong>.</p>
			<p>Check the Mailpit dashboard to see this email.</p>
			<p><a href="http://localhost:8025">Click here to open Mailpit Dashboard</a></p>
		</div>
	</body>
	</html>`

	htmlBytes, err := builder.
		AddHeader("from", config.Sender).
		AddHeader("to", receivers).
		AddHeader("subject", "HTML Email Test").
		SetContent(htmlBody).
		Build()

	if err != nil {
		log.Printf("HTML Build Error: %v", err)
		return
	} else {
		err = mailer.SendByGoSMTPPackage(&mailer.Message{
			Host:     config.Host,
			Port:     config.Port,
			Sender:   config.Sender,
			Receiver: receivers,
			Data:     string(htmlBytes),
		})

		if err != nil {
			log.Printf("HTML Email Send Error: %v", err)
		} else {
			log.Printf("HTML email send to mailpit")
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 EMAILS DISPATCHED TO MAILPIT")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("\n🌐 OPEN MAILPIT DASHBOARD:")
	fmt.Println("   URL: http://localhost:8025")
	fmt.Println("\n📨 You should see:")
	fmt.Println("   1. Plain Text Test")
	fmt.Println("   2. HTML Email Test")
	fmt.Println("\n🔍 In Mailpit dashboard you can:")
	fmt.Println("   • View email content")
	fmt.Println("   • See HTML rendering")
	fmt.Println("   • Check email headers")
	fmt.Println("   • View raw source")
	fmt.Println(strings.Repeat("=", 50))

}
