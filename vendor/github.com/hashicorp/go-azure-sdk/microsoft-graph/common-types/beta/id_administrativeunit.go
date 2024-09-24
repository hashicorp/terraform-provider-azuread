package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitId{}

// AdministrativeUnitId is a struct representing the Resource ID for a Administrative Unit
type AdministrativeUnitId struct {
	AdministrativeUnitId string
}

// NewAdministrativeUnitID returns a new AdministrativeUnitId struct
func NewAdministrativeUnitID(administrativeUnitId string) AdministrativeUnitId {
	return AdministrativeUnitId{
		AdministrativeUnitId: administrativeUnitId,
	}
}

// ParseAdministrativeUnitID parses 'input' into a AdministrativeUnitId
func ParseAdministrativeUnitID(input string) (*AdministrativeUnitId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdministrativeUnitIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitIDInsensitively(input string) (*AdministrativeUnitId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdministrativeUnitId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	return nil
}

// ValidateAdministrativeUnitID checks that 'input' can be parsed as a Administrative Unit ID
func ValidateAdministrativeUnitID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit ID
func (id AdministrativeUnitId) ID() string {
	fmtString := "/administrativeUnits/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit ID
func (id AdministrativeUnitId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
	}
}

// String returns a human-readable description of this Administrative Unit ID
func (id AdministrativeUnitId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
	}
	return fmt.Sprintf("Administrative Unit (%s)", strings.Join(components, "\n"))
}
