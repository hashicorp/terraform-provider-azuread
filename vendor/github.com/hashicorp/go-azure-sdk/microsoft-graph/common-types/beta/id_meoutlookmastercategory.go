package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookMasterCategoryId{}

// MeOutlookMasterCategoryId is a struct representing the Resource ID for a Me Outlook Master Category
type MeOutlookMasterCategoryId struct {
	OutlookCategoryId string
}

// NewMeOutlookMasterCategoryID returns a new MeOutlookMasterCategoryId struct
func NewMeOutlookMasterCategoryID(outlookCategoryId string) MeOutlookMasterCategoryId {
	return MeOutlookMasterCategoryId{
		OutlookCategoryId: outlookCategoryId,
	}
}

// ParseMeOutlookMasterCategoryID parses 'input' into a MeOutlookMasterCategoryId
func ParseMeOutlookMasterCategoryID(input string) (*MeOutlookMasterCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookMasterCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookMasterCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookMasterCategoryIDInsensitively parses 'input' case-insensitively into a MeOutlookMasterCategoryId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookMasterCategoryIDInsensitively(input string) (*MeOutlookMasterCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookMasterCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookMasterCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookMasterCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookCategoryId, ok = input.Parsed["outlookCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookCategoryId", input)
	}

	return nil
}

// ValidateMeOutlookMasterCategoryID checks that 'input' can be parsed as a Me Outlook Master Category ID
func ValidateMeOutlookMasterCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookMasterCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Master Category ID
func (id MeOutlookMasterCategoryId) ID() string {
	fmtString := "/me/outlook/masterCategories/%s"
	return fmt.Sprintf(fmtString, id.OutlookCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Master Category ID
func (id MeOutlookMasterCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("masterCategories", "masterCategories", "masterCategories"),
		resourceids.UserSpecifiedSegment("outlookCategoryId", "outlookCategoryId"),
	}
}

// String returns a human-readable description of this Me Outlook Master Category ID
func (id MeOutlookMasterCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Category: %q", id.OutlookCategoryId),
	}
	return fmt.Sprintf("Me Outlook Master Category (%s)", strings.Join(components, "\n"))
}
