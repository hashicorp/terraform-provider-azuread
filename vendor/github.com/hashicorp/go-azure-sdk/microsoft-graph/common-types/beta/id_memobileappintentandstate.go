package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMobileAppIntentAndStateId{}

// MeMobileAppIntentAndStateId is a struct representing the Resource ID for a Me Mobile App Intent And State
type MeMobileAppIntentAndStateId struct {
	MobileAppIntentAndStateId string
}

// NewMeMobileAppIntentAndStateID returns a new MeMobileAppIntentAndStateId struct
func NewMeMobileAppIntentAndStateID(mobileAppIntentAndStateId string) MeMobileAppIntentAndStateId {
	return MeMobileAppIntentAndStateId{
		MobileAppIntentAndStateId: mobileAppIntentAndStateId,
	}
}

// ParseMeMobileAppIntentAndStateID parses 'input' into a MeMobileAppIntentAndStateId
func ParseMeMobileAppIntentAndStateID(input string) (*MeMobileAppIntentAndStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMobileAppIntentAndStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMobileAppIntentAndStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMobileAppIntentAndStateIDInsensitively parses 'input' case-insensitively into a MeMobileAppIntentAndStateId
// note: this method should only be used for API response data and not user input
func ParseMeMobileAppIntentAndStateIDInsensitively(input string) (*MeMobileAppIntentAndStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMobileAppIntentAndStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMobileAppIntentAndStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMobileAppIntentAndStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobileAppIntentAndStateId, ok = input.Parsed["mobileAppIntentAndStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppIntentAndStateId", input)
	}

	return nil
}

// ValidateMeMobileAppIntentAndStateID checks that 'input' can be parsed as a Me Mobile App Intent And State ID
func ValidateMeMobileAppIntentAndStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMobileAppIntentAndStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mobile App Intent And State ID
func (id MeMobileAppIntentAndStateId) ID() string {
	fmtString := "/me/mobileAppIntentAndStates/%s"
	return fmt.Sprintf(fmtString, id.MobileAppIntentAndStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mobile App Intent And State ID
func (id MeMobileAppIntentAndStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mobileAppIntentAndStates", "mobileAppIntentAndStates", "mobileAppIntentAndStates"),
		resourceids.UserSpecifiedSegment("mobileAppIntentAndStateId", "mobileAppIntentAndStateId"),
	}
}

// String returns a human-readable description of this Me Mobile App Intent And State ID
func (id MeMobileAppIntentAndStateId) String() string {
	components := []string{
		fmt.Sprintf("Mobile App Intent And State: %q", id.MobileAppIntentAndStateId),
	}
	return fmt.Sprintf("Me Mobile App Intent And State (%s)", strings.Join(components, "\n"))
}
