package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdMemberId{}

// GroupIdMemberId is a struct representing the Resource ID for a Group Id Member
type GroupIdMemberId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdMemberID returns a new GroupIdMemberId struct
func NewGroupIdMemberID(groupId string, directoryObjectId string) GroupIdMemberId {
	return GroupIdMemberId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdMemberID parses 'input' into a GroupIdMemberId
func ParseGroupIdMemberID(input string) (*GroupIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdMemberIDInsensitively parses 'input' case-insensitively into a GroupIdMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdMemberIDInsensitively(input string) (*GroupIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdMemberID checks that 'input' can be parsed as a Group Id Member ID
func ValidateGroupIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Member ID
func (id GroupIdMemberId) ID() string {
	fmtString := "/groups/%s/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Member ID
func (id GroupIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Member ID
func (id GroupIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Member (%s)", strings.Join(components, "\n"))
}
