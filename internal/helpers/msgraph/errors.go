package msgraph

import "fmt"

type AlreadyExistsError struct {
	obj string
	id  string
}

func (e AlreadyExistsError) Error() string {
	return fmt.Sprintf("An existing %s already exists with the ID %q", e.obj, e.id)
}

type CredentialError struct {
	str  string
	attr string
}

func (e CredentialError) Attr() string {
	return e.attr
}

func (e CredentialError) Error() string {
	return e.str
}
