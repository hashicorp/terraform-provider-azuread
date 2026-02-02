package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenotePageId{}

// GroupIdOnenotePageId is a struct representing the Resource ID for a Group Id Onenote Page
type GroupIdOnenotePageId struct {
	GroupId       string
	OnenotePageId string
}

// NewGroupIdOnenotePageID returns a new GroupIdOnenotePageId struct
func NewGroupIdOnenotePageID(groupId string, onenotePageId string) GroupIdOnenotePageId {
	return GroupIdOnenotePageId{
		GroupId:       groupId,
		OnenotePageId: onenotePageId,
	}
}

// ParseGroupIdOnenotePageID parses 'input' into a GroupIdOnenotePageId
func ParseGroupIdOnenotePageID(input string) (*GroupIdOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenotePageIDInsensitively parses 'input' case-insensitively into a GroupIdOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenotePageIDInsensitively(input string) (*GroupIdOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenotePageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdOnenotePageID checks that 'input' can be parsed as a Group Id Onenote Page ID
func ValidateGroupIdOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Page ID
func (id GroupIdOnenotePageId) ID() string {
	fmtString := "/groups/%s/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Page ID
func (id GroupIdOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Page ID
func (id GroupIdOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Onenote Page (%s)", strings.Join(components, "\n"))
}
