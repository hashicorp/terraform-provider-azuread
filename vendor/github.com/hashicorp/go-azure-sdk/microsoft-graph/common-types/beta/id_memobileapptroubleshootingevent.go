package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMobileAppTroubleshootingEventId{}

// MeMobileAppTroubleshootingEventId is a struct representing the Resource ID for a Me Mobile App Troubleshooting Event
type MeMobileAppTroubleshootingEventId struct {
	MobileAppTroubleshootingEventId string
}

// NewMeMobileAppTroubleshootingEventID returns a new MeMobileAppTroubleshootingEventId struct
func NewMeMobileAppTroubleshootingEventID(mobileAppTroubleshootingEventId string) MeMobileAppTroubleshootingEventId {
	return MeMobileAppTroubleshootingEventId{
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
	}
}

// ParseMeMobileAppTroubleshootingEventID parses 'input' into a MeMobileAppTroubleshootingEventId
func ParseMeMobileAppTroubleshootingEventID(input string) (*MeMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMobileAppTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMobileAppTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a MeMobileAppTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseMeMobileAppTroubleshootingEventIDInsensitively(input string) (*MeMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMobileAppTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMobileAppTroubleshootingEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobileAppTroubleshootingEventId, ok = input.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", input)
	}

	return nil
}

// ValidateMeMobileAppTroubleshootingEventID checks that 'input' can be parsed as a Me Mobile App Troubleshooting Event ID
func ValidateMeMobileAppTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMobileAppTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mobile App Troubleshooting Event ID
func (id MeMobileAppTroubleshootingEventId) ID() string {
	fmtString := "/me/mobileAppTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.MobileAppTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mobile App Troubleshooting Event ID
func (id MeMobileAppTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventId"),
	}
}

// String returns a human-readable description of this Me Mobile App Troubleshooting Event ID
func (id MeMobileAppTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
	}
	return fmt.Sprintf("Me Mobile App Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
