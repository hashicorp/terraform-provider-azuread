package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAdministrativeUnitIdScopedRoleMemberId{}

// DirectoryAdministrativeUnitIdScopedRoleMemberId is a struct representing the Resource ID for a Directory Administrative Unit Id Scoped Role Member
type DirectoryAdministrativeUnitIdScopedRoleMemberId struct {
	AdministrativeUnitId   string
	ScopedRoleMembershipId string
}

// NewDirectoryAdministrativeUnitIdScopedRoleMemberID returns a new DirectoryAdministrativeUnitIdScopedRoleMemberId struct
func NewDirectoryAdministrativeUnitIdScopedRoleMemberID(administrativeUnitId string, scopedRoleMembershipId string) DirectoryAdministrativeUnitIdScopedRoleMemberId {
	return DirectoryAdministrativeUnitIdScopedRoleMemberId{
		AdministrativeUnitId:   administrativeUnitId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseDirectoryAdministrativeUnitIdScopedRoleMemberID parses 'input' into a DirectoryAdministrativeUnitIdScopedRoleMemberId
func ParseDirectoryAdministrativeUnitIdScopedRoleMemberID(input string) (*DirectoryAdministrativeUnitIdScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdScopedRoleMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdScopedRoleMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitIdScopedRoleMemberIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitIdScopedRoleMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitIdScopedRoleMemberIDInsensitively(input string) (*DirectoryAdministrativeUnitIdScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAdministrativeUnitIdScopedRoleMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAdministrativeUnitIdScopedRoleMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryAdministrativeUnitIdScopedRoleMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.ScopedRoleMembershipId, ok = input.Parsed["scopedRoleMembershipId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", input)
	}

	return nil
}

// ValidateDirectoryAdministrativeUnitIdScopedRoleMemberID checks that 'input' can be parsed as a Directory Administrative Unit Id Scoped Role Member ID
func ValidateDirectoryAdministrativeUnitIdScopedRoleMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitIdScopedRoleMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Id Scoped Role Member ID
func (id DirectoryAdministrativeUnitIdScopedRoleMemberId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/scopedRoleMembers/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Id Scoped Role Member ID
func (id DirectoryAdministrativeUnitIdScopedRoleMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("scopedRoleMembers", "scopedRoleMembers", "scopedRoleMembers"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipId"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Id Scoped Role Member ID
func (id DirectoryAdministrativeUnitIdScopedRoleMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("Directory Administrative Unit Id Scoped Role Member (%s)", strings.Join(components, "\n"))
}
