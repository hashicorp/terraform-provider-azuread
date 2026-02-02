package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdListItemDocumentSetVersionId{}

// GroupIdDriveIdItemIdListItemDocumentSetVersionId is a struct representing the Resource ID for a Group Id Drive Id Item Id List Item Document Set Version
type GroupIdDriveIdItemIdListItemDocumentSetVersionId struct {
	GroupId              string
	DriveId              string
	DriveItemId          string
	DocumentSetVersionId string
}

// NewGroupIdDriveIdItemIdListItemDocumentSetVersionID returns a new GroupIdDriveIdItemIdListItemDocumentSetVersionId struct
func NewGroupIdDriveIdItemIdListItemDocumentSetVersionID(groupId string, driveId string, driveItemId string, documentSetVersionId string) GroupIdDriveIdItemIdListItemDocumentSetVersionId {
	return GroupIdDriveIdItemIdListItemDocumentSetVersionId{
		GroupId:              groupId,
		DriveId:              driveId,
		DriveItemId:          driveItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseGroupIdDriveIdItemIdListItemDocumentSetVersionID parses 'input' into a GroupIdDriveIdItemIdListItemDocumentSetVersionId
func ParseGroupIdDriveIdItemIdListItemDocumentSetVersionID(input string) (*GroupIdDriveIdItemIdListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdListItemDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdListItemDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdListItemDocumentSetVersionIDInsensitively(input string) (*GroupIdDriveIdItemIdListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdListItemDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdListItemDocumentSetVersionID checks that 'input' can be parsed as a Group Id Drive Id Item Id List Item Document Set Version ID
func ValidateGroupIdDriveIdItemIdListItemDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdListItemDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id List Item Document Set Version ID
func (id GroupIdDriveIdItemIdListItemDocumentSetVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/listItem/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id List Item Document Set Version ID
func (id GroupIdDriveIdItemIdListItemDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id List Item Document Set Version ID
func (id GroupIdDriveIdItemIdListItemDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id List Item Document Set Version (%s)", strings.Join(components, "\n"))
}
