/*
 * send email
 *
 */
package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
)

type Mailer struct {
	From    string
	To      string
	Subject string
	Body    string

	Password string
	Address  string
}

/*
 * app at addr, switches to TLS if possible,
 * authenticates with mechanism a if possible, and then sends an email from
 * address from, to addresses to, with message msg.
 */
func (m *Mailer) SendMail() error {

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = m.From
	headers["To"] = m.To
	headers["Subject"] = m.Subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + m.Body

	// Connect to the SMTP Server

	host, _, _ := net.SplitHostPort(m.Address)

	auth := smtp.PlainAuth("", m.From, m.Password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(m.Address)
	if err != nil {
		log.Panic(err)
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(m.From); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(m.To); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	return c.Quit()
}
