package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileDeviceManagementPolicyId{}

// PolicyMobileDeviceManagementPolicyId is a struct representing the Resource ID for a Policy Mobile Device Management Policy
type PolicyMobileDeviceManagementPolicyId struct {
	MobilityManagementPolicyId string
}

// NewPolicyMobileDeviceManagementPolicyID returns a new PolicyMobileDeviceManagementPolicyId struct
func NewPolicyMobileDeviceManagementPolicyID(mobilityManagementPolicyId string) PolicyMobileDeviceManagementPolicyId {
	return PolicyMobileDeviceManagementPolicyId{
		MobilityManagementPolicyId: mobilityManagementPolicyId,
	}
}

// ParsePolicyMobileDeviceManagementPolicyID parses 'input' into a PolicyMobileDeviceManagementPolicyId
func ParsePolicyMobileDeviceManagementPolicyID(input string) (*PolicyMobileDeviceManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileDeviceManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileDeviceManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyMobileDeviceManagementPolicyIDInsensitively parses 'input' case-insensitively into a PolicyMobileDeviceManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyMobileDeviceManagementPolicyIDInsensitively(input string) (*PolicyMobileDeviceManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileDeviceManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileDeviceManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyMobileDeviceManagementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobilityManagementPolicyId, ok = input.Parsed["mobilityManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", input)
	}

	return nil
}

// ValidatePolicyMobileDeviceManagementPolicyID checks that 'input' can be parsed as a Policy Mobile Device Management Policy ID
func ValidatePolicyMobileDeviceManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMobileDeviceManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Mobile Device Management Policy ID
func (id PolicyMobileDeviceManagementPolicyId) ID() string {
	fmtString := "/policies/mobileDeviceManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.MobilityManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Mobile Device Management Policy ID
func (id PolicyMobileDeviceManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("mobileDeviceManagementPolicies", "mobileDeviceManagementPolicies", "mobileDeviceManagementPolicies"),
		resourceids.UserSpecifiedSegment("mobilityManagementPolicyId", "mobilityManagementPolicyId"),
	}
}

// String returns a human-readable description of this Policy Mobile Device Management Policy ID
func (id PolicyMobileDeviceManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Mobility Management Policy: %q", id.MobilityManagementPolicyId),
	}
	return fmt.Sprintf("Policy Mobile Device Management Policy (%s)", strings.Join(components, "\n"))
}
