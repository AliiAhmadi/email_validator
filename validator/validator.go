package validator

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/AliiAhmadi/email_validator/log"
)

type Validator struct {
	error  error
	status bool
	logger *log.Logger
}

var (
	//
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	FROM_EMAIL = "aliahmadi@ut.ac.ir"
)

func New() *Validator {
	return &Validator{
		error:  nil,
		status: false,
		logger: log.New(),
	}
}

func (v *Validator) Email(email string) {
	if err := v.syntax(email, EmailRX); err != nil {
		v.logger.Warning(err)
		v.status = false
		return
	}

	// Make a buffer for reading response from SMTP server
	buf := make([]byte, 1024)

	// Extract parts of email address
	parts := strings.Split(email, "@")

	// DNS MX lookup to find the SMTP server for the recipient's domain
	mx, err := net.LookupMX(parts[1])
	if err != nil {
		v.error = fmt.Errorf("error finding MX records: %w", err)
		return
	}

	// Connecting to SMTP server
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", mx[0].Host, 25), 5*time.Second)
	if err != nil {
		v.logger.Warning(fmt.Errorf("invalid host: %s", mx[0].Host))
		return
	}
	defer conn.Close()

	// Read initial response from the SMTP server
	_, err = conn.Read(buf)
	if err != nil {
		v.error = fmt.Errorf("error reading response from SMTP server: %w", err)
		return
	}

	// Send EHLO command
	conn.Write([]byte("EHLO lexurr.ir\r\n"))
	_, err = conn.Read(buf)
	if err != nil {
		v.error = fmt.Errorf("error reading response from SMTP server: %w", err)
		return
	}

	// Send MAIL FROM command
	conn.Write([]byte("MAIL FROM:<" + FROM_EMAIL + ">\r\n"))
	_, err = conn.Read(buf)
	if err != nil {
		v.error = fmt.Errorf("error reading response from SMTP server: %w", err)
		return
	}

	// Send RCPT TO command with the email address to check
	conn.Write([]byte("RCPT TO:<" + email + ">\r\n"))
	_, err = conn.Read(buf)
	if err != nil {
		v.error = fmt.Errorf("error reading response from SMTP server: %w", err)
		return
	}

	// Check the response code to determine if the email exists
	if strings.Contains(string(buf), "250") {
		// Email exists
		v.status = true
	}
}

func (v *Validator) syntax(email string, rx *regexp.Regexp) error {
	if !rx.MatchString(email) {
		return fmt.Errorf("invalid syntax: %s", email)
	}

	return nil
}

func (v *Validator) Valid() error {
	return v.error
}

func (v *Validator) Status() bool {
	return v.status
}
