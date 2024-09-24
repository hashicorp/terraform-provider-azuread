package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdClaimsMappingPolicyId{}

// ServicePrincipalIdClaimsMappingPolicyId is a struct representing the Resource ID for a Service Principal Id Claims Mapping Policy
type ServicePrincipalIdClaimsMappingPolicyId struct {
	ServicePrincipalId    string
	ClaimsMappingPolicyId string
}

// NewServicePrincipalIdClaimsMappingPolicyID returns a new ServicePrincipalIdClaimsMappingPolicyId struct
func NewServicePrincipalIdClaimsMappingPolicyID(servicePrincipalId string, claimsMappingPolicyId string) ServicePrincipalIdClaimsMappingPolicyId {
	return ServicePrincipalIdClaimsMappingPolicyId{
		ServicePrincipalId:    servicePrincipalId,
		ClaimsMappingPolicyId: claimsMappingPolicyId,
	}
}

// ParseServicePrincipalIdClaimsMappingPolicyID parses 'input' into a ServicePrincipalIdClaimsMappingPolicyId
func ParseServicePrincipalIdClaimsMappingPolicyID(input string) (*ServicePrincipalIdClaimsMappingPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdClaimsMappingPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdClaimsMappingPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdClaimsMappingPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdClaimsMappingPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdClaimsMappingPolicyIDInsensitively(input string) (*ServicePrincipalIdClaimsMappingPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdClaimsMappingPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdClaimsMappingPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdClaimsMappingPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.ClaimsMappingPolicyId, ok = input.Parsed["claimsMappingPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", input)
	}

	return nil
}

// ValidateServicePrincipalIdClaimsMappingPolicyID checks that 'input' can be parsed as a Service Principal Id Claims Mapping Policy ID
func ValidateServicePrincipalIdClaimsMappingPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdClaimsMappingPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Claims Mapping Policy ID
func (id ServicePrincipalIdClaimsMappingPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/claimsMappingPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.ClaimsMappingPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Claims Mapping Policy ID
func (id ServicePrincipalIdClaimsMappingPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("claimsMappingPolicies", "claimsMappingPolicies", "claimsMappingPolicies"),
		resourceids.UserSpecifiedSegment("claimsMappingPolicyId", "claimsMappingPolicyId"),
	}
}

// String returns a human-readable description of this Service Principal Id Claims Mapping Policy ID
func (id ServicePrincipalIdClaimsMappingPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Claims Mapping Policy: %q", id.ClaimsMappingPolicyId),
	}
	return fmt.Sprintf("Service Principal Id Claims Mapping Policy (%s)", strings.Join(components, "\n"))
}
