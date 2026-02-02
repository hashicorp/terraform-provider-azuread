package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteOperationId{}

// GroupIdSiteIdOnenoteOperationId is a struct representing the Resource ID for a Group Id Site Id Onenote Operation
type GroupIdSiteIdOnenoteOperationId struct {
	GroupId            string
	SiteId             string
	OnenoteOperationId string
}

// NewGroupIdSiteIdOnenoteOperationID returns a new GroupIdSiteIdOnenoteOperationId struct
func NewGroupIdSiteIdOnenoteOperationID(groupId string, siteId string, onenoteOperationId string) GroupIdSiteIdOnenoteOperationId {
	return GroupIdSiteIdOnenoteOperationId{
		GroupId:            groupId,
		SiteId:             siteId,
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseGroupIdSiteIdOnenoteOperationID parses 'input' into a GroupIdSiteIdOnenoteOperationId
func ParseGroupIdSiteIdOnenoteOperationID(input string) (*GroupIdSiteIdOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteOperationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteOperationIDInsensitively(input string) (*GroupIdSiteIdOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.OnenoteOperationId, ok = input.Parsed["onenoteOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteOperationID checks that 'input' can be parsed as a Group Id Site Id Onenote Operation ID
func ValidateGroupIdSiteIdOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Operation ID
func (id GroupIdSiteIdOnenoteOperationId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Operation ID
func (id GroupIdSiteIdOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Operation ID
func (id GroupIdSiteIdOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Operation (%s)", strings.Join(components, "\n"))
}
