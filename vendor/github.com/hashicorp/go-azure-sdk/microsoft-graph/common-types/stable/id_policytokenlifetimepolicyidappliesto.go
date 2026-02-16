package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyTokenLifetimePolicyIdAppliesToId{}

// PolicyTokenLifetimePolicyIdAppliesToId is a struct representing the Resource ID for a Policy Token Lifetime Policy Id Applies To
type PolicyTokenLifetimePolicyIdAppliesToId struct {
	TokenLifetimePolicyId string
	DirectoryObjectId     string
}

// NewPolicyTokenLifetimePolicyIdAppliesToID returns a new PolicyTokenLifetimePolicyIdAppliesToId struct
func NewPolicyTokenLifetimePolicyIdAppliesToID(tokenLifetimePolicyId string, directoryObjectId string) PolicyTokenLifetimePolicyIdAppliesToId {
	return PolicyTokenLifetimePolicyIdAppliesToId{
		TokenLifetimePolicyId: tokenLifetimePolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyTokenLifetimePolicyIdAppliesToID parses 'input' into a PolicyTokenLifetimePolicyIdAppliesToId
func ParsePolicyTokenLifetimePolicyIdAppliesToID(input string) (*PolicyTokenLifetimePolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyTokenLifetimePolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyTokenLifetimePolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyTokenLifetimePolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyTokenLifetimePolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyTokenLifetimePolicyIdAppliesToIDInsensitively(input string) (*PolicyTokenLifetimePolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyTokenLifetimePolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyTokenLifetimePolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyTokenLifetimePolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TokenLifetimePolicyId, ok = input.Parsed["tokenLifetimePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyTokenLifetimePolicyIdAppliesToID checks that 'input' can be parsed as a Policy Token Lifetime Policy Id Applies To ID
func ValidatePolicyTokenLifetimePolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyTokenLifetimePolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Token Lifetime Policy Id Applies To ID
func (id PolicyTokenLifetimePolicyIdAppliesToId) ID() string {
	fmtString := "/policies/tokenLifetimePolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.TokenLifetimePolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Token Lifetime Policy Id Applies To ID
func (id PolicyTokenLifetimePolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("tokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy Token Lifetime Policy Id Applies To ID
func (id PolicyTokenLifetimePolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Token Lifetime Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
