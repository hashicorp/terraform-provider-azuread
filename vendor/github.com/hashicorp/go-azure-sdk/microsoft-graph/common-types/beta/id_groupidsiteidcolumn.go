package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdColumnId{}

// GroupIdSiteIdColumnId is a struct representing the Resource ID for a Group Id Site Id Column
type GroupIdSiteIdColumnId struct {
	GroupId            string
	SiteId             string
	ColumnDefinitionId string
}

// NewGroupIdSiteIdColumnID returns a new GroupIdSiteIdColumnId struct
func NewGroupIdSiteIdColumnID(groupId string, siteId string, columnDefinitionId string) GroupIdSiteIdColumnId {
	return GroupIdSiteIdColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdSiteIdColumnID parses 'input' into a GroupIdSiteIdColumnId
func ParseGroupIdSiteIdColumnID(input string) (*GroupIdSiteIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdColumnIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdColumnIDInsensitively(input string) (*GroupIdSiteIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdColumnID checks that 'input' can be parsed as a Group Id Site Id Column ID
func ValidateGroupIdSiteIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Column ID
func (id GroupIdSiteIdColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Column ID
func (id GroupIdSiteIdColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Column ID
func (id GroupIdSiteIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Site Id Column (%s)", strings.Join(components, "\n"))
}
