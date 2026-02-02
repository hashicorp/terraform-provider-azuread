package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteSectionGroupId{}

// GroupIdOnenoteSectionGroupId is a struct representing the Resource ID for a Group Id Onenote Section Group
type GroupIdOnenoteSectionGroupId struct {
	GroupId        string
	SectionGroupId string
}

// NewGroupIdOnenoteSectionGroupID returns a new GroupIdOnenoteSectionGroupId struct
func NewGroupIdOnenoteSectionGroupID(groupId string, sectionGroupId string) GroupIdOnenoteSectionGroupId {
	return GroupIdOnenoteSectionGroupId{
		GroupId:        groupId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupIdOnenoteSectionGroupID parses 'input' into a GroupIdOnenoteSectionGroupId
func ParseGroupIdOnenoteSectionGroupID(input string) (*GroupIdOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteSectionGroupIDInsensitively(input string) (*GroupIdOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteSectionGroupID checks that 'input' can be parsed as a Group Id Onenote Section Group ID
func ValidateGroupIdOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Section Group ID
func (id GroupIdOnenoteSectionGroupId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Section Group ID
func (id GroupIdOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Section Group ID
func (id GroupIdOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Id Onenote Section Group (%s)", strings.Join(components, "\n"))
}
