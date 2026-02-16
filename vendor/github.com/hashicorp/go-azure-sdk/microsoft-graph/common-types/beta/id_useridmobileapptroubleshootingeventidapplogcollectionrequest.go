package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}

// UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId is a struct representing the Resource ID for a User Id Mobile App Troubleshooting Event Id App Log Collection Request
type UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId struct {
	UserId                          string
	MobileAppTroubleshootingEventId string
	AppLogCollectionRequestId       string
}

// NewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID returns a new UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId struct
func NewUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(userId string, mobileAppTroubleshootingEventId string, appLogCollectionRequestId string) UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId {
	return UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
		UserId:                          userId,
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
		AppLogCollectionRequestId:       appLogCollectionRequestId,
	}
}

// ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID parses 'input' into a UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId
func ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(input string) (*UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(input string) (*UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MobileAppTroubleshootingEventId, ok = input.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", input)
	}

	if id.AppLogCollectionRequestId, ok = input.Parsed["appLogCollectionRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", input)
	}

	return nil
}

// ValidateUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID checks that 'input' can be parsed as a User Id Mobile App Troubleshooting Event Id App Log Collection Request ID
func ValidateUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId) ID() string {
	fmtString := "/users/%s/mobileAppTroubleshootingEvents/%s/appLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MobileAppTroubleshootingEventId, id.AppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventId"),
		resourceids.StaticSegment("appLogCollectionRequests", "appLogCollectionRequests", "appLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("appLogCollectionRequestId", "appLogCollectionRequestId"),
	}
}

// String returns a human-readable description of this User Id Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
		fmt.Sprintf("App Log Collection Request: %q", id.AppLogCollectionRequestId),
	}
	return fmt.Sprintf("User Id Mobile App Troubleshooting Event Id App Log Collection Request (%s)", strings.Join(components, "\n"))
}
