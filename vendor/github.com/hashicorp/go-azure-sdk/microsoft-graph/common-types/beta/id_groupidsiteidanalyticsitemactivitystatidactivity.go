package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{}

// GroupIdSiteIdAnalyticsItemActivityStatIdActivityId is a struct representing the Resource ID for a Group Id Site Id Analytics Item Activity Stat Id Activity
type GroupIdSiteIdAnalyticsItemActivityStatIdActivityId struct {
	GroupId            string
	SiteId             string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewGroupIdSiteIdAnalyticsItemActivityStatIdActivityID returns a new GroupIdSiteIdAnalyticsItemActivityStatIdActivityId struct
func NewGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(groupId string, siteId string, itemActivityStatId string, itemActivityId string) GroupIdSiteIdAnalyticsItemActivityStatIdActivityId {
	return GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{
		GroupId:            groupId,
		SiteId:             siteId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityID parses 'input' into a GroupIdSiteIdAnalyticsItemActivityStatIdActivityId
func ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(input string) (*GroupIdSiteIdAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdAnalyticsItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityIDInsensitively(input string) (*GroupIdSiteIdAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdAnalyticsItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityId, ok = input.Parsed["itemActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdAnalyticsItemActivityStatIdActivityID checks that 'input' can be parsed as a Group Id Site Id Analytics Item Activity Stat Id Activity ID
func ValidateGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Analytics Item Activity Stat Id Activity ID
func (id GroupIdSiteIdAnalyticsItemActivityStatIdActivityId) ID() string {
	fmtString := "/groups/%s/sites/%s/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Analytics Item Activity Stat Id Activity ID
func (id GroupIdSiteIdAnalyticsItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Analytics Item Activity Stat Id Activity ID
func (id GroupIdSiteIdAnalyticsItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Group Id Site Id Analytics Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
