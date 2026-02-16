package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAdministrativeUnitIdExtensionId{}

// DirectoryAdministrativeUnitIdExtensionId is a struct representing the Resource ID for a Directory Administrative Unit Id Extension
type DirectoryAdministrativeUnitIdExtensionId struct {
	AdministrativeUnitId string
	ExtensionId          string
}

// NewDirectoryAdministrativeUnitIdExtensionID returns a new DirectoryAdministrativeUnitIdExtensionId struct
func NewDirectoryAdministrativeUnitIdExtensionID(administrativeUnitId string, extensionId string) DirectoryAdministrativeUnitIdExtensionId {
	return DirectoryAdministrativeUnitIdExtensionId{
		AdministrativeUnitId: administrativeUnitId,
		ExtensionId:          extensionId,
	}
}

// ParseDirectoryAdministrativeUnitIdExtensionID parses 'input' into a DirectoryAdministrativeUnitIdExtensionId
func ParseDirectoryAdministrativeUnitIdExtensionID(input string) (*DirectoryAdministrativeUnitIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitIdExtensionIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitIdExtensionIDInsensitively(input string) (*DirectoryAdministrativeUnitIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryAdministrativeUnitIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateDirectoryAdministrativeUnitIdExtensionID checks that 'input' can be parsed as a Directory Administrative Unit Id Extension ID
func ValidateDirectoryAdministrativeUnitIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Id Extension ID
func (id DirectoryAdministrativeUnitIdExtensionId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Id Extension ID
func (id DirectoryAdministrativeUnitIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Id Extension ID
func (id DirectoryAdministrativeUnitIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Directory Administrative Unit Id Extension (%s)", strings.Join(components, "\n"))
}
