package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootListItemDocumentSetVersionId{}

// GroupIdDriveIdRootListItemDocumentSetVersionId is a struct representing the Resource ID for a Group Id Drive Id Root List Item Document Set Version
type GroupIdDriveIdRootListItemDocumentSetVersionId struct {
	GroupId              string
	DriveId              string
	DocumentSetVersionId string
}

// NewGroupIdDriveIdRootListItemDocumentSetVersionID returns a new GroupIdDriveIdRootListItemDocumentSetVersionId struct
func NewGroupIdDriveIdRootListItemDocumentSetVersionID(groupId string, driveId string, documentSetVersionId string) GroupIdDriveIdRootListItemDocumentSetVersionId {
	return GroupIdDriveIdRootListItemDocumentSetVersionId{
		GroupId:              groupId,
		DriveId:              driveId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseGroupIdDriveIdRootListItemDocumentSetVersionID parses 'input' into a GroupIdDriveIdRootListItemDocumentSetVersionId
func ParseGroupIdDriveIdRootListItemDocumentSetVersionID(input string) (*GroupIdDriveIdRootListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootListItemDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootListItemDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootListItemDocumentSetVersionIDInsensitively(input string) (*GroupIdDriveIdRootListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootListItemDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootListItemDocumentSetVersionID checks that 'input' can be parsed as a Group Id Drive Id Root List Item Document Set Version ID
func ValidateGroupIdDriveIdRootListItemDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootListItemDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root List Item Document Set Version ID
func (id GroupIdDriveIdRootListItemDocumentSetVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/listItem/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root List Item Document Set Version ID
func (id GroupIdDriveIdRootListItemDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root List Item Document Set Version ID
func (id GroupIdDriveIdRootListItemDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root List Item Document Set Version (%s)", strings.Join(components, "\n"))
}
