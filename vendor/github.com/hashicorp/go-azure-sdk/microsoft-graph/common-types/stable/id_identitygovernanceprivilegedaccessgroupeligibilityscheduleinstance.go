package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{}

// IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Eligibility Schedule Instance
type IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId struct {
	PrivilegedAccessGroupEligibilityScheduleInstanceId string
}

// NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID returns a new IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId struct
func NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(privilegedAccessGroupEligibilityScheduleInstanceId string) IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId {
	return IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{
		PrivilegedAccessGroupEligibilityScheduleInstanceId: privilegedAccessGroupEligibilityScheduleInstanceId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID parses 'input' into a IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId
func ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(input string) (*IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupEligibilityScheduleInstanceId, ok = input.Parsed["privilegedAccessGroupEligibilityScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupEligibilityScheduleInstanceId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Eligibility Schedule Instance ID
func ValidateIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Eligibility Schedule Instance ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupEligibilityScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Eligibility Schedule Instance ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("eligibilityScheduleInstances", "eligibilityScheduleInstances", "eligibilityScheduleInstances"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupEligibilityScheduleInstanceId", "privilegedAccessGroupEligibilityScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Eligibility Schedule Instance ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Eligibility Schedule Instance: %q", id.PrivilegedAccessGroupEligibilityScheduleInstanceId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Eligibility Schedule Instance (%s)", strings.Join(components, "\n"))
}
