package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdDeletedMemberId{}

// AdministrativeUnitIdDeletedMemberId is a struct representing the Resource ID for a Administrative Unit Id Deleted Member
type AdministrativeUnitIdDeletedMemberId struct {
	AdministrativeUnitId string
	DirectoryObjectId    string
}

// NewAdministrativeUnitIdDeletedMemberID returns a new AdministrativeUnitIdDeletedMemberId struct
func NewAdministrativeUnitIdDeletedMemberID(administrativeUnitId string, directoryObjectId string) AdministrativeUnitIdDeletedMemberId {
	return AdministrativeUnitIdDeletedMemberId{
		AdministrativeUnitId: administrativeUnitId,
		DirectoryObjectId:    directoryObjectId,
	}
}

// ParseAdministrativeUnitIdDeletedMemberID parses 'input' into a AdministrativeUnitIdDeletedMemberId
func ParseAdministrativeUnitIdDeletedMemberID(input string) (*AdministrativeUnitIdDeletedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdDeletedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdDeletedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdministrativeUnitIdDeletedMemberIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitIdDeletedMemberId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitIdDeletedMemberIDInsensitively(input string) (*AdministrativeUnitIdDeletedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdDeletedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdDeletedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdministrativeUnitIdDeletedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateAdministrativeUnitIdDeletedMemberID checks that 'input' can be parsed as a Administrative Unit Id Deleted Member ID
func ValidateAdministrativeUnitIdDeletedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitIdDeletedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Id Deleted Member ID
func (id AdministrativeUnitIdDeletedMemberId) ID() string {
	fmtString := "/administrativeUnits/%s/deletedMembers/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Id Deleted Member ID
func (id AdministrativeUnitIdDeletedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("deletedMembers", "deletedMembers", "deletedMembers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Administrative Unit Id Deleted Member ID
func (id AdministrativeUnitIdDeletedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Administrative Unit Id Deleted Member (%s)", strings.Join(components, "\n"))
}
