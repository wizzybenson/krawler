package types

type Contact struct {
	Emailset *EmailSet
	Phoneset *PhoneSet
}

func NewContact() *Contact {
	return &Contact{&EmailSet{}, &PhoneSet{}}
}