package types

import "strings"

type PhoneSet map[string]bool

func (e *PhoneSet) Add(phones []string) {
	for _, phone := range phones {
		if _, ok := (*e)[phone]; !ok {
			(*e)[phone] = true
		}
	}

}

func (e PhoneSet) ToString() string {
	phones := []string{}
	for phone := range e {
		phones = append(phones, phone)
	}

	return strings.Join(phones, ", ")
}
