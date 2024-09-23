package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdDocumentSetVersionId{}

// MeDriveIdListItemIdDocumentSetVersionId is a struct representing the Resource ID for a Me Drive Id List Item Id Document Set Version
type MeDriveIdListItemIdDocumentSetVersionId struct {
	DriveId              string
	ListItemId           string
	DocumentSetVersionId string
}

// NewMeDriveIdListItemIdDocumentSetVersionID returns a new MeDriveIdListItemIdDocumentSetVersionId struct
func NewMeDriveIdListItemIdDocumentSetVersionID(driveId string, listItemId string, documentSetVersionId string) MeDriveIdListItemIdDocumentSetVersionId {
	return MeDriveIdListItemIdDocumentSetVersionId{
		DriveId:              driveId,
		ListItemId:           listItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseMeDriveIdListItemIdDocumentSetVersionID parses 'input' into a MeDriveIdListItemIdDocumentSetVersionId
func ParseMeDriveIdListItemIdDocumentSetVersionID(input string) (*MeDriveIdListItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListItemIdDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdListItemIdDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListItemIdDocumentSetVersionIDInsensitively(input string) (*MeDriveIdListItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListItemIdDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeDriveIdListItemIdDocumentSetVersionID checks that 'input' can be parsed as a Me Drive Id List Item Id Document Set Version ID
func ValidateMeDriveIdListItemIdDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListItemIdDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Item Id Document Set Version ID
func (id MeDriveIdListItemIdDocumentSetVersionId) ID() string {
	fmtString := "/me/drives/%s/list/items/%s/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ListItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Item Id Document Set Version ID
func (id MeDriveIdListItemIdDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Item Id Document Set Version ID
func (id MeDriveIdListItemIdDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Me Drive Id List Item Id Document Set Version (%s)", strings.Join(components, "\n"))
}
