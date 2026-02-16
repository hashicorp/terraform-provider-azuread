package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdSubscriptionId{}

// GroupIdSiteIdListIdSubscriptionId is a struct representing the Resource ID for a Group Id Site Id List Id Subscription
type GroupIdSiteIdListIdSubscriptionId struct {
	GroupId        string
	SiteId         string
	ListId         string
	SubscriptionId string
}

// NewGroupIdSiteIdListIdSubscriptionID returns a new GroupIdSiteIdListIdSubscriptionId struct
func NewGroupIdSiteIdListIdSubscriptionID(groupId string, siteId string, listId string, subscriptionId string) GroupIdSiteIdListIdSubscriptionId {
	return GroupIdSiteIdListIdSubscriptionId{
		GroupId:        groupId,
		SiteId:         siteId,
		ListId:         listId,
		SubscriptionId: subscriptionId,
	}
}

// ParseGroupIdSiteIdListIdSubscriptionID parses 'input' into a GroupIdSiteIdListIdSubscriptionId
func ParseGroupIdSiteIdListIdSubscriptionID(input string) (*GroupIdSiteIdListIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdSubscriptionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdSubscriptionIDInsensitively(input string) (*GroupIdSiteIdListIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdSubscriptionID checks that 'input' can be parsed as a Group Id Site Id List Id Subscription ID
func ValidateGroupIdSiteIdListIdSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Subscription ID
func (id GroupIdSiteIdListIdSubscriptionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Subscription ID
func (id GroupIdSiteIdListIdSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Subscription ID
func (id GroupIdSiteIdListIdSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Subscription (%s)", strings.Join(components, "\n"))
}
