package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdMemberOfId{}

// GroupIdMemberOfId is a struct representing the Resource ID for a Group Id Member Of
type GroupIdMemberOfId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdMemberOfID returns a new GroupIdMemberOfId struct
func NewGroupIdMemberOfID(groupId string, directoryObjectId string) GroupIdMemberOfId {
	return GroupIdMemberOfId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdMemberOfID parses 'input' into a GroupIdMemberOfId
func ParseGroupIdMemberOfID(input string) (*GroupIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdMemberOfIDInsensitively parses 'input' case-insensitively into a GroupIdMemberOfId
// note: this method should only be used for API response data and not user input
func ParseGroupIdMemberOfIDInsensitively(input string) (*GroupIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdMemberOfID checks that 'input' can be parsed as a Group Id Member Of ID
func ValidateGroupIdMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Member Of ID
func (id GroupIdMemberOfId) ID() string {
	fmtString := "/groups/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Member Of ID
func (id GroupIdMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("memberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Member Of ID
func (id GroupIdMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Member Of (%s)", strings.Join(components, "\n"))
}
