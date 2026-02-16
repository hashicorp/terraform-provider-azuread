package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyClaimsMappingPolicyId{}

// PolicyClaimsMappingPolicyId is a struct representing the Resource ID for a Policy Claims Mapping Policy
type PolicyClaimsMappingPolicyId struct {
	ClaimsMappingPolicyId string
}

// NewPolicyClaimsMappingPolicyID returns a new PolicyClaimsMappingPolicyId struct
func NewPolicyClaimsMappingPolicyID(claimsMappingPolicyId string) PolicyClaimsMappingPolicyId {
	return PolicyClaimsMappingPolicyId{
		ClaimsMappingPolicyId: claimsMappingPolicyId,
	}
}

// ParsePolicyClaimsMappingPolicyID parses 'input' into a PolicyClaimsMappingPolicyId
func ParsePolicyClaimsMappingPolicyID(input string) (*PolicyClaimsMappingPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyClaimsMappingPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyClaimsMappingPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyClaimsMappingPolicyIDInsensitively parses 'input' case-insensitively into a PolicyClaimsMappingPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyClaimsMappingPolicyIDInsensitively(input string) (*PolicyClaimsMappingPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyClaimsMappingPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyClaimsMappingPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyClaimsMappingPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ClaimsMappingPolicyId, ok = input.Parsed["claimsMappingPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", input)
	}

	return nil
}

// ValidatePolicyClaimsMappingPolicyID checks that 'input' can be parsed as a Policy Claims Mapping Policy ID
func ValidatePolicyClaimsMappingPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyClaimsMappingPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Claims Mapping Policy ID
func (id PolicyClaimsMappingPolicyId) ID() string {
	fmtString := "/policies/claimsMappingPolicies/%s"
	return fmt.Sprintf(fmtString, id.ClaimsMappingPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Claims Mapping Policy ID
func (id PolicyClaimsMappingPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("claimsMappingPolicies", "claimsMappingPolicies", "claimsMappingPolicies"),
		resourceids.UserSpecifiedSegment("claimsMappingPolicyId", "claimsMappingPolicyId"),
	}
}

// String returns a human-readable description of this Policy Claims Mapping Policy ID
func (id PolicyClaimsMappingPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Claims Mapping Policy: %q", id.ClaimsMappingPolicyId),
	}
	return fmt.Sprintf("Policy Claims Mapping Policy (%s)", strings.Join(components, "\n"))
}
