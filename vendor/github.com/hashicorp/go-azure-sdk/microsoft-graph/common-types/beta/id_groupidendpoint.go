package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEndpointId{}

// GroupIdEndpointId is a struct representing the Resource ID for a Group Id Endpoint
type GroupIdEndpointId struct {
	GroupId    string
	EndpointId string
}

// NewGroupIdEndpointID returns a new GroupIdEndpointId struct
func NewGroupIdEndpointID(groupId string, endpointId string) GroupIdEndpointId {
	return GroupIdEndpointId{
		GroupId:    groupId,
		EndpointId: endpointId,
	}
}

// ParseGroupIdEndpointID parses 'input' into a GroupIdEndpointId
func ParseGroupIdEndpointID(input string) (*GroupIdEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEndpointIDInsensitively parses 'input' case-insensitively into a GroupIdEndpointId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEndpointIDInsensitively(input string) (*GroupIdEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEndpointId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EndpointId, ok = input.Parsed["endpointId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "endpointId", input)
	}

	return nil
}

// ValidateGroupIdEndpointID checks that 'input' can be parsed as a Group Id Endpoint ID
func ValidateGroupIdEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Endpoint ID
func (id GroupIdEndpointId) ID() string {
	fmtString := "/groups/%s/endpoints/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EndpointId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Endpoint ID
func (id GroupIdEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("endpoints", "endpoints", "endpoints"),
		resourceids.UserSpecifiedSegment("endpointId", "endpointId"),
	}
}

// String returns a human-readable description of this Group Id Endpoint ID
func (id GroupIdEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Endpoint: %q", id.EndpointId),
	}
	return fmt.Sprintf("Group Id Endpoint (%s)", strings.Join(components, "\n"))
}
