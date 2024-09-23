package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteSectionIdPageId{}

// GroupIdOnenoteSectionIdPageId is a struct representing the Resource ID for a Group Id Onenote Section Id Page
type GroupIdOnenoteSectionIdPageId struct {
	GroupId          string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupIdOnenoteSectionIdPageID returns a new GroupIdOnenoteSectionIdPageId struct
func NewGroupIdOnenoteSectionIdPageID(groupId string, onenoteSectionId string, onenotePageId string) GroupIdOnenoteSectionIdPageId {
	return GroupIdOnenoteSectionIdPageId{
		GroupId:          groupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupIdOnenoteSectionIdPageID parses 'input' into a GroupIdOnenoteSectionIdPageId
func ParseGroupIdOnenoteSectionIdPageID(input string) (*GroupIdOnenoteSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteSectionIdPageIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteSectionIdPageIDInsensitively(input string) (*GroupIdOnenoteSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteSectionIdPageID checks that 'input' can be parsed as a Group Id Onenote Section Id Page ID
func ValidateGroupIdOnenoteSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Section Id Page ID
func (id GroupIdOnenoteSectionIdPageId) ID() string {
	fmtString := "/groups/%s/onenote/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Section Id Page ID
func (id GroupIdOnenoteSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Section Id Page ID
func (id GroupIdOnenoteSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Onenote Section Id Page (%s)", strings.Join(components, "\n"))
}
