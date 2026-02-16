package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteSectionId{}

// GroupIdOnenoteSectionId is a struct representing the Resource ID for a Group Id Onenote Section
type GroupIdOnenoteSectionId struct {
	GroupId          string
	OnenoteSectionId string
}

// NewGroupIdOnenoteSectionID returns a new GroupIdOnenoteSectionId struct
func NewGroupIdOnenoteSectionID(groupId string, onenoteSectionId string) GroupIdOnenoteSectionId {
	return GroupIdOnenoteSectionId{
		GroupId:          groupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdOnenoteSectionID parses 'input' into a GroupIdOnenoteSectionId
func ParseGroupIdOnenoteSectionID(input string) (*GroupIdOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteSectionIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteSectionIDInsensitively(input string) (*GroupIdOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteSectionID checks that 'input' can be parsed as a Group Id Onenote Section ID
func ValidateGroupIdOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Section ID
func (id GroupIdOnenoteSectionId) ID() string {
	fmtString := "/groups/%s/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Section ID
func (id GroupIdOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Section ID
func (id GroupIdOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Onenote Section (%s)", strings.Join(components, "\n"))
}
