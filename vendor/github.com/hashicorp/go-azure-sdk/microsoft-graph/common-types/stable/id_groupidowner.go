package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOwnerId{}

// GroupIdOwnerId is a struct representing the Resource ID for a Group Id Owner
type GroupIdOwnerId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdOwnerID returns a new GroupIdOwnerId struct
func NewGroupIdOwnerID(groupId string, directoryObjectId string) GroupIdOwnerId {
	return GroupIdOwnerId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdOwnerID parses 'input' into a GroupIdOwnerId
func ParseGroupIdOwnerID(input string) (*GroupIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOwnerIDInsensitively parses 'input' case-insensitively into a GroupIdOwnerId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOwnerIDInsensitively(input string) (*GroupIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdOwnerID checks that 'input' can be parsed as a Group Id Owner ID
func ValidateGroupIdOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Owner ID
func (id GroupIdOwnerId) ID() string {
	fmtString := "/groups/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Owner ID
func (id GroupIdOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("owners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Owner ID
func (id GroupIdOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Owner (%s)", strings.Join(components, "\n"))
}
