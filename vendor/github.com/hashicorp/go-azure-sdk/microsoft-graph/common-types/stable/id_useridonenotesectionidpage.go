package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteSectionIdPageId{}

// UserIdOnenoteSectionIdPageId is a struct representing the Resource ID for a User Id Onenote Section Id Page
type UserIdOnenoteSectionIdPageId struct {
	UserId           string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewUserIdOnenoteSectionIdPageID returns a new UserIdOnenoteSectionIdPageId struct
func NewUserIdOnenoteSectionIdPageID(userId string, onenoteSectionId string, onenotePageId string) UserIdOnenoteSectionIdPageId {
	return UserIdOnenoteSectionIdPageId{
		UserId:           userId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseUserIdOnenoteSectionIdPageID parses 'input' into a UserIdOnenoteSectionIdPageId
func ParseUserIdOnenoteSectionIdPageID(input string) (*UserIdOnenoteSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteSectionIdPageIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteSectionIdPageIDInsensitively(input string) (*UserIdOnenoteSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateUserIdOnenoteSectionIdPageID checks that 'input' can be parsed as a User Id Onenote Section Id Page ID
func ValidateUserIdOnenoteSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Section Id Page ID
func (id UserIdOnenoteSectionIdPageId) ID() string {
	fmtString := "/users/%s/onenote/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Section Id Page ID
func (id UserIdOnenoteSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this User Id Onenote Section Id Page ID
func (id UserIdOnenoteSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Id Onenote Section Id Page (%s)", strings.Join(components, "\n"))
}
