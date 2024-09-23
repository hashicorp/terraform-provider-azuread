package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPermissionGrantId{}

// GroupIdPermissionGrantId is a struct representing the Resource ID for a Group Id Permission Grant
type GroupIdPermissionGrantId struct {
	GroupId                           string
	ResourceSpecificPermissionGrantId string
}

// NewGroupIdPermissionGrantID returns a new GroupIdPermissionGrantId struct
func NewGroupIdPermissionGrantID(groupId string, resourceSpecificPermissionGrantId string) GroupIdPermissionGrantId {
	return GroupIdPermissionGrantId{
		GroupId:                           groupId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseGroupIdPermissionGrantID parses 'input' into a GroupIdPermissionGrantId
func ParseGroupIdPermissionGrantID(input string) (*GroupIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdPermissionGrantIDInsensitively parses 'input' case-insensitively into a GroupIdPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseGroupIdPermissionGrantIDInsensitively(input string) (*GroupIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateGroupIdPermissionGrantID checks that 'input' can be parsed as a Group Id Permission Grant ID
func ValidateGroupIdPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Permission Grant ID
func (id GroupIdPermissionGrantId) ID() string {
	fmtString := "/groups/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Permission Grant ID
func (id GroupIdPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this Group Id Permission Grant ID
func (id GroupIdPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Group Id Permission Grant (%s)", strings.Join(components, "\n"))
}
