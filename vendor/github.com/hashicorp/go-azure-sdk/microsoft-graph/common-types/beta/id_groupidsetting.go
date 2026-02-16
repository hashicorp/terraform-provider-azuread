package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSettingId{}

// GroupIdSettingId is a struct representing the Resource ID for a Group Id Setting
type GroupIdSettingId struct {
	GroupId            string
	DirectorySettingId string
}

// NewGroupIdSettingID returns a new GroupIdSettingId struct
func NewGroupIdSettingID(groupId string, directorySettingId string) GroupIdSettingId {
	return GroupIdSettingId{
		GroupId:            groupId,
		DirectorySettingId: directorySettingId,
	}
}

// ParseGroupIdSettingID parses 'input' into a GroupIdSettingId
func ParseGroupIdSettingID(input string) (*GroupIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSettingIDInsensitively parses 'input' case-insensitively into a GroupIdSettingId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSettingIDInsensitively(input string) (*GroupIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DirectorySettingId, ok = input.Parsed["directorySettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directorySettingId", input)
	}

	return nil
}

// ValidateGroupIdSettingID checks that 'input' can be parsed as a Group Id Setting ID
func ValidateGroupIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Setting ID
func (id GroupIdSettingId) ID() string {
	fmtString := "/groups/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectorySettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Setting ID
func (id GroupIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("directorySettingId", "directorySettingId"),
	}
}

// String returns a human-readable description of this Group Id Setting ID
func (id GroupIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Setting: %q", id.DirectorySettingId),
	}
	return fmt.Sprintf("Group Id Setting (%s)", strings.Join(components, "\n"))
}
