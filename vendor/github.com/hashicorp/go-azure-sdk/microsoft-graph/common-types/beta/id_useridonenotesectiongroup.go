package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteSectionGroupId{}

// UserIdOnenoteSectionGroupId is a struct representing the Resource ID for a User Id Onenote Section Group
type UserIdOnenoteSectionGroupId struct {
	UserId         string
	SectionGroupId string
}

// NewUserIdOnenoteSectionGroupID returns a new UserIdOnenoteSectionGroupId struct
func NewUserIdOnenoteSectionGroupID(userId string, sectionGroupId string) UserIdOnenoteSectionGroupId {
	return UserIdOnenoteSectionGroupId{
		UserId:         userId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseUserIdOnenoteSectionGroupID parses 'input' into a UserIdOnenoteSectionGroupId
func ParseUserIdOnenoteSectionGroupID(input string) (*UserIdOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteSectionGroupIDInsensitively(input string) (*UserIdOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	return nil
}

// ValidateUserIdOnenoteSectionGroupID checks that 'input' can be parsed as a User Id Onenote Section Group ID
func ValidateUserIdOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Section Group ID
func (id UserIdOnenoteSectionGroupId) ID() string {
	fmtString := "/users/%s/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Section Group ID
func (id UserIdOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this User Id Onenote Section Group ID
func (id UserIdOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("User Id Onenote Section Group (%s)", strings.Join(components, "\n"))
}
