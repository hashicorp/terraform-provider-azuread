package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdExtensionId{}

// AdministrativeUnitIdExtensionId is a struct representing the Resource ID for a Administrative Unit Id Extension
type AdministrativeUnitIdExtensionId struct {
	AdministrativeUnitId string
	ExtensionId          string
}

// NewAdministrativeUnitIdExtensionID returns a new AdministrativeUnitIdExtensionId struct
func NewAdministrativeUnitIdExtensionID(administrativeUnitId string, extensionId string) AdministrativeUnitIdExtensionId {
	return AdministrativeUnitIdExtensionId{
		AdministrativeUnitId: administrativeUnitId,
		ExtensionId:          extensionId,
	}
}

// ParseAdministrativeUnitIdExtensionID parses 'input' into a AdministrativeUnitIdExtensionId
func ParseAdministrativeUnitIdExtensionID(input string) (*AdministrativeUnitIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdministrativeUnitIdExtensionIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitIdExtensionIDInsensitively(input string) (*AdministrativeUnitIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdministrativeUnitIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateAdministrativeUnitIdExtensionID checks that 'input' can be parsed as a Administrative Unit Id Extension ID
func ValidateAdministrativeUnitIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Id Extension ID
func (id AdministrativeUnitIdExtensionId) ID() string {
	fmtString := "/administrativeUnits/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Id Extension ID
func (id AdministrativeUnitIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Administrative Unit Id Extension ID
func (id AdministrativeUnitIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Administrative Unit Id Extension (%s)", strings.Join(components, "\n"))
}
