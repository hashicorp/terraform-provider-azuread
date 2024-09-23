package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyServicePrincipalCreationPolicyIdExcludeId{}

// PolicyServicePrincipalCreationPolicyIdExcludeId is a struct representing the Resource ID for a Policy Service Principal Creation Policy Id Exclude
type PolicyServicePrincipalCreationPolicyIdExcludeId struct {
	ServicePrincipalCreationPolicyId       string
	ServicePrincipalCreationConditionSetId string
}

// NewPolicyServicePrincipalCreationPolicyIdExcludeID returns a new PolicyServicePrincipalCreationPolicyIdExcludeId struct
func NewPolicyServicePrincipalCreationPolicyIdExcludeID(servicePrincipalCreationPolicyId string, servicePrincipalCreationConditionSetId string) PolicyServicePrincipalCreationPolicyIdExcludeId {
	return PolicyServicePrincipalCreationPolicyIdExcludeId{
		ServicePrincipalCreationPolicyId:       servicePrincipalCreationPolicyId,
		ServicePrincipalCreationConditionSetId: servicePrincipalCreationConditionSetId,
	}
}

// ParsePolicyServicePrincipalCreationPolicyIdExcludeID parses 'input' into a PolicyServicePrincipalCreationPolicyIdExcludeId
func ParsePolicyServicePrincipalCreationPolicyIdExcludeID(input string) (*PolicyServicePrincipalCreationPolicyIdExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyServicePrincipalCreationPolicyIdExcludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyServicePrincipalCreationPolicyIdExcludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyServicePrincipalCreationPolicyIdExcludeIDInsensitively parses 'input' case-insensitively into a PolicyServicePrincipalCreationPolicyIdExcludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyServicePrincipalCreationPolicyIdExcludeIDInsensitively(input string) (*PolicyServicePrincipalCreationPolicyIdExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyServicePrincipalCreationPolicyIdExcludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyServicePrincipalCreationPolicyIdExcludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyServicePrincipalCreationPolicyIdExcludeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalCreationPolicyId, ok = input.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", input)
	}

	if id.ServicePrincipalCreationConditionSetId, ok = input.Parsed["servicePrincipalCreationConditionSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationConditionSetId", input)
	}

	return nil
}

// ValidatePolicyServicePrincipalCreationPolicyIdExcludeID checks that 'input' can be parsed as a Policy Service Principal Creation Policy Id Exclude ID
func ValidatePolicyServicePrincipalCreationPolicyIdExcludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyServicePrincipalCreationPolicyIdExcludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Service Principal Creation Policy Id Exclude ID
func (id PolicyServicePrincipalCreationPolicyIdExcludeId) ID() string {
	fmtString := "/policies/servicePrincipalCreationPolicies/%s/excludes/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalCreationPolicyId, id.ServicePrincipalCreationConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Service Principal Creation Policy Id Exclude ID
func (id PolicyServicePrincipalCreationPolicyIdExcludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationPolicyId", "servicePrincipalCreationPolicyId"),
		resourceids.StaticSegment("excludes", "excludes", "excludes"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationConditionSetId", "servicePrincipalCreationConditionSetId"),
	}
}

// String returns a human-readable description of this Policy Service Principal Creation Policy Id Exclude ID
func (id PolicyServicePrincipalCreationPolicyIdExcludeId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal Creation Policy: %q", id.ServicePrincipalCreationPolicyId),
		fmt.Sprintf("Service Principal Creation Condition Set: %q", id.ServicePrincipalCreationConditionSetId),
	}
	return fmt.Sprintf("Policy Service Principal Creation Policy Id Exclude (%s)", strings.Join(components, "\n"))
}
