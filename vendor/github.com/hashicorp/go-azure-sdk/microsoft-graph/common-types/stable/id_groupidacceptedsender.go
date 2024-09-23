package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdAcceptedSenderId{}

// GroupIdAcceptedSenderId is a struct representing the Resource ID for a Group Id Accepted Sender
type GroupIdAcceptedSenderId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupIdAcceptedSenderID returns a new GroupIdAcceptedSenderId struct
func NewGroupIdAcceptedSenderID(groupId string, directoryObjectId string) GroupIdAcceptedSenderId {
	return GroupIdAcceptedSenderId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupIdAcceptedSenderID parses 'input' into a GroupIdAcceptedSenderId
func ParseGroupIdAcceptedSenderID(input string) (*GroupIdAcceptedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdAcceptedSenderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdAcceptedSenderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdAcceptedSenderIDInsensitively parses 'input' case-insensitively into a GroupIdAcceptedSenderId
// note: this method should only be used for API response data and not user input
func ParseGroupIdAcceptedSenderIDInsensitively(input string) (*GroupIdAcceptedSenderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdAcceptedSenderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdAcceptedSenderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdAcceptedSenderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateGroupIdAcceptedSenderID checks that 'input' can be parsed as a Group Id Accepted Sender ID
func ValidateGroupIdAcceptedSenderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdAcceptedSenderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Accepted Sender ID
func (id GroupIdAcceptedSenderId) ID() string {
	fmtString := "/groups/%s/acceptedSenders/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Accepted Sender ID
func (id GroupIdAcceptedSenderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("acceptedSenders", "acceptedSenders", "acceptedSenders"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Group Id Accepted Sender ID
func (id GroupIdAcceptedSenderId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Id Accepted Sender (%s)", strings.Join(components, "\n"))
}
