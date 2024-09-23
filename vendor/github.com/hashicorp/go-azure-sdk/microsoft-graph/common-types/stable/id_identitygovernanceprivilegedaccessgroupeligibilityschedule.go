package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{}

// IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Eligibility Schedule
type IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId struct {
	PrivilegedAccessGroupEligibilityScheduleId string
}

// NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID returns a new IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId struct
func NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(privilegedAccessGroupEligibilityScheduleId string) IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId {
	return IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{
		PrivilegedAccessGroupEligibilityScheduleId: privilegedAccessGroupEligibilityScheduleId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID parses 'input' into a IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId
func ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(input string) (*IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupEligibilityScheduleId, ok = input.Parsed["privilegedAccessGroupEligibilityScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupEligibilityScheduleId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Eligibility Schedule ID
func ValidateIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Eligibility Schedule ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/eligibilitySchedules/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupEligibilityScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Eligibility Schedule ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("eligibilitySchedules", "eligibilitySchedules", "eligibilitySchedules"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupEligibilityScheduleId", "privilegedAccessGroupEligibilityScheduleId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Eligibility Schedule ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Eligibility Schedule: %q", id.PrivilegedAccessGroupEligibilityScheduleId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Eligibility Schedule (%s)", strings.Join(components, "\n"))
}
