package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdMemberId{}

// AdministrativeUnitIdMemberId is a struct representing the Resource ID for a Administrative Unit Id Member
type AdministrativeUnitIdMemberId struct {
	AdministrativeUnitId string
	DirectoryObjectId    string
}

// NewAdministrativeUnitIdMemberID returns a new AdministrativeUnitIdMemberId struct
func NewAdministrativeUnitIdMemberID(administrativeUnitId string, directoryObjectId string) AdministrativeUnitIdMemberId {
	return AdministrativeUnitIdMemberId{
		AdministrativeUnitId: administrativeUnitId,
		DirectoryObjectId:    directoryObjectId,
	}
}

// ParseAdministrativeUnitIdMemberID parses 'input' into a AdministrativeUnitIdMemberId
func ParseAdministrativeUnitIdMemberID(input string) (*AdministrativeUnitIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdministrativeUnitIdMemberIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitIdMemberId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitIdMemberIDInsensitively(input string) (*AdministrativeUnitIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdministrativeUnitIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateAdministrativeUnitIdMemberID checks that 'input' can be parsed as a Administrative Unit Id Member ID
func ValidateAdministrativeUnitIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Id Member ID
func (id AdministrativeUnitIdMemberId) ID() string {
	fmtString := "/administrativeUnits/%s/members/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Id Member ID
func (id AdministrativeUnitIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Administrative Unit Id Member ID
func (id AdministrativeUnitIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Administrative Unit Id Member (%s)", strings.Join(components, "\n"))
}
