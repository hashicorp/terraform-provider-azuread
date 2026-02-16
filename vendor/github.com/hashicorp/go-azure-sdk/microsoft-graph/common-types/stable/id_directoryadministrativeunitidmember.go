package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAdministrativeUnitIdMemberId{}

// DirectoryAdministrativeUnitIdMemberId is a struct representing the Resource ID for a Directory Administrative Unit Id Member
type DirectoryAdministrativeUnitIdMemberId struct {
	AdministrativeUnitId string
	DirectoryObjectId    string
}

// NewDirectoryAdministrativeUnitIdMemberID returns a new DirectoryAdministrativeUnitIdMemberId struct
func NewDirectoryAdministrativeUnitIdMemberID(administrativeUnitId string, directoryObjectId string) DirectoryAdministrativeUnitIdMemberId {
	return DirectoryAdministrativeUnitIdMemberId{
		AdministrativeUnitId: administrativeUnitId,
		DirectoryObjectId:    directoryObjectId,
	}
}

// ParseDirectoryAdministrativeUnitIdMemberID parses 'input' into a DirectoryAdministrativeUnitIdMemberId
func ParseDirectoryAdministrativeUnitIdMemberID(input string) (*DirectoryAdministrativeUnitIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitIdMemberIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitIdMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitIdMemberIDInsensitively(input string) (*DirectoryAdministrativeUnitIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryAdministrativeUnitIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryAdministrativeUnitIdMemberID checks that 'input' can be parsed as a Directory Administrative Unit Id Member ID
func ValidateDirectoryAdministrativeUnitIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Id Member ID
func (id DirectoryAdministrativeUnitIdMemberId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/members/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Id Member ID
func (id DirectoryAdministrativeUnitIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Id Member ID
func (id DirectoryAdministrativeUnitIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Administrative Unit Id Member (%s)", strings.Join(components, "\n"))
}
