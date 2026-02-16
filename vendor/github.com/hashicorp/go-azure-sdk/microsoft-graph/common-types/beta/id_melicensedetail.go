package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeLicenseDetailId{}

// MeLicenseDetailId is a struct representing the Resource ID for a Me License Detail
type MeLicenseDetailId struct {
	LicenseDetailsId string
}

// NewMeLicenseDetailID returns a new MeLicenseDetailId struct
func NewMeLicenseDetailID(licenseDetailsId string) MeLicenseDetailId {
	return MeLicenseDetailId{
		LicenseDetailsId: licenseDetailsId,
	}
}

// ParseMeLicenseDetailID parses 'input' into a MeLicenseDetailId
func ParseMeLicenseDetailID(input string) (*MeLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeLicenseDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeLicenseDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeLicenseDetailIDInsensitively parses 'input' case-insensitively into a MeLicenseDetailId
// note: this method should only be used for API response data and not user input
func ParseMeLicenseDetailIDInsensitively(input string) (*MeLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeLicenseDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeLicenseDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeLicenseDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LicenseDetailsId, ok = input.Parsed["licenseDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", input)
	}

	return nil
}

// ValidateMeLicenseDetailID checks that 'input' can be parsed as a Me License Detail ID
func ValidateMeLicenseDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeLicenseDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me License Detail ID
func (id MeLicenseDetailId) ID() string {
	fmtString := "/me/licenseDetails/%s"
	return fmt.Sprintf(fmtString, id.LicenseDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me License Detail ID
func (id MeLicenseDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("licenseDetails", "licenseDetails", "licenseDetails"),
		resourceids.UserSpecifiedSegment("licenseDetailsId", "licenseDetailsId"),
	}
}

// String returns a human-readable description of this Me License Detail ID
func (id MeLicenseDetailId) String() string {
	components := []string{
		fmt.Sprintf("License Details: %q", id.LicenseDetailsId),
	}
	return fmt.Sprintf("Me License Detail (%s)", strings.Join(components, "\n"))
}
