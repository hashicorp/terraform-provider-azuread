package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMobileAppIntentAndStateId{}

// UserIdMobileAppIntentAndStateId is a struct representing the Resource ID for a User Id Mobile App Intent And State
type UserIdMobileAppIntentAndStateId struct {
	UserId                    string
	MobileAppIntentAndStateId string
}

// NewUserIdMobileAppIntentAndStateID returns a new UserIdMobileAppIntentAndStateId struct
func NewUserIdMobileAppIntentAndStateID(userId string, mobileAppIntentAndStateId string) UserIdMobileAppIntentAndStateId {
	return UserIdMobileAppIntentAndStateId{
		UserId:                    userId,
		MobileAppIntentAndStateId: mobileAppIntentAndStateId,
	}
}

// ParseUserIdMobileAppIntentAndStateID parses 'input' into a UserIdMobileAppIntentAndStateId
func ParseUserIdMobileAppIntentAndStateID(input string) (*UserIdMobileAppIntentAndStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMobileAppIntentAndStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMobileAppIntentAndStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMobileAppIntentAndStateIDInsensitively parses 'input' case-insensitively into a UserIdMobileAppIntentAndStateId
// note: this method should only be used for API response data and not user input
func ParseUserIdMobileAppIntentAndStateIDInsensitively(input string) (*UserIdMobileAppIntentAndStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMobileAppIntentAndStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMobileAppIntentAndStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMobileAppIntentAndStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MobileAppIntentAndStateId, ok = input.Parsed["mobileAppIntentAndStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppIntentAndStateId", input)
	}

	return nil
}

// ValidateUserIdMobileAppIntentAndStateID checks that 'input' can be parsed as a User Id Mobile App Intent And State ID
func ValidateUserIdMobileAppIntentAndStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMobileAppIntentAndStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mobile App Intent And State ID
func (id UserIdMobileAppIntentAndStateId) ID() string {
	fmtString := "/users/%s/mobileAppIntentAndStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MobileAppIntentAndStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mobile App Intent And State ID
func (id UserIdMobileAppIntentAndStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mobileAppIntentAndStates", "mobileAppIntentAndStates", "mobileAppIntentAndStates"),
		resourceids.UserSpecifiedSegment("mobileAppIntentAndStateId", "mobileAppIntentAndStateId"),
	}
}

// String returns a human-readable description of this User Id Mobile App Intent And State ID
func (id UserIdMobileAppIntentAndStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mobile App Intent And State: %q", id.MobileAppIntentAndStateId),
	}
	return fmt.Sprintf("User Id Mobile App Intent And State (%s)", strings.Join(components, "\n"))
}
