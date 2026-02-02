package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyClaimsMappingPolicyIdAppliesToId{}

// PolicyClaimsMappingPolicyIdAppliesToId is a struct representing the Resource ID for a Policy Claims Mapping Policy Id Applies To
type PolicyClaimsMappingPolicyIdAppliesToId struct {
	ClaimsMappingPolicyId string
	DirectoryObjectId     string
}

// NewPolicyClaimsMappingPolicyIdAppliesToID returns a new PolicyClaimsMappingPolicyIdAppliesToId struct
func NewPolicyClaimsMappingPolicyIdAppliesToID(claimsMappingPolicyId string, directoryObjectId string) PolicyClaimsMappingPolicyIdAppliesToId {
	return PolicyClaimsMappingPolicyIdAppliesToId{
		ClaimsMappingPolicyId: claimsMappingPolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyClaimsMappingPolicyIdAppliesToID parses 'input' into a PolicyClaimsMappingPolicyIdAppliesToId
func ParsePolicyClaimsMappingPolicyIdAppliesToID(input string) (*PolicyClaimsMappingPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyClaimsMappingPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyClaimsMappingPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyClaimsMappingPolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyClaimsMappingPolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyClaimsMappingPolicyIdAppliesToIDInsensitively(input string) (*PolicyClaimsMappingPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyClaimsMappingPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyClaimsMappingPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyClaimsMappingPolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ClaimsMappingPolicyId, ok = input.Parsed["claimsMappingPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyClaimsMappingPolicyIdAppliesToID checks that 'input' can be parsed as a Policy Claims Mapping Policy Id Applies To ID
func ValidatePolicyClaimsMappingPolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyClaimsMappingPolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Claims Mapping Policy Id Applies To ID
func (id PolicyClaimsMappingPolicyIdAppliesToId) ID() string {
	fmtString := "/policies/claimsMappingPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.ClaimsMappingPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Claims Mapping Policy Id Applies To ID
func (id PolicyClaimsMappingPolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("claimsMappingPolicies", "claimsMappingPolicies", "claimsMappingPolicies"),
		resourceids.UserSpecifiedSegment("claimsMappingPolicyId", "claimsMappingPolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy Claims Mapping Policy Id Applies To ID
func (id PolicyClaimsMappingPolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Claims Mapping Policy: %q", id.ClaimsMappingPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Claims Mapping Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
