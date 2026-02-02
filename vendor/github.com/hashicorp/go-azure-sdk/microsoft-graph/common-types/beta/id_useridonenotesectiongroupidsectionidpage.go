package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteSectionGroupIdSectionIdPageId{}

// UserIdOnenoteSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a User Id Onenote Section Group Id Section Id Page
type UserIdOnenoteSectionGroupIdSectionIdPageId struct {
	UserId           string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserIdOnenoteSectionGroupIdSectionIdPageID returns a new UserIdOnenoteSectionGroupIdSectionIdPageId struct
func NewUserIdOnenoteSectionGroupIdSectionIdPageID(userId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) UserIdOnenoteSectionGroupIdSectionIdPageId {
	return UserIdOnenoteSectionGroupIdSectionIdPageId{
		UserId:           userId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserIdOnenoteSectionGroupIdSectionIdPageID parses 'input' into a UserIdOnenoteSectionGroupIdSectionIdPageId
func ParseUserIdOnenoteSectionGroupIdSectionIdPageID(input string) (*UserIdOnenoteSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteSectionGroupIdSectionIdPageIDInsensitively(input string) (*UserIdOnenoteSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateUserIdOnenoteSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a User Id Onenote Section Group Id Section Id Page ID
func ValidateUserIdOnenoteSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Section Group Id Section Id Page ID
func (id UserIdOnenoteSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Section Group Id Section Id Page ID
func (id UserIdOnenoteSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this User Id Onenote Section Group Id Section Id Page ID
func (id UserIdOnenoteSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Id Onenote Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
