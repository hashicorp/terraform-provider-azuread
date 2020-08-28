package graph

import "fmt"

type AlreadyExistsError struct {
	obj string
	id  string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s with ID %q already exists", e.obj, e.id)
}
