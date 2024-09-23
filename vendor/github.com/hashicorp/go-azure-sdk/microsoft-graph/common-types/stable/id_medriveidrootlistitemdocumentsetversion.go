package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootListItemDocumentSetVersionId{}

// MeDriveIdRootListItemDocumentSetVersionId is a struct representing the Resource ID for a Me Drive Id Root List Item Document Set Version
type MeDriveIdRootListItemDocumentSetVersionId struct {
	DriveId              string
	DocumentSetVersionId string
}

// NewMeDriveIdRootListItemDocumentSetVersionID returns a new MeDriveIdRootListItemDocumentSetVersionId struct
func NewMeDriveIdRootListItemDocumentSetVersionID(driveId string, documentSetVersionId string) MeDriveIdRootListItemDocumentSetVersionId {
	return MeDriveIdRootListItemDocumentSetVersionId{
		DriveId:              driveId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseMeDriveIdRootListItemDocumentSetVersionID parses 'input' into a MeDriveIdRootListItemDocumentSetVersionId
func ParseMeDriveIdRootListItemDocumentSetVersionID(input string) (*MeDriveIdRootListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootListItemDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootListItemDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootListItemDocumentSetVersionIDInsensitively(input string) (*MeDriveIdRootListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootListItemDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootListItemDocumentSetVersionID checks that 'input' can be parsed as a Me Drive Id Root List Item Document Set Version ID
func ValidateMeDriveIdRootListItemDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootListItemDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root List Item Document Set Version ID
func (id MeDriveIdRootListItemDocumentSetVersionId) ID() string {
	fmtString := "/me/drives/%s/root/listItem/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root List Item Document Set Version ID
func (id MeDriveIdRootListItemDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root List Item Document Set Version ID
func (id MeDriveIdRootListItemDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Me Drive Id Root List Item Document Set Version (%s)", strings.Join(components, "\n"))
}
