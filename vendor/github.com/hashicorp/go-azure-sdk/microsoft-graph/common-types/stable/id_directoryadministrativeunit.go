package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAdministrativeUnitId{}

// DirectoryAdministrativeUnitId is a struct representing the Resource ID for a Directory Administrative Unit
type DirectoryAdministrativeUnitId struct {
	AdministrativeUnitId string
}

// NewDirectoryAdministrativeUnitID returns a new DirectoryAdministrativeUnitId struct
func NewDirectoryAdministrativeUnitID(administrativeUnitId string) DirectoryAdministrativeUnitId {
	return DirectoryAdministrativeUnitId{
		AdministrativeUnitId: administrativeUnitId,
	}
}

// ParseDirectoryAdministrativeUnitID parses 'input' into a DirectoryAdministrativeUnitId
func ParseDirectoryAdministrativeUnitID(input string) (*DirectoryAdministrativeUnitId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitIDInsensitively(input string) (*DirectoryAdministrativeUnitId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryAdministrativeUnitId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	return nil
}

// ValidateDirectoryAdministrativeUnitID checks that 'input' can be parsed as a Directory Administrative Unit ID
func ValidateDirectoryAdministrativeUnitID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit ID
func (id DirectoryAdministrativeUnitId) ID() string {
	fmtString := "/directory/administrativeUnits/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit ID
func (id DirectoryAdministrativeUnitId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit ID
func (id DirectoryAdministrativeUnitId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
	}
	return fmt.Sprintf("Directory Administrative Unit (%s)", strings.Join(components, "\n"))
}
