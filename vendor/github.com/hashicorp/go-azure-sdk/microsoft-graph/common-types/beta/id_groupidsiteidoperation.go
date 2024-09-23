package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOperationId{}

// GroupIdSiteIdOperationId is a struct representing the Resource ID for a Group Id Site Id Operation
type GroupIdSiteIdOperationId struct {
	GroupId                    string
	SiteId                     string
	RichLongRunningOperationId string
}

// NewGroupIdSiteIdOperationID returns a new GroupIdSiteIdOperationId struct
func NewGroupIdSiteIdOperationID(groupId string, siteId string, richLongRunningOperationId string) GroupIdSiteIdOperationId {
	return GroupIdSiteIdOperationId{
		GroupId:                    groupId,
		SiteId:                     siteId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseGroupIdSiteIdOperationID parses 'input' into a GroupIdSiteIdOperationId
func ParseGroupIdSiteIdOperationID(input string) (*GroupIdSiteIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOperationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOperationIDInsensitively(input string) (*GroupIdSiteIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.RichLongRunningOperationId, ok = input.Parsed["richLongRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOperationID checks that 'input' can be parsed as a Group Id Site Id Operation ID
func ValidateGroupIdSiteIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Operation ID
func (id GroupIdSiteIdOperationId) ID() string {
	fmtString := "/groups/%s/sites/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Operation ID
func (id GroupIdSiteIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Operation ID
func (id GroupIdSiteIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("Group Id Site Id Operation (%s)", strings.Join(components, "\n"))
}
