package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdTokenIssuancePolicyId{}

// ServicePrincipalIdTokenIssuancePolicyId is a struct representing the Resource ID for a Service Principal Id Token Issuance Policy
type ServicePrincipalIdTokenIssuancePolicyId struct {
	ServicePrincipalId    string
	TokenIssuancePolicyId string
}

// NewServicePrincipalIdTokenIssuancePolicyID returns a new ServicePrincipalIdTokenIssuancePolicyId struct
func NewServicePrincipalIdTokenIssuancePolicyID(servicePrincipalId string, tokenIssuancePolicyId string) ServicePrincipalIdTokenIssuancePolicyId {
	return ServicePrincipalIdTokenIssuancePolicyId{
		ServicePrincipalId:    servicePrincipalId,
		TokenIssuancePolicyId: tokenIssuancePolicyId,
	}
}

// ParseServicePrincipalIdTokenIssuancePolicyID parses 'input' into a ServicePrincipalIdTokenIssuancePolicyId
func ParseServicePrincipalIdTokenIssuancePolicyID(input string) (*ServicePrincipalIdTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdTokenIssuancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdTokenIssuancePolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdTokenIssuancePolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdTokenIssuancePolicyIDInsensitively(input string) (*ServicePrincipalIdTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdTokenIssuancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdTokenIssuancePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.TokenIssuancePolicyId, ok = input.Parsed["tokenIssuancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", input)
	}

	return nil
}

// ValidateServicePrincipalIdTokenIssuancePolicyID checks that 'input' can be parsed as a Service Principal Id Token Issuance Policy ID
func ValidateServicePrincipalIdTokenIssuancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdTokenIssuancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Token Issuance Policy ID
func (id ServicePrincipalIdTokenIssuancePolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/tokenIssuancePolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.TokenIssuancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Token Issuance Policy ID
func (id ServicePrincipalIdTokenIssuancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("tokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyId"),
	}
}

// String returns a human-readable description of this Service Principal Id Token Issuance Policy ID
func (id ServicePrincipalIdTokenIssuancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
	}
	return fmt.Sprintf("Service Principal Id Token Issuance Policy (%s)", strings.Join(components, "\n"))
}
