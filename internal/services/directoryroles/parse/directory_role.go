package parse

import "fmt"

type DirectoryRoleId struct {
	val string
}

func NewDirectoryRoleID(input string) DirectoryRoleId {
	return DirectoryRoleId{val: input}
}

func (id DirectoryRoleId) ID() string {
	return id.val
}

func (id DirectoryRoleId) String() string {
	return fmt.Sprintf("Directory Role (Object ID: %q)", id.val)
}
