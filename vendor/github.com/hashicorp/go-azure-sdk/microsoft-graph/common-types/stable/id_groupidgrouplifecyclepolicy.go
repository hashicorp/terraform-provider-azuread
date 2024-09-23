package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdGroupLifecyclePolicyId{}

// GroupIdGroupLifecyclePolicyId is a struct representing the Resource ID for a Group Id Group Lifecycle Policy
type GroupIdGroupLifecyclePolicyId struct {
	GroupId                string
	GroupLifecyclePolicyId string
}

// NewGroupIdGroupLifecyclePolicyID returns a new GroupIdGroupLifecyclePolicyId struct
func NewGroupIdGroupLifecyclePolicyID(groupId string, groupLifecyclePolicyId string) GroupIdGroupLifecyclePolicyId {
	return GroupIdGroupLifecyclePolicyId{
		GroupId:                groupId,
		GroupLifecyclePolicyId: groupLifecyclePolicyId,
	}
}

// ParseGroupIdGroupLifecyclePolicyID parses 'input' into a GroupIdGroupLifecyclePolicyId
func ParseGroupIdGroupLifecyclePolicyID(input string) (*GroupIdGroupLifecyclePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdGroupLifecyclePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdGroupLifecyclePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdGroupLifecyclePolicyIDInsensitively parses 'input' case-insensitively into a GroupIdGroupLifecyclePolicyId
// note: this method should only be used for API response data and not user input
func ParseGroupIdGroupLifecyclePolicyIDInsensitively(input string) (*GroupIdGroupLifecyclePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdGroupLifecyclePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdGroupLifecyclePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdGroupLifecyclePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.GroupLifecyclePolicyId, ok = input.Parsed["groupLifecyclePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupLifecyclePolicyId", input)
	}

	return nil
}

// ValidateGroupIdGroupLifecyclePolicyID checks that 'input' can be parsed as a Group Id Group Lifecycle Policy ID
func ValidateGroupIdGroupLifecyclePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdGroupLifecyclePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Group Lifecycle Policy ID
func (id GroupIdGroupLifecyclePolicyId) ID() string {
	fmtString := "/groups/%s/groupLifecyclePolicies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.GroupLifecyclePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Group Lifecycle Policy ID
func (id GroupIdGroupLifecyclePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("groupLifecyclePolicies", "groupLifecyclePolicies", "groupLifecyclePolicies"),
		resourceids.UserSpecifiedSegment("groupLifecyclePolicyId", "groupLifecyclePolicyId"),
	}
}

// String returns a human-readable description of this Group Id Group Lifecycle Policy ID
func (id GroupIdGroupLifecyclePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Group Lifecycle Policy: %q", id.GroupLifecyclePolicyId),
	}
	return fmt.Sprintf("Group Id Group Lifecycle Policy (%s)", strings.Join(components, "\n"))
}
