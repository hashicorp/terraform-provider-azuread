package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookIdSectionGroupIdSectionId{}

// UserIdOnenoteNotebookIdSectionGroupIdSectionId is a struct representing the Resource ID for a User Id Onenote Notebook Id Section Group Id Section
type UserIdOnenoteNotebookIdSectionGroupIdSectionId struct {
	UserId           string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewUserIdOnenoteNotebookIdSectionGroupIdSectionID returns a new UserIdOnenoteNotebookIdSectionGroupIdSectionId struct
func NewUserIdOnenoteNotebookIdSectionGroupIdSectionID(userId string, notebookId string, sectionGroupId string, onenoteSectionId string) UserIdOnenoteNotebookIdSectionGroupIdSectionId {
	return UserIdOnenoteNotebookIdSectionGroupIdSectionId{
		UserId:           userId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserIdOnenoteNotebookIdSectionGroupIdSectionID parses 'input' into a UserIdOnenoteNotebookIdSectionGroupIdSectionId
func ParseUserIdOnenoteNotebookIdSectionGroupIdSectionID(input string) (*UserIdOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteNotebookIdSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIDInsensitively(input string) (*UserIdOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteNotebookIdSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateUserIdOnenoteNotebookIdSectionGroupIdSectionID checks that 'input' can be parsed as a User Id Onenote Notebook Id Section Group Id Section ID
func ValidateUserIdOnenoteNotebookIdSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteNotebookIdSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Notebook Id Section Group Id Section ID
func (id UserIdOnenoteNotebookIdSectionGroupIdSectionId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Notebook Id Section Group Id Section ID
func (id UserIdOnenoteNotebookIdSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this User Id Onenote Notebook Id Section Group Id Section ID
func (id UserIdOnenoteNotebookIdSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Id Onenote Notebook Id Section Group Id Section (%s)", strings.Join(components, "\n"))
}
