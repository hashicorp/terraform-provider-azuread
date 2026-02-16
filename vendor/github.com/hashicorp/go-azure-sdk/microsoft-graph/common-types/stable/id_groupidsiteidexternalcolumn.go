package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdExternalColumnId{}

// GroupIdSiteIdExternalColumnId is a struct representing the Resource ID for a Group Id Site Id External Column
type GroupIdSiteIdExternalColumnId struct {
	GroupId            string
	SiteId             string
	ColumnDefinitionId string
}

// NewGroupIdSiteIdExternalColumnID returns a new GroupIdSiteIdExternalColumnId struct
func NewGroupIdSiteIdExternalColumnID(groupId string, siteId string, columnDefinitionId string) GroupIdSiteIdExternalColumnId {
	return GroupIdSiteIdExternalColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdSiteIdExternalColumnID parses 'input' into a GroupIdSiteIdExternalColumnId
func ParseGroupIdSiteIdExternalColumnID(input string) (*GroupIdSiteIdExternalColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdExternalColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdExternalColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdExternalColumnIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdExternalColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdExternalColumnIDInsensitively(input string) (*GroupIdSiteIdExternalColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdExternalColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdExternalColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdExternalColumnId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdSiteIdExternalColumnID checks that 'input' can be parsed as a Group Id Site Id External Column ID
func ValidateGroupIdSiteIdExternalColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdExternalColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id External Column ID
func (id GroupIdSiteIdExternalColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/externalColumns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id External Column ID
func (id GroupIdSiteIdExternalColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("externalColumns", "externalColumns", "externalColumns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id External Column ID
func (id GroupIdSiteIdExternalColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Site Id External Column (%s)", strings.Join(components, "\n"))
}
