package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdExtensionId{}

// GroupIdExtensionId is a struct representing the Resource ID for a Group Id Extension
type GroupIdExtensionId struct {
	GroupId     string
	ExtensionId string
}

// NewGroupIdExtensionID returns a new GroupIdExtensionId struct
func NewGroupIdExtensionID(groupId string, extensionId string) GroupIdExtensionId {
	return GroupIdExtensionId{
		GroupId:     groupId,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdExtensionID parses 'input' into a GroupIdExtensionId
func ParseGroupIdExtensionID(input string) (*GroupIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdExtensionIDInsensitively(input string) (*GroupIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdExtensionID checks that 'input' can be parsed as a Group Id Extension ID
func ValidateGroupIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Extension ID
func (id GroupIdExtensionId) ID() string {
	fmtString := "/groups/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Extension ID
func (id GroupIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Extension ID
func (id GroupIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Extension (%s)", strings.Join(components, "\n"))
}
