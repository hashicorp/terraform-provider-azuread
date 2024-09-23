package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{}

// IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Eligibility Schedule Request
type IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId struct {
	PrivilegedAccessGroupEligibilityScheduleRequestId string
}

// NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID returns a new IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId struct
func NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(privilegedAccessGroupEligibilityScheduleRequestId string) IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId {
	return IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{
		PrivilegedAccessGroupEligibilityScheduleRequestId: privilegedAccessGroupEligibilityScheduleRequestId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID parses 'input' into a IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId
func ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(input string) (*IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupEligibilityScheduleRequestId, ok = input.Parsed["privilegedAccessGroupEligibilityScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupEligibilityScheduleRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Eligibility Schedule Request ID
func ValidateIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Eligibility Schedule Request ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupEligibilityScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Eligibility Schedule Request ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("eligibilityScheduleRequests", "eligibilityScheduleRequests", "eligibilityScheduleRequests"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupEligibilityScheduleRequestId", "privilegedAccessGroupEligibilityScheduleRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Eligibility Schedule Request ID
func (id IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Eligibility Schedule Request: %q", id.PrivilegedAccessGroupEligibilityScheduleRequestId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Eligibility Schedule Request (%s)", strings.Join(components, "\n"))
}
