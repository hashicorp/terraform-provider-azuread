package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdAnalyticsItemActivityStatId{}

// GroupIdSiteIdAnalyticsItemActivityStatId is a struct representing the Resource ID for a Group Id Site Id Analytics Item Activity Stat
type GroupIdSiteIdAnalyticsItemActivityStatId struct {
	GroupId            string
	SiteId             string
	ItemActivityStatId string
}

// NewGroupIdSiteIdAnalyticsItemActivityStatID returns a new GroupIdSiteIdAnalyticsItemActivityStatId struct
func NewGroupIdSiteIdAnalyticsItemActivityStatID(groupId string, siteId string, itemActivityStatId string) GroupIdSiteIdAnalyticsItemActivityStatId {
	return GroupIdSiteIdAnalyticsItemActivityStatId{
		GroupId:            groupId,
		SiteId:             siteId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseGroupIdSiteIdAnalyticsItemActivityStatID parses 'input' into a GroupIdSiteIdAnalyticsItemActivityStatId
func ParseGroupIdSiteIdAnalyticsItemActivityStatID(input string) (*GroupIdSiteIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdAnalyticsItemActivityStatIDInsensitively(input string) (*GroupIdSiteIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdAnalyticsItemActivityStatID checks that 'input' can be parsed as a Group Id Site Id Analytics Item Activity Stat ID
func ValidateGroupIdSiteIdAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Analytics Item Activity Stat ID
func (id GroupIdSiteIdAnalyticsItemActivityStatId) ID() string {
	fmtString := "/groups/%s/sites/%s/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Analytics Item Activity Stat ID
func (id GroupIdSiteIdAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Analytics Item Activity Stat ID
func (id GroupIdSiteIdAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Group Id Site Id Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
