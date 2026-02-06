package mailer

import (
	"fmt"

	"github.com/emersion/go-smtp"
)

func SendByGoSMTPPackage(m *Message) error {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", m.Host, m.Port))
	if err != nil {
		return err
	}

	// Set the sender and recipient first
	if err := c.Mail(m.Sender, nil); err != nil {
		return err
	}
	if err := c.Rcpt(m.Receiver, nil); err != nil {
		return err
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(wc, m.Data)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return err
	}
	return nil
}
