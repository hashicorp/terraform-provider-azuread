package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemIdActivityId{}

// GroupIdDriveIdListItemIdActivityId is a struct representing the Resource ID for a Group Id Drive Id List Item Id Activity
type GroupIdDriveIdListItemIdActivityId struct {
	GroupId           string
	DriveId           string
	ListItemId        string
	ItemActivityOLDId string
}

// NewGroupIdDriveIdListItemIdActivityID returns a new GroupIdDriveIdListItemIdActivityId struct
func NewGroupIdDriveIdListItemIdActivityID(groupId string, driveId string, listItemId string, itemActivityOLDId string) GroupIdDriveIdListItemIdActivityId {
	return GroupIdDriveIdListItemIdActivityId{
		GroupId:           groupId,
		DriveId:           driveId,
		ListItemId:        listItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdDriveIdListItemIdActivityID parses 'input' into a GroupIdDriveIdListItemIdActivityId
func ParseGroupIdDriveIdListItemIdActivityID(input string) (*GroupIdDriveIdListItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListItemIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListItemIdActivityIDInsensitively(input string) (*GroupIdDriveIdListItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListItemIdActivityID checks that 'input' can be parsed as a Group Id Drive Id List Item Id Activity ID
func ValidateGroupIdDriveIdListItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Item Id Activity ID
func (id GroupIdDriveIdListItemIdActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ListItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Item Id Activity ID
func (id GroupIdDriveIdListItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Item Id Activity ID
func (id GroupIdDriveIdListItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Drive Id List Item Id Activity (%s)", strings.Join(components, "\n"))
}
