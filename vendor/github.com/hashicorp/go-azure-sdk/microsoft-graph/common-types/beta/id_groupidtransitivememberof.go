package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTransitiveMemberOfId{}

// GroupIdTransitiveMemberOfId is a struct representing the Resource ID for a Group Id Transitive Member Of
type GroupIdTransitiveMemberOfId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdTransitiveMemberOfID returns a new GroupIdTransitiveMemberOfId struct
func NewGroupIdTransitiveMemberOfID(groupId string, directoryObjectId string) GroupIdTransitiveMemberOfId {
	return GroupIdTransitiveMemberOfId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdTransitiveMemberOfID parses 'input' into a GroupIdTransitiveMemberOfId
func ParseGroupIdTransitiveMemberOfID(input string) (*GroupIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a GroupIdTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTransitiveMemberOfIDInsensitively(input string) (*GroupIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTransitiveMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdTransitiveMemberOfID checks that 'input' can be parsed as a Group Id Transitive Member Of ID
func ValidateGroupIdTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Transitive Member Of ID
func (id GroupIdTransitiveMemberOfId) ID() string {
	fmtString := "/groups/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Transitive Member Of ID
func (id GroupIdTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("transitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Transitive Member Of ID
func (id GroupIdTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Transitive Member Of (%s)", strings.Join(components, "\n"))
}
