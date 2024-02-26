package validator

import (
	"fmt"
	"regexp"

	"github.com/AliiAhmadi/email_validator/log"
)

type Validator struct {
	error  error
	status bool
	logger *log.Logger
}

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
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

	v.status = true
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
