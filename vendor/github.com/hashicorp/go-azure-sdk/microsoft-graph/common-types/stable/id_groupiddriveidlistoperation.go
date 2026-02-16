package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListOperationId{}

// GroupIdDriveIdListOperationId is a struct representing the Resource ID for a Group Id Drive Id List Operation
type GroupIdDriveIdListOperationId struct {
	GroupId                    string
	DriveId                    string
	RichLongRunningOperationId string
}

// NewGroupIdDriveIdListOperationID returns a new GroupIdDriveIdListOperationId struct
func NewGroupIdDriveIdListOperationID(groupId string, driveId string, richLongRunningOperationId string) GroupIdDriveIdListOperationId {
	return GroupIdDriveIdListOperationId{
		GroupId:                    groupId,
		DriveId:                    driveId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseGroupIdDriveIdListOperationID parses 'input' into a GroupIdDriveIdListOperationId
func ParseGroupIdDriveIdListOperationID(input string) (*GroupIdDriveIdListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListOperationIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListOperationIDInsensitively(input string) (*GroupIdDriveIdListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.RichLongRunningOperationId, ok = input.Parsed["richLongRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListOperationID checks that 'input' can be parsed as a Group Id Drive Id List Operation ID
func ValidateGroupIdDriveIdListOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Operation ID
func (id GroupIdDriveIdListOperationId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Operation ID
func (id GroupIdDriveIdListOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Operation ID
func (id GroupIdDriveIdListOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("Group Id Drive Id List Operation (%s)", strings.Join(components, "\n"))
}
