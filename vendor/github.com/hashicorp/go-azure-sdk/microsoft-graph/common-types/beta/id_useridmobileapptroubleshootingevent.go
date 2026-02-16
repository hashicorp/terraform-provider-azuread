package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMobileAppTroubleshootingEventId{}

// UserIdMobileAppTroubleshootingEventId is a struct representing the Resource ID for a User Id Mobile App Troubleshooting Event
type UserIdMobileAppTroubleshootingEventId struct {
	UserId                          string
	MobileAppTroubleshootingEventId string
}

// NewUserIdMobileAppTroubleshootingEventID returns a new UserIdMobileAppTroubleshootingEventId struct
func NewUserIdMobileAppTroubleshootingEventID(userId string, mobileAppTroubleshootingEventId string) UserIdMobileAppTroubleshootingEventId {
	return UserIdMobileAppTroubleshootingEventId{
		UserId:                          userId,
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
	}
}

// ParseUserIdMobileAppTroubleshootingEventID parses 'input' into a UserIdMobileAppTroubleshootingEventId
func ParseUserIdMobileAppTroubleshootingEventID(input string) (*UserIdMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMobileAppTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMobileAppTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a UserIdMobileAppTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseUserIdMobileAppTroubleshootingEventIDInsensitively(input string) (*UserIdMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMobileAppTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMobileAppTroubleshootingEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MobileAppTroubleshootingEventId, ok = input.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", input)
	}

	return nil
}

// ValidateUserIdMobileAppTroubleshootingEventID checks that 'input' can be parsed as a User Id Mobile App Troubleshooting Event ID
func ValidateUserIdMobileAppTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMobileAppTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mobile App Troubleshooting Event ID
func (id UserIdMobileAppTroubleshootingEventId) ID() string {
	fmtString := "/users/%s/mobileAppTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MobileAppTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mobile App Troubleshooting Event ID
func (id UserIdMobileAppTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventId"),
	}
}

// String returns a human-readable description of this User Id Mobile App Troubleshooting Event ID
func (id UserIdMobileAppTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
	}
	return fmt.Sprintf("User Id Mobile App Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
