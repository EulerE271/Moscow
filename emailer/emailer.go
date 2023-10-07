package emailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"russianwords/tsvreader"
	"strings"
)

func SendMail(from string, password string, to []string, subject string, body string) error {
	smtpHost := "smtp-relay.brevo.com"
	smtpPort := "587"

	// Set up authentication information
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Connect to the server, switch to TLS, and then authenticate
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true, // This is for testing; ideally, you'd want this to be false
		ServerName:         smtpHost,
	}

	conn, err := smtp.Dial(smtpHost + ":" + smtpPort)
	if err != nil {
		return err
	}

	// Start TLS
	if err = conn.StartTLS(tlsconfig); err != nil {
		return err
	}

	// Authenticate
	if err = conn.Auth(auth); err != nil {
		return err
	}

	// Set the sender and recipient
	if err = conn.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = conn.Rcpt(addr); err != nil {
			return err
		}
	}

	// Send the email body
	headers := make(map[string]string)
	headers["From"] = "noreply@glossaryru.com"
	headers["To"] = to[0] // assuming you're sending to one recipient for simplicity; adjust if needed
	headers["Subject"] = subject
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"UTF-8\""

	email := ""
	for k, v := range headers {
		email += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	email += "\r\n" + body

	// Send the email body
	w, err := conn.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(email))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	// Close the connection
	conn.Quit()

	return nil
}

func SendGlossary(email string, glossaryCount int) error {
	// Fetch required info from environment variables
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	subject := "Your Requested Glossary"

	// Get random glossary entries from TSV file
	// Define n as per your requirement of number of lines you want to fetch
	n := glossaryCount
	records, err := tsvreader.ReadRandomLines("russianwords.tsv", n)
	if err != nil {
		return err // You might handle this error differently in production
	}

	// Convert the records to a formatted string
	var lines []string
	for _, record := range records {
		lines = append(lines, strings.Join(record, " - "))
	}
	glossaryContent := strings.Join(lines, "\n")

	// Form the email body
	body := "Here is the glossary you requested:\n\n" + glossaryContent

	// Call SendMail with the above values
	return SendMail(from, password, []string{email}, subject, body)
}
