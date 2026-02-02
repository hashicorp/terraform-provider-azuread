package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyTokenIssuancePolicyIdAppliesToId{}

// PolicyTokenIssuancePolicyIdAppliesToId is a struct representing the Resource ID for a Policy Token Issuance Policy Id Applies To
type PolicyTokenIssuancePolicyIdAppliesToId struct {
	TokenIssuancePolicyId string
	DirectoryObjectId     string
}

// NewPolicyTokenIssuancePolicyIdAppliesToID returns a new PolicyTokenIssuancePolicyIdAppliesToId struct
func NewPolicyTokenIssuancePolicyIdAppliesToID(tokenIssuancePolicyId string, directoryObjectId string) PolicyTokenIssuancePolicyIdAppliesToId {
	return PolicyTokenIssuancePolicyIdAppliesToId{
		TokenIssuancePolicyId: tokenIssuancePolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyTokenIssuancePolicyIdAppliesToID parses 'input' into a PolicyTokenIssuancePolicyIdAppliesToId
func ParsePolicyTokenIssuancePolicyIdAppliesToID(input string) (*PolicyTokenIssuancePolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyTokenIssuancePolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyTokenIssuancePolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyTokenIssuancePolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyTokenIssuancePolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyTokenIssuancePolicyIdAppliesToIDInsensitively(input string) (*PolicyTokenIssuancePolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyTokenIssuancePolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyTokenIssuancePolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyTokenIssuancePolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TokenIssuancePolicyId, ok = input.Parsed["tokenIssuancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyTokenIssuancePolicyIdAppliesToID checks that 'input' can be parsed as a Policy Token Issuance Policy Id Applies To ID
func ValidatePolicyTokenIssuancePolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyTokenIssuancePolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Token Issuance Policy Id Applies To ID
func (id PolicyTokenIssuancePolicyIdAppliesToId) ID() string {
	fmtString := "/policies/tokenIssuancePolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.TokenIssuancePolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Token Issuance Policy Id Applies To ID
func (id PolicyTokenIssuancePolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("tokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy Token Issuance Policy Id Applies To ID
func (id PolicyTokenIssuancePolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Token Issuance Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
