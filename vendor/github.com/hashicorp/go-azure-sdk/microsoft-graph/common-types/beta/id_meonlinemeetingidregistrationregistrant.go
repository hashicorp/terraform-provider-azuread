package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdRegistrationRegistrantId{}

// MeOnlineMeetingIdRegistrationRegistrantId is a struct representing the Resource ID for a Me Online Meeting Id Registration Registrant
type MeOnlineMeetingIdRegistrationRegistrantId struct {
	OnlineMeetingId         string
	MeetingRegistrantBaseId string
}

// NewMeOnlineMeetingIdRegistrationRegistrantID returns a new MeOnlineMeetingIdRegistrationRegistrantId struct
func NewMeOnlineMeetingIdRegistrationRegistrantID(onlineMeetingId string, meetingRegistrantBaseId string) MeOnlineMeetingIdRegistrationRegistrantId {
	return MeOnlineMeetingIdRegistrationRegistrantId{
		OnlineMeetingId:         onlineMeetingId,
		MeetingRegistrantBaseId: meetingRegistrantBaseId,
	}
}

// ParseMeOnlineMeetingIdRegistrationRegistrantID parses 'input' into a MeOnlineMeetingIdRegistrationRegistrantId
func ParseMeOnlineMeetingIdRegistrationRegistrantID(input string) (*MeOnlineMeetingIdRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdRegistrationRegistrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdRegistrationRegistrantIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdRegistrationRegistrantId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdRegistrationRegistrantIDInsensitively(input string) (*MeOnlineMeetingIdRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdRegistrationRegistrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdRegistrationRegistrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.MeetingRegistrantBaseId, ok = input.Parsed["meetingRegistrantBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrantBaseId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdRegistrationRegistrantID checks that 'input' can be parsed as a Me Online Meeting Id Registration Registrant ID
func ValidateMeOnlineMeetingIdRegistrationRegistrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdRegistrationRegistrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Registration Registrant ID
func (id MeOnlineMeetingIdRegistrationRegistrantId) ID() string {
	fmtString := "/me/onlineMeetings/%s/registration/registrants/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingRegistrantBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Registration Registrant ID
func (id MeOnlineMeetingIdRegistrationRegistrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("registration", "registration", "registration"),
		resourceids.StaticSegment("registrants", "registrants", "registrants"),
		resourceids.UserSpecifiedSegment("meetingRegistrantBaseId", "meetingRegistrantBaseId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Registration Registrant ID
func (id MeOnlineMeetingIdRegistrationRegistrantId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registrant Base: %q", id.MeetingRegistrantBaseId),
	}
	return fmt.Sprintf("Me Online Meeting Id Registration Registrant (%s)", strings.Join(components, "\n"))
}
