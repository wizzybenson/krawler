package types

import "strings"

type EmailSet map[string]bool

func (e *EmailSet) Add(emails []string) {
	for _, email := range emails {
		if _, ok := (*e)[email]; !ok {
			(*e)[email] = true
		}
	}

}

func (e EmailSet) ToString() string {
	emails := []string{}
	for email := range e {
		emails = append(emails, email)
	}

	return strings.Join(emails, " ")
}
