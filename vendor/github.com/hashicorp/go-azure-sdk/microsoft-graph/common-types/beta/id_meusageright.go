package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeUsageRightId{}

// MeUsageRightId is a struct representing the Resource ID for a Me Usage Right
type MeUsageRightId struct {
	UsageRightId string
}

// NewMeUsageRightID returns a new MeUsageRightId struct
func NewMeUsageRightID(usageRightId string) MeUsageRightId {
	return MeUsageRightId{
		UsageRightId: usageRightId,
	}
}

// ParseMeUsageRightID parses 'input' into a MeUsageRightId
func ParseMeUsageRightID(input string) (*MeUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeUsageRightIDInsensitively parses 'input' case-insensitively into a MeUsageRightId
// note: this method should only be used for API response data and not user input
func ParseMeUsageRightIDInsensitively(input string) (*MeUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeUsageRightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UsageRightId, ok = input.Parsed["usageRightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", input)
	}

	return nil
}

// ValidateMeUsageRightID checks that 'input' can be parsed as a Me Usage Right ID
func ValidateMeUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Usage Right ID
func (id MeUsageRightId) ID() string {
	fmtString := "/me/usageRights/%s"
	return fmt.Sprintf(fmtString, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Usage Right ID
func (id MeUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("usageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightId"),
	}
}

// String returns a human-readable description of this Me Usage Right ID
func (id MeUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("Me Usage Right (%s)", strings.Join(components, "\n"))
}
