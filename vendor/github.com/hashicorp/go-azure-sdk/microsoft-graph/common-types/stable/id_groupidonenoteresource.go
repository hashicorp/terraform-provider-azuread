package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteResourceId{}

// GroupIdOnenoteResourceId is a struct representing the Resource ID for a Group Id Onenote Resource
type GroupIdOnenoteResourceId struct {
	GroupId           string
	OnenoteResourceId string
}

// NewGroupIdOnenoteResourceID returns a new GroupIdOnenoteResourceId struct
func NewGroupIdOnenoteResourceID(groupId string, onenoteResourceId string) GroupIdOnenoteResourceId {
	return GroupIdOnenoteResourceId{
		GroupId:           groupId,
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseGroupIdOnenoteResourceID parses 'input' into a GroupIdOnenoteResourceId
func ParseGroupIdOnenoteResourceID(input string) (*GroupIdOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteResourceIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteResourceIDInsensitively(input string) (*GroupIdOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OnenoteResourceId, ok = input.Parsed["onenoteResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteResourceID checks that 'input' can be parsed as a Group Id Onenote Resource ID
func ValidateGroupIdOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Resource ID
func (id GroupIdOnenoteResourceId) ID() string {
	fmtString := "/groups/%s/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Resource ID
func (id GroupIdOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Resource ID
func (id GroupIdOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("Group Id Onenote Resource (%s)", strings.Join(components, "\n"))
}
