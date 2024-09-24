package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDirectReportId{}

// MeDirectReportId is a struct representing the Resource ID for a Me Direct Report
type MeDirectReportId struct {
	DirectoryObjectId string
}

// NewMeDirectReportID returns a new MeDirectReportId struct
func NewMeDirectReportID(directoryObjectId string) MeDirectReportId {
	return MeDirectReportId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDirectReportID parses 'input' into a MeDirectReportId
func ParseMeDirectReportID(input string) (*MeDirectReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDirectReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDirectReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDirectReportIDInsensitively parses 'input' case-insensitively into a MeDirectReportId
// note: this method should only be used for API response data and not user input
func ParseMeDirectReportIDInsensitively(input string) (*MeDirectReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDirectReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDirectReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDirectReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeDirectReportID checks that 'input' can be parsed as a Me Direct Report ID
func ValidateMeDirectReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDirectReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Direct Report ID
func (id MeDirectReportId) ID() string {
	fmtString := "/me/directReports/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Direct Report ID
func (id MeDirectReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("directReports", "directReports", "directReports"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Direct Report ID
func (id MeDirectReportId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Direct Report (%s)", strings.Join(components, "\n"))
}
