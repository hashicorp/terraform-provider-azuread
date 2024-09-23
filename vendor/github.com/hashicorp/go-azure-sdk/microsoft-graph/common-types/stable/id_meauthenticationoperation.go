package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationOperationId{}

// MeAuthenticationOperationId is a struct representing the Resource ID for a Me Authentication Operation
type MeAuthenticationOperationId struct {
	LongRunningOperationId string
}

// NewMeAuthenticationOperationID returns a new MeAuthenticationOperationId struct
func NewMeAuthenticationOperationID(longRunningOperationId string) MeAuthenticationOperationId {
	return MeAuthenticationOperationId{
		LongRunningOperationId: longRunningOperationId,
	}
}

// ParseMeAuthenticationOperationID parses 'input' into a MeAuthenticationOperationId
func ParseMeAuthenticationOperationID(input string) (*MeAuthenticationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationOperationIDInsensitively parses 'input' case-insensitively into a MeAuthenticationOperationId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationOperationIDInsensitively(input string) (*MeAuthenticationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LongRunningOperationId, ok = input.Parsed["longRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longRunningOperationId", input)
	}

	return nil
}

// ValidateMeAuthenticationOperationID checks that 'input' can be parsed as a Me Authentication Operation ID
func ValidateMeAuthenticationOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Operation ID
func (id MeAuthenticationOperationId) ID() string {
	fmtString := "/me/authentication/operations/%s"
	return fmt.Sprintf(fmtString, id.LongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Operation ID
func (id MeAuthenticationOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("longRunningOperationId", "longRunningOperationId"),
	}
}

// String returns a human-readable description of this Me Authentication Operation ID
func (id MeAuthenticationOperationId) String() string {
	components := []string{
		fmt.Sprintf("Long Running Operation: %q", id.LongRunningOperationId),
	}
	return fmt.Sprintf("Me Authentication Operation (%s)", strings.Join(components, "\n"))
}
