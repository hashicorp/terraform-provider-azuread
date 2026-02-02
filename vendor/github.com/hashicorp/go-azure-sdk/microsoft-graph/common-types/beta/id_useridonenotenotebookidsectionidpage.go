package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookIdSectionIdPageId{}

// UserIdOnenoteNotebookIdSectionIdPageId is a struct representing the Resource ID for a User Id Onenote Notebook Id Section Id Page
type UserIdOnenoteNotebookIdSectionIdPageId struct {
	UserId           string
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserIdOnenoteNotebookIdSectionIdPageID returns a new UserIdOnenoteNotebookIdSectionIdPageId struct
func NewUserIdOnenoteNotebookIdSectionIdPageID(userId string, notebookId string, onenoteSectionId string, onenotePageId string) UserIdOnenoteNotebookIdSectionIdPageId {
	return UserIdOnenoteNotebookIdSectionIdPageId{
		UserId:           userId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserIdOnenoteNotebookIdSectionIdPageID parses 'input' into a UserIdOnenoteNotebookIdSectionIdPageId
func ParseUserIdOnenoteNotebookIdSectionIdPageID(input string) (*UserIdOnenoteNotebookIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteNotebookIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteNotebookIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteNotebookIdSectionIdPageIDInsensitively(input string) (*UserIdOnenoteNotebookIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteNotebookIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateUserIdOnenoteNotebookIdSectionIdPageID checks that 'input' can be parsed as a User Id Onenote Notebook Id Section Id Page ID
func ValidateUserIdOnenoteNotebookIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteNotebookIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Notebook Id Section Id Page ID
func (id UserIdOnenoteNotebookIdSectionIdPageId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Notebook Id Section Id Page ID
func (id UserIdOnenoteNotebookIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this User Id Onenote Notebook Id Section Id Page ID
func (id UserIdOnenoteNotebookIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Id Onenote Notebook Id Section Id Page (%s)", strings.Join(components, "\n"))
}
