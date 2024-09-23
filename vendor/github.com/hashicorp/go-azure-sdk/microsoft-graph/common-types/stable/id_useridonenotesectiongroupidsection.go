package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteSectionGroupIdSectionId{}

// UserIdOnenoteSectionGroupIdSectionId is a struct representing the Resource ID for a User Id Onenote Section Group Id Section
type UserIdOnenoteSectionGroupIdSectionId struct {
	UserId           string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewUserIdOnenoteSectionGroupIdSectionID returns a new UserIdOnenoteSectionGroupIdSectionId struct
func NewUserIdOnenoteSectionGroupIdSectionID(userId string, sectionGroupId string, onenoteSectionId string) UserIdOnenoteSectionGroupIdSectionId {
	return UserIdOnenoteSectionGroupIdSectionId{
		UserId:           userId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserIdOnenoteSectionGroupIdSectionID parses 'input' into a UserIdOnenoteSectionGroupIdSectionId
func ParseUserIdOnenoteSectionGroupIdSectionID(input string) (*UserIdOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteSectionGroupIdSectionIDInsensitively(input string) (*UserIdOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateUserIdOnenoteSectionGroupIdSectionID checks that 'input' can be parsed as a User Id Onenote Section Group Id Section ID
func ValidateUserIdOnenoteSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Section Group Id Section ID
func (id UserIdOnenoteSectionGroupIdSectionId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Section Group Id Section ID
func (id UserIdOnenoteSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this User Id Onenote Section Group Id Section ID
func (id UserIdOnenoteSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Id Onenote Section Group Id Section (%s)", strings.Join(components, "\n"))
}
