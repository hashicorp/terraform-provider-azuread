package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookIdSectionId{}

// UserIdOnenoteNotebookIdSectionId is a struct representing the Resource ID for a User Id Onenote Notebook Id Section
type UserIdOnenoteNotebookIdSectionId struct {
	UserId           string
	NotebookId       string
	OnenoteSectionId string
}

// NewUserIdOnenoteNotebookIdSectionID returns a new UserIdOnenoteNotebookIdSectionId struct
func NewUserIdOnenoteNotebookIdSectionID(userId string, notebookId string, onenoteSectionId string) UserIdOnenoteNotebookIdSectionId {
	return UserIdOnenoteNotebookIdSectionId{
		UserId:           userId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserIdOnenoteNotebookIdSectionID parses 'input' into a UserIdOnenoteNotebookIdSectionId
func ParseUserIdOnenoteNotebookIdSectionID(input string) (*UserIdOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteNotebookIdSectionIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteNotebookIdSectionId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteNotebookIdSectionIDInsensitively(input string) (*UserIdOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteNotebookIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateUserIdOnenoteNotebookIdSectionID checks that 'input' can be parsed as a User Id Onenote Notebook Id Section ID
func ValidateUserIdOnenoteNotebookIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteNotebookIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Notebook Id Section ID
func (id UserIdOnenoteNotebookIdSectionId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Notebook Id Section ID
func (id UserIdOnenoteNotebookIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this User Id Onenote Notebook Id Section ID
func (id UserIdOnenoteNotebookIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Id Onenote Notebook Id Section (%s)", strings.Join(components, "\n"))
}
