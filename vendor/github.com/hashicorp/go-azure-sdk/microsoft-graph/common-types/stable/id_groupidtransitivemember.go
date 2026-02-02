package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTransitiveMemberId{}

// GroupIdTransitiveMemberId is a struct representing the Resource ID for a Group Id Transitive Member
type GroupIdTransitiveMemberId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdTransitiveMemberID returns a new GroupIdTransitiveMemberId struct
func NewGroupIdTransitiveMemberID(groupId string, directoryObjectId string) GroupIdTransitiveMemberId {
	return GroupIdTransitiveMemberId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdTransitiveMemberID parses 'input' into a GroupIdTransitiveMemberId
func ParseGroupIdTransitiveMemberID(input string) (*GroupIdTransitiveMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTransitiveMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTransitiveMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTransitiveMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTransitiveMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTransitiveMemberIDInsensitively(input string) (*GroupIdTransitiveMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTransitiveMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTransitiveMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTransitiveMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdTransitiveMemberID checks that 'input' can be parsed as a Group Id Transitive Member ID
func ValidateGroupIdTransitiveMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTransitiveMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Transitive Member ID
func (id GroupIdTransitiveMemberId) ID() string {
	fmtString := "/groups/%s/transitiveMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Transitive Member ID
func (id GroupIdTransitiveMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("transitiveMembers", "transitiveMembers", "transitiveMembers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Transitive Member ID
func (id GroupIdTransitiveMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Transitive Member (%s)", strings.Join(components, "\n"))
}
