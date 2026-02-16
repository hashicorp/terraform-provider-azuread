package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupId{}

// GroupId is a struct representing the Resource ID for a Group
type GroupId struct {
	GroupId string
}

// NewGroupID returns a new GroupId struct
func NewGroupID(groupId string) GroupId {
	return GroupId{
		GroupId: groupId,
	}
}

// ParseGroupID parses 'input' into a GroupId
func ParseGroupID(input string) (*GroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIDInsensitively parses 'input' case-insensitively into a GroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIDInsensitively(input string) (*GroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	return nil
}

// ValidateGroupID checks that 'input' can be parsed as a Group ID
func ValidateGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group ID
func (id GroupId) ID() string {
	fmtString := "/groups/%s"
	return fmt.Sprintf(fmtString, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group ID
func (id GroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
	}
}

// String returns a human-readable description of this Group ID
func (id GroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Group (%s)", strings.Join(components, "\n"))
}
