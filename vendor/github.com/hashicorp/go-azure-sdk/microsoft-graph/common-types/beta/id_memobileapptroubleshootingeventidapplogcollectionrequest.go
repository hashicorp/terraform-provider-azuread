package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}

// MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId is a struct representing the Resource ID for a Me Mobile App Troubleshooting Event Id App Log Collection Request
type MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId struct {
	MobileAppTroubleshootingEventId string
	AppLogCollectionRequestId       string
}

// NewMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID returns a new MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId struct
func NewMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID(mobileAppTroubleshootingEventId string, appLogCollectionRequestId string) MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId {
	return MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
		AppLogCollectionRequestId:       appLogCollectionRequestId,
	}
}

// ParseMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID parses 'input' into a MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId
func ParseMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID(input string) (*MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseMeMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(input string) (*MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobileAppTroubleshootingEventId, ok = input.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", input)
	}

	if id.AppLogCollectionRequestId, ok = input.Parsed["appLogCollectionRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", input)
	}

	return nil
}

// ValidateMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID checks that 'input' can be parsed as a Me Mobile App Troubleshooting Event Id App Log Collection Request ID
func ValidateMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMobileAppTroubleshootingEventIdAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId) ID() string {
	fmtString := "/me/mobileAppTroubleshootingEvents/%s/appLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.MobileAppTroubleshootingEventId, id.AppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventId"),
		resourceids.StaticSegment("appLogCollectionRequests", "appLogCollectionRequests", "appLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("appLogCollectionRequestId", "appLogCollectionRequestId"),
	}
}

// String returns a human-readable description of this Me Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id MeMobileAppTroubleshootingEventIdAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
		fmt.Sprintf("App Log Collection Request: %q", id.AppLogCollectionRequestId),
	}
	return fmt.Sprintf("Me Mobile App Troubleshooting Event Id App Log Collection Request (%s)", strings.Join(components, "\n"))
}
