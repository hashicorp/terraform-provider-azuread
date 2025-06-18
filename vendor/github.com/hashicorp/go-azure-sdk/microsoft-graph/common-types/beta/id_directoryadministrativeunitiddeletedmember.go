package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAdministrativeUnitIdDeletedMemberId{}

// DirectoryAdministrativeUnitIdDeletedMemberId is a struct representing the Resource ID for a Directory Administrative Unit Id Deleted Member
type DirectoryAdministrativeUnitIdDeletedMemberId struct {
	AdministrativeUnitId string
	DirectoryObjectId    string
}

// NewDirectoryAdministrativeUnitIdDeletedMemberID returns a new DirectoryAdministrativeUnitIdDeletedMemberId struct
func NewDirectoryAdministrativeUnitIdDeletedMemberID(administrativeUnitId string, directoryObjectId string) DirectoryAdministrativeUnitIdDeletedMemberId {
	return DirectoryAdministrativeUnitIdDeletedMemberId{
		AdministrativeUnitId: administrativeUnitId,
		DirectoryObjectId:    directoryObjectId,
	}
}

// ParseDirectoryAdministrativeUnitIdDeletedMemberID parses 'input' into a DirectoryAdministrativeUnitIdDeletedMemberId
func ParseDirectoryAdministrativeUnitIdDeletedMemberID(input string) (*DirectoryAdministrativeUnitIdDeletedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdDeletedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdDeletedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitIdDeletedMemberIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitIdDeletedMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitIdDeletedMemberIDInsensitively(input string) (*DirectoryAdministrativeUnitIdDeletedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdDeletedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdDeletedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryAdministrativeUnitIdDeletedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryAdministrativeUnitIdDeletedMemberID checks that 'input' can be parsed as a Directory Administrative Unit Id Deleted Member ID
func ValidateDirectoryAdministrativeUnitIdDeletedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitIdDeletedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Id Deleted Member ID
func (id DirectoryAdministrativeUnitIdDeletedMemberId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/deletedMembers/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Id Deleted Member ID
func (id DirectoryAdministrativeUnitIdDeletedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("deletedMembers", "deletedMembers", "deletedMembers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Id Deleted Member ID
func (id DirectoryAdministrativeUnitIdDeletedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Administrative Unit Id Deleted Member (%s)", strings.Join(components, "\n"))
}
