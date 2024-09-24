package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookId{}

// UserIdOnenoteNotebookId is a struct representing the Resource ID for a User Id Onenote Notebook
type UserIdOnenoteNotebookId struct {
	UserId     string
	NotebookId string
}

// NewUserIdOnenoteNotebookID returns a new UserIdOnenoteNotebookId struct
func NewUserIdOnenoteNotebookID(userId string, notebookId string) UserIdOnenoteNotebookId {
	return UserIdOnenoteNotebookId{
		UserId:     userId,
		NotebookId: notebookId,
	}
}

// ParseUserIdOnenoteNotebookID parses 'input' into a UserIdOnenoteNotebookId
func ParseUserIdOnenoteNotebookID(input string) (*UserIdOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteNotebookIDInsensitively(input string) (*UserIdOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteNotebookId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	return nil
}

// ValidateUserIdOnenoteNotebookID checks that 'input' can be parsed as a User Id Onenote Notebook ID
func ValidateUserIdOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Notebook ID
func (id UserIdOnenoteNotebookId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Notebook ID
func (id UserIdOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
	}
}

// String returns a human-readable description of this User Id Onenote Notebook ID
func (id UserIdOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("User Id Onenote Notebook (%s)", strings.Join(components, "\n"))
}
