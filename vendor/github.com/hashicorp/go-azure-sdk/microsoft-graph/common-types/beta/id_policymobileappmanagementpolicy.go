package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileAppManagementPolicyId{}

// PolicyMobileAppManagementPolicyId is a struct representing the Resource ID for a Policy Mobile App Management Policy
type PolicyMobileAppManagementPolicyId struct {
	MobilityManagementPolicyId string
}

// NewPolicyMobileAppManagementPolicyID returns a new PolicyMobileAppManagementPolicyId struct
func NewPolicyMobileAppManagementPolicyID(mobilityManagementPolicyId string) PolicyMobileAppManagementPolicyId {
	return PolicyMobileAppManagementPolicyId{
		MobilityManagementPolicyId: mobilityManagementPolicyId,
	}
}

// ParsePolicyMobileAppManagementPolicyID parses 'input' into a PolicyMobileAppManagementPolicyId
func ParsePolicyMobileAppManagementPolicyID(input string) (*PolicyMobileAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileAppManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyMobileAppManagementPolicyIDInsensitively parses 'input' case-insensitively into a PolicyMobileAppManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyMobileAppManagementPolicyIDInsensitively(input string) (*PolicyMobileAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileAppManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyMobileAppManagementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobilityManagementPolicyId, ok = input.Parsed["mobilityManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", input)
	}

	return nil
}

// ValidatePolicyMobileAppManagementPolicyID checks that 'input' can be parsed as a Policy Mobile App Management Policy ID
func ValidatePolicyMobileAppManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMobileAppManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Mobile App Management Policy ID
func (id PolicyMobileAppManagementPolicyId) ID() string {
	fmtString := "/policies/mobileAppManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.MobilityManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Mobile App Management Policy ID
func (id PolicyMobileAppManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("mobileAppManagementPolicies", "mobileAppManagementPolicies", "mobileAppManagementPolicies"),
		resourceids.UserSpecifiedSegment("mobilityManagementPolicyId", "mobilityManagementPolicyId"),
	}
}

// String returns a human-readable description of this Policy Mobile App Management Policy ID
func (id PolicyMobileAppManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Mobility Management Policy: %q", id.MobilityManagementPolicyId),
	}
	return fmt.Sprintf("Policy Mobile App Management Policy (%s)", strings.Join(components, "\n"))
}
