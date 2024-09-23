package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootListItemVersionId{}

// MeDriveIdRootListItemVersionId is a struct representing the Resource ID for a Me Drive Id Root List Item Version
type MeDriveIdRootListItemVersionId struct {
	DriveId           string
	ListItemVersionId string
}

// NewMeDriveIdRootListItemVersionID returns a new MeDriveIdRootListItemVersionId struct
func NewMeDriveIdRootListItemVersionID(driveId string, listItemVersionId string) MeDriveIdRootListItemVersionId {
	return MeDriveIdRootListItemVersionId{
		DriveId:           driveId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseMeDriveIdRootListItemVersionID parses 'input' into a MeDriveIdRootListItemVersionId
func ParseMeDriveIdRootListItemVersionID(input string) (*MeDriveIdRootListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootListItemVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootListItemVersionIDInsensitively(input string) (*MeDriveIdRootListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootListItemVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootListItemVersionID checks that 'input' can be parsed as a Me Drive Id Root List Item Version ID
func ValidateMeDriveIdRootListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root List Item Version ID
func (id MeDriveIdRootListItemVersionId) ID() string {
	fmtString := "/me/drives/%s/root/listItem/versions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root List Item Version ID
func (id MeDriveIdRootListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root List Item Version ID
func (id MeDriveIdRootListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Me Drive Id Root List Item Version (%s)", strings.Join(components, "\n"))
}
