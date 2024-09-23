// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type DirectoryRoleId struct {
	DirectoryRoleId string
}

func NewDirectoryRoleID(input string) DirectoryRoleId {
	return DirectoryRoleId{DirectoryRoleId: input}
}

func (id DirectoryRoleId) ID() string {
	return id.DirectoryRoleId
}

func (id DirectoryRoleId) String() string {
	return fmt.Sprintf("Directory Role (Object ID: %q)", id.DirectoryRoleId)
}
