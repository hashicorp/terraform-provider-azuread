package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemIdDocumentSetVersionId{}

// GroupIdDriveIdListItemIdDocumentSetVersionId is a struct representing the Resource ID for a Group Id Drive Id List Item Id Document Set Version
type GroupIdDriveIdListItemIdDocumentSetVersionId struct {
	GroupId              string
	DriveId              string
	ListItemId           string
	DocumentSetVersionId string
}

// NewGroupIdDriveIdListItemIdDocumentSetVersionID returns a new GroupIdDriveIdListItemIdDocumentSetVersionId struct
func NewGroupIdDriveIdListItemIdDocumentSetVersionID(groupId string, driveId string, listItemId string, documentSetVersionId string) GroupIdDriveIdListItemIdDocumentSetVersionId {
	return GroupIdDriveIdListItemIdDocumentSetVersionId{
		GroupId:              groupId,
		DriveId:              driveId,
		ListItemId:           listItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseGroupIdDriveIdListItemIdDocumentSetVersionID parses 'input' into a GroupIdDriveIdListItemIdDocumentSetVersionId
func ParseGroupIdDriveIdListItemIdDocumentSetVersionID(input string) (*GroupIdDriveIdListItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListItemIdDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListItemIdDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListItemIdDocumentSetVersionIDInsensitively(input string) (*GroupIdDriveIdListItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListItemIdDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListItemIdDocumentSetVersionID checks that 'input' can be parsed as a Group Id Drive Id List Item Id Document Set Version ID
func ValidateGroupIdDriveIdListItemIdDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListItemIdDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Item Id Document Set Version ID
func (id GroupIdDriveIdListItemIdDocumentSetVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/items/%s/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ListItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Item Id Document Set Version ID
func (id GroupIdDriveIdListItemIdDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Item Id Document Set Version ID
func (id GroupIdDriveIdListItemIdDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Item Id Document Set Version (%s)", strings.Join(components, "\n"))
}
