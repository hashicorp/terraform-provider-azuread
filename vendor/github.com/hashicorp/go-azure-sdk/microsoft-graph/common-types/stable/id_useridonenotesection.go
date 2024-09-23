package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteSectionId{}

// UserIdOnenoteSectionId is a struct representing the Resource ID for a User Id Onenote Section
type UserIdOnenoteSectionId struct {
	UserId           string
	OnenoteSectionId string
}

// NewUserIdOnenoteSectionID returns a new UserIdOnenoteSectionId struct
func NewUserIdOnenoteSectionID(userId string, onenoteSectionId string) UserIdOnenoteSectionId {
	return UserIdOnenoteSectionId{
		UserId:           userId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseUserIdOnenoteSectionID parses 'input' into a UserIdOnenoteSectionId
func ParseUserIdOnenoteSectionID(input string) (*UserIdOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteSectionIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteSectionIDInsensitively(input string) (*UserIdOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateUserIdOnenoteSectionID checks that 'input' can be parsed as a User Id Onenote Section ID
func ValidateUserIdOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Section ID
func (id UserIdOnenoteSectionId) ID() string {
	fmtString := "/users/%s/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Section ID
func (id UserIdOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this User Id Onenote Section ID
func (id UserIdOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("User Id Onenote Section (%s)", strings.Join(components, "\n"))
}
