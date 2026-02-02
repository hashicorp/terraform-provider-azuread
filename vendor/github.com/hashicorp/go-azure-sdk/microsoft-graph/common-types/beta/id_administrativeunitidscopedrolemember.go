package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdScopedRoleMemberId{}

// AdministrativeUnitIdScopedRoleMemberId is a struct representing the Resource ID for a Administrative Unit Id Scoped Role Member
type AdministrativeUnitIdScopedRoleMemberId struct {
	AdministrativeUnitId   string
	ScopedRoleMembershipId string
}

// NewAdministrativeUnitIdScopedRoleMemberID returns a new AdministrativeUnitIdScopedRoleMemberId struct
func NewAdministrativeUnitIdScopedRoleMemberID(administrativeUnitId string, scopedRoleMembershipId string) AdministrativeUnitIdScopedRoleMemberId {
	return AdministrativeUnitIdScopedRoleMemberId{
		AdministrativeUnitId:   administrativeUnitId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseAdministrativeUnitIdScopedRoleMemberID parses 'input' into a AdministrativeUnitIdScopedRoleMemberId
func ParseAdministrativeUnitIdScopedRoleMemberID(input string) (*AdministrativeUnitIdScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdScopedRoleMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdScopedRoleMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdministrativeUnitIdScopedRoleMemberIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitIdScopedRoleMemberId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitIdScopedRoleMemberIDInsensitively(input string) (*AdministrativeUnitIdScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdministrativeUnitIdScopedRoleMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdministrativeUnitIdScopedRoleMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdministrativeUnitIdScopedRoleMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdministrativeUnitId, ok = input.Parsed["administrativeUnitId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", input)
	}

	if id.ScopedRoleMembershipId, ok = input.Parsed["scopedRoleMembershipId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", input)
	}

	return nil
}

// ValidateAdministrativeUnitIdScopedRoleMemberID checks that 'input' can be parsed as a Administrative Unit Id Scoped Role Member ID
func ValidateAdministrativeUnitIdScopedRoleMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitIdScopedRoleMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Id Scoped Role Member ID
func (id AdministrativeUnitIdScopedRoleMemberId) ID() string {
	fmtString := "/administrativeUnits/%s/scopedRoleMembers/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Id Scoped Role Member ID
func (id AdministrativeUnitIdScopedRoleMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("administrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitId"),
		resourceids.StaticSegment("scopedRoleMembers", "scopedRoleMembers", "scopedRoleMembers"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipId"),
	}
}

// String returns a human-readable description of this Administrative Unit Id Scoped Role Member ID
func (id AdministrativeUnitIdScopedRoleMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("Administrative Unit Id Scoped Role Member (%s)", strings.Join(components, "\n"))
}
