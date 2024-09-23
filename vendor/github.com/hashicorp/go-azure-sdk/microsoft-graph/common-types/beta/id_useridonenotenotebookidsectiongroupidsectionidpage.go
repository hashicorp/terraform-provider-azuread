package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}

// UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a User Id Onenote Notebook Id Section Group Id Section Id Page
type UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId struct {
	UserId           string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID returns a new UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId struct
func NewUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(userId string, notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId {
	return UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{
		UserId:           userId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID parses 'input' into a UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId
func ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(input string) (*UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively(input string) (*UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a User Id Onenote Notebook Id Section Group Id Section Id Page ID
func ValidateUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this User Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id UserIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Id Onenote Notebook Id Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
