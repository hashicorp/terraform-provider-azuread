package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdRejectedSenderId{}

// GroupIdRejectedSenderId is a struct representing the Resource ID for a Group Id Rejected Sender
type GroupIdRejectedSenderId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdRejectedSenderID returns a new GroupIdRejectedSenderId struct
func NewGroupIdRejectedSenderID(groupId string, directoryObjectId string) GroupIdRejectedSenderId {
	return GroupIdRejectedSenderId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdRejectedSenderID parses 'input' into a GroupIdRejectedSenderId
func ParseGroupIdRejectedSenderID(input string) (*GroupIdRejectedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdRejectedSenderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdRejectedSenderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdRejectedSenderIDInsensitively parses 'input' case-insensitively into a GroupIdRejectedSenderId
// note: this method should only be used for API response data and not user input
func ParseGroupIdRejectedSenderIDInsensitively(input string) (*GroupIdRejectedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdRejectedSenderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdRejectedSenderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdRejectedSenderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdRejectedSenderID checks that 'input' can be parsed as a Group Id Rejected Sender ID
func ValidateGroupIdRejectedSenderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdRejectedSenderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Rejected Sender ID
func (id GroupIdRejectedSenderId) ID() string {
	fmtString := "/groups/%s/rejectedSenders/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Rejected Sender ID
func (id GroupIdRejectedSenderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("rejectedSenders", "rejectedSenders", "rejectedSenders"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Rejected Sender ID
func (id GroupIdRejectedSenderId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Rejected Sender (%s)", strings.Join(components, "\n"))
}
