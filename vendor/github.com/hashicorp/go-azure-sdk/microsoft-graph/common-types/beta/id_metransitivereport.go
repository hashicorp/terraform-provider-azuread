package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTransitiveReportId{}

// MeTransitiveReportId is a struct representing the Resource ID for a Me Transitive Report
type MeTransitiveReportId struct {
	DirectoryObjectId string
}

// NewMeTransitiveReportID returns a new MeTransitiveReportId struct
func NewMeTransitiveReportID(directoryObjectId string) MeTransitiveReportId {
	return MeTransitiveReportId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeTransitiveReportID parses 'input' into a MeTransitiveReportId
func ParseMeTransitiveReportID(input string) (*MeTransitiveReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTransitiveReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTransitiveReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTransitiveReportIDInsensitively parses 'input' case-insensitively into a MeTransitiveReportId
// note: this method should only be used for API response data and not user input
func ParseMeTransitiveReportIDInsensitively(input string) (*MeTransitiveReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTransitiveReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTransitiveReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTransitiveReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeTransitiveReportID checks that 'input' can be parsed as a Me Transitive Report ID
func ValidateMeTransitiveReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTransitiveReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Transitive Report ID
func (id MeTransitiveReportId) ID() string {
	fmtString := "/me/transitiveReports/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Transitive Report ID
func (id MeTransitiveReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("transitiveReports", "transitiveReports", "transitiveReports"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Transitive Report ID
func (id MeTransitiveReportId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Transitive Report (%s)", strings.Join(components, "\n"))
}
