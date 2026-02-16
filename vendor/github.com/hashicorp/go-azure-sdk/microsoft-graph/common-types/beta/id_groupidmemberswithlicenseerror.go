package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdMembersWithLicenseErrorId{}

// GroupIdMembersWithLicenseErrorId is a struct representing the Resource ID for a Group Id Members With License Error
type GroupIdMembersWithLicenseErrorId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdMembersWithLicenseErrorID returns a new GroupIdMembersWithLicenseErrorId struct
func NewGroupIdMembersWithLicenseErrorID(groupId string, directoryObjectId string) GroupIdMembersWithLicenseErrorId {
	return GroupIdMembersWithLicenseErrorId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdMembersWithLicenseErrorID parses 'input' into a GroupIdMembersWithLicenseErrorId
func ParseGroupIdMembersWithLicenseErrorID(input string) (*GroupIdMembersWithLicenseErrorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdMembersWithLicenseErrorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdMembersWithLicenseErrorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdMembersWithLicenseErrorIDInsensitively parses 'input' case-insensitively into a GroupIdMembersWithLicenseErrorId
// note: this method should only be used for API response data and not user input
func ParseGroupIdMembersWithLicenseErrorIDInsensitively(input string) (*GroupIdMembersWithLicenseErrorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdMembersWithLicenseErrorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdMembersWithLicenseErrorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdMembersWithLicenseErrorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdMembersWithLicenseErrorID checks that 'input' can be parsed as a Group Id Members With License Error ID
func ValidateGroupIdMembersWithLicenseErrorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdMembersWithLicenseErrorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Members With License Error ID
func (id GroupIdMembersWithLicenseErrorId) ID() string {
	fmtString := "/groups/%s/membersWithLicenseErrors/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Members With License Error ID
func (id GroupIdMembersWithLicenseErrorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("membersWithLicenseErrors", "membersWithLicenseErrors", "membersWithLicenseErrors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Members With License Error ID
func (id GroupIdMembersWithLicenseErrorId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Members With License Error (%s)", strings.Join(components, "\n"))
}
