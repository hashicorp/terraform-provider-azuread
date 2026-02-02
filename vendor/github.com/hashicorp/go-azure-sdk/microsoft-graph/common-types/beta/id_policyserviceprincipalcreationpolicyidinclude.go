package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyServicePrincipalCreationPolicyIdIncludeId{}

// PolicyServicePrincipalCreationPolicyIdIncludeId is a struct representing the Resource ID for a Policy Service Principal Creation Policy Id Include
type PolicyServicePrincipalCreationPolicyIdIncludeId struct {
	ServicePrincipalCreationPolicyId       string
	ServicePrincipalCreationConditionSetId string
}

// NewPolicyServicePrincipalCreationPolicyIdIncludeID returns a new PolicyServicePrincipalCreationPolicyIdIncludeId struct
func NewPolicyServicePrincipalCreationPolicyIdIncludeID(servicePrincipalCreationPolicyId string, servicePrincipalCreationConditionSetId string) PolicyServicePrincipalCreationPolicyIdIncludeId {
	return PolicyServicePrincipalCreationPolicyIdIncludeId{
		ServicePrincipalCreationPolicyId:       servicePrincipalCreationPolicyId,
		ServicePrincipalCreationConditionSetId: servicePrincipalCreationConditionSetId,
	}
}

// ParsePolicyServicePrincipalCreationPolicyIdIncludeID parses 'input' into a PolicyServicePrincipalCreationPolicyIdIncludeId
func ParsePolicyServicePrincipalCreationPolicyIdIncludeID(input string) (*PolicyServicePrincipalCreationPolicyIdIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyServicePrincipalCreationPolicyIdIncludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyServicePrincipalCreationPolicyIdIncludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyServicePrincipalCreationPolicyIdIncludeIDInsensitively parses 'input' case-insensitively into a PolicyServicePrincipalCreationPolicyIdIncludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyServicePrincipalCreationPolicyIdIncludeIDInsensitively(input string) (*PolicyServicePrincipalCreationPolicyIdIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyServicePrincipalCreationPolicyIdIncludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyServicePrincipalCreationPolicyIdIncludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyServicePrincipalCreationPolicyIdIncludeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalCreationPolicyId, ok = input.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", input)
	}

	if id.ServicePrincipalCreationConditionSetId, ok = input.Parsed["servicePrincipalCreationConditionSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationConditionSetId", input)
	}

	return nil
}

// ValidatePolicyServicePrincipalCreationPolicyIdIncludeID checks that 'input' can be parsed as a Policy Service Principal Creation Policy Id Include ID
func ValidatePolicyServicePrincipalCreationPolicyIdIncludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyServicePrincipalCreationPolicyIdIncludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Service Principal Creation Policy Id Include ID
func (id PolicyServicePrincipalCreationPolicyIdIncludeId) ID() string {
	fmtString := "/policies/servicePrincipalCreationPolicies/%s/includes/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalCreationPolicyId, id.ServicePrincipalCreationConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Service Principal Creation Policy Id Include ID
func (id PolicyServicePrincipalCreationPolicyIdIncludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationPolicyId", "servicePrincipalCreationPolicyId"),
		resourceids.StaticSegment("includes", "includes", "includes"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationConditionSetId", "servicePrincipalCreationConditionSetId"),
	}
}

// String returns a human-readable description of this Policy Service Principal Creation Policy Id Include ID
func (id PolicyServicePrincipalCreationPolicyIdIncludeId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal Creation Policy: %q", id.ServicePrincipalCreationPolicyId),
		fmt.Sprintf("Service Principal Creation Condition Set: %q", id.ServicePrincipalCreationConditionSetId),
	}
	return fmt.Sprintf("Policy Service Principal Creation Policy Id Include (%s)", strings.Join(components, "\n"))
}
