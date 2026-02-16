package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdOperationId{}

// GroupIdSiteIdListIdOperationId is a struct representing the Resource ID for a Group Id Site Id List Id Operation
type GroupIdSiteIdListIdOperationId struct {
	GroupId                    string
	SiteId                     string
	ListId                     string
	RichLongRunningOperationId string
}

// NewGroupIdSiteIdListIdOperationID returns a new GroupIdSiteIdListIdOperationId struct
func NewGroupIdSiteIdListIdOperationID(groupId string, siteId string, listId string, richLongRunningOperationId string) GroupIdSiteIdListIdOperationId {
	return GroupIdSiteIdListIdOperationId{
		GroupId:                    groupId,
		SiteId:                     siteId,
		ListId:                     listId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseGroupIdSiteIdListIdOperationID parses 'input' into a GroupIdSiteIdListIdOperationId
func ParseGroupIdSiteIdListIdOperationID(input string) (*GroupIdSiteIdListIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdOperationIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdOperationIDInsensitively(input string) (*GroupIdSiteIdListIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ListId, ok = input.Parsed["listId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listId", input)
	}

	if id.RichLongRunningOperationId, ok = input.Parsed["richLongRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdOperationID checks that 'input' can be parsed as a Group Id Site Id List Id Operation ID
func ValidateGroupIdSiteIdListIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Operation ID
func (id GroupIdSiteIdListIdOperationId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Operation ID
func (id GroupIdSiteIdListIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Operation ID
func (id GroupIdSiteIdListIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Operation (%s)", strings.Join(components, "\n"))
}
