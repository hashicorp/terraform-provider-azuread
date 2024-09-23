package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyServicePrincipalCreationPolicyId{}

// PolicyServicePrincipalCreationPolicyId is a struct representing the Resource ID for a Policy Service Principal Creation Policy
type PolicyServicePrincipalCreationPolicyId struct {
	ServicePrincipalCreationPolicyId string
}

// NewPolicyServicePrincipalCreationPolicyID returns a new PolicyServicePrincipalCreationPolicyId struct
func NewPolicyServicePrincipalCreationPolicyID(servicePrincipalCreationPolicyId string) PolicyServicePrincipalCreationPolicyId {
	return PolicyServicePrincipalCreationPolicyId{
		ServicePrincipalCreationPolicyId: servicePrincipalCreationPolicyId,
	}
}

// ParsePolicyServicePrincipalCreationPolicyID parses 'input' into a PolicyServicePrincipalCreationPolicyId
func ParsePolicyServicePrincipalCreationPolicyID(input string) (*PolicyServicePrincipalCreationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyServicePrincipalCreationPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyServicePrincipalCreationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyServicePrincipalCreationPolicyIDInsensitively parses 'input' case-insensitively into a PolicyServicePrincipalCreationPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyServicePrincipalCreationPolicyIDInsensitively(input string) (*PolicyServicePrincipalCreationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyServicePrincipalCreationPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyServicePrincipalCreationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyServicePrincipalCreationPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalCreationPolicyId, ok = input.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", input)
	}

	return nil
}

// ValidatePolicyServicePrincipalCreationPolicyID checks that 'input' can be parsed as a Policy Service Principal Creation Policy ID
func ValidatePolicyServicePrincipalCreationPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyServicePrincipalCreationPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Service Principal Creation Policy ID
func (id PolicyServicePrincipalCreationPolicyId) ID() string {
	fmtString := "/policies/servicePrincipalCreationPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalCreationPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Service Principal Creation Policy ID
func (id PolicyServicePrincipalCreationPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationPolicyId", "servicePrincipalCreationPolicyId"),
	}
}

// String returns a human-readable description of this Policy Service Principal Creation Policy ID
func (id PolicyServicePrincipalCreationPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal Creation Policy: %q", id.ServicePrincipalCreationPolicyId),
	}
	return fmt.Sprintf("Policy Service Principal Creation Policy (%s)", strings.Join(components, "\n"))
}
