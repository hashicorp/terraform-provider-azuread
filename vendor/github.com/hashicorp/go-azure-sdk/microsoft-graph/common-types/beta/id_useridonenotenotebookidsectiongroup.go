package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookIdSectionGroupId{}

// UserIdOnenoteNotebookIdSectionGroupId is a struct representing the Resource ID for a User Id Onenote Notebook Id Section Group
type UserIdOnenoteNotebookIdSectionGroupId struct {
	UserId         string
	NotebookId     string
	SectionGroupId string
}

// NewUserIdOnenoteNotebookIdSectionGroupID returns a new UserIdOnenoteNotebookIdSectionGroupId struct
func NewUserIdOnenoteNotebookIdSectionGroupID(userId string, notebookId string, sectionGroupId string) UserIdOnenoteNotebookIdSectionGroupId {
	return UserIdOnenoteNotebookIdSectionGroupId{
		UserId:         userId,
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseUserIdOnenoteNotebookIdSectionGroupID parses 'input' into a UserIdOnenoteNotebookIdSectionGroupId
func ParseUserIdOnenoteNotebookIdSectionGroupID(input string) (*UserIdOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteNotebookIdSectionGroupIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteNotebookIdSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteNotebookIdSectionGroupIDInsensitively(input string) (*UserIdOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteNotebookIdSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdOnenoteNotebookIdSectionGroupID checks that 'input' can be parsed as a User Id Onenote Notebook Id Section Group ID
func ValidateUserIdOnenoteNotebookIdSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteNotebookIdSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Notebook Id Section Group ID
func (id UserIdOnenoteNotebookIdSectionGroupId) ID() string {
	fmtString := "/users/%s/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Notebook Id Section Group ID
func (id UserIdOnenoteNotebookIdSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this User Id Onenote Notebook Id Section Group ID
func (id UserIdOnenoteNotebookIdSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("User Id Onenote Notebook Id Section Group (%s)", strings.Join(components, "\n"))
}
