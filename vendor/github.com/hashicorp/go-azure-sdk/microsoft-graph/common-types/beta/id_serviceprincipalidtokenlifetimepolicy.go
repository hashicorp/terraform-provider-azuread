package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdTokenLifetimePolicyId{}

// ServicePrincipalIdTokenLifetimePolicyId is a struct representing the Resource ID for a Service Principal Id Token Lifetime Policy
type ServicePrincipalIdTokenLifetimePolicyId struct {
	ServicePrincipalId    string
	TokenLifetimePolicyId string
}

// NewServicePrincipalIdTokenLifetimePolicyID returns a new ServicePrincipalIdTokenLifetimePolicyId struct
func NewServicePrincipalIdTokenLifetimePolicyID(servicePrincipalId string, tokenLifetimePolicyId string) ServicePrincipalIdTokenLifetimePolicyId {
	return ServicePrincipalIdTokenLifetimePolicyId{
		ServicePrincipalId:    servicePrincipalId,
		TokenLifetimePolicyId: tokenLifetimePolicyId,
	}
}

// ParseServicePrincipalIdTokenLifetimePolicyID parses 'input' into a ServicePrincipalIdTokenLifetimePolicyId
func ParseServicePrincipalIdTokenLifetimePolicyID(input string) (*ServicePrincipalIdTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdTokenLifetimePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdTokenLifetimePolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdTokenLifetimePolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdTokenLifetimePolicyIDInsensitively(input string) (*ServicePrincipalIdTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdTokenLifetimePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdTokenLifetimePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.TokenLifetimePolicyId, ok = input.Parsed["tokenLifetimePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", input)
	}

	return nil
}

// ValidateServicePrincipalIdTokenLifetimePolicyID checks that 'input' can be parsed as a Service Principal Id Token Lifetime Policy ID
func ValidateServicePrincipalIdTokenLifetimePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdTokenLifetimePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Token Lifetime Policy ID
func (id ServicePrincipalIdTokenLifetimePolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/tokenLifetimePolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.TokenLifetimePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Token Lifetime Policy ID
func (id ServicePrincipalIdTokenLifetimePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("tokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyId"),
	}
}

// String returns a human-readable description of this Service Principal Id Token Lifetime Policy ID
func (id ServicePrincipalIdTokenLifetimePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
	}
	return fmt.Sprintf("Service Principal Id Token Lifetime Policy (%s)", strings.Join(components, "\n"))
}
