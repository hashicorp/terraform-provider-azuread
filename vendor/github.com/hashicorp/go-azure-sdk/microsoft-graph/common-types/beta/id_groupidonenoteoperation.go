package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteOperationId{}

// GroupIdOnenoteOperationId is a struct representing the Resource ID for a Group Id Onenote Operation
type GroupIdOnenoteOperationId struct {
	GroupId            string
	OnenoteOperationId string
}

// NewGroupIdOnenoteOperationID returns a new GroupIdOnenoteOperationId struct
func NewGroupIdOnenoteOperationID(groupId string, onenoteOperationId string) GroupIdOnenoteOperationId {
	return GroupIdOnenoteOperationId{
		GroupId:            groupId,
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseGroupIdOnenoteOperationID parses 'input' into a GroupIdOnenoteOperationId
func ParseGroupIdOnenoteOperationID(input string) (*GroupIdOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteOperationIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteOperationIDInsensitively(input string) (*GroupIdOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OnenoteOperationId, ok = input.Parsed["onenoteOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteOperationID checks that 'input' can be parsed as a Group Id Onenote Operation ID
func ValidateGroupIdOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Operation ID
func (id GroupIdOnenoteOperationId) ID() string {
	fmtString := "/groups/%s/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Operation ID
func (id GroupIdOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Operation ID
func (id GroupIdOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("Group Id Onenote Operation (%s)", strings.Join(components, "\n"))
}
