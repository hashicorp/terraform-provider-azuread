package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdHomeRealmDiscoveryPolicyId{}

// ServicePrincipalIdHomeRealmDiscoveryPolicyId is a struct representing the Resource ID for a Service Principal Id Home Realm Discovery Policy
type ServicePrincipalIdHomeRealmDiscoveryPolicyId struct {
	ServicePrincipalId         string
	HomeRealmDiscoveryPolicyId string
}

// NewServicePrincipalIdHomeRealmDiscoveryPolicyID returns a new ServicePrincipalIdHomeRealmDiscoveryPolicyId struct
func NewServicePrincipalIdHomeRealmDiscoveryPolicyID(servicePrincipalId string, homeRealmDiscoveryPolicyId string) ServicePrincipalIdHomeRealmDiscoveryPolicyId {
	return ServicePrincipalIdHomeRealmDiscoveryPolicyId{
		ServicePrincipalId:         servicePrincipalId,
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// ParseServicePrincipalIdHomeRealmDiscoveryPolicyID parses 'input' into a ServicePrincipalIdHomeRealmDiscoveryPolicyId
func ParseServicePrincipalIdHomeRealmDiscoveryPolicyID(input string) (*ServicePrincipalIdHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdHomeRealmDiscoveryPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdHomeRealmDiscoveryPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdHomeRealmDiscoveryPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdHomeRealmDiscoveryPolicyIDInsensitively(input string) (*ServicePrincipalIdHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdHomeRealmDiscoveryPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdHomeRealmDiscoveryPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.HomeRealmDiscoveryPolicyId, ok = input.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", input)
	}

	return nil
}

// ValidateServicePrincipalIdHomeRealmDiscoveryPolicyID checks that 'input' can be parsed as a Service Principal Id Home Realm Discovery Policy ID
func ValidateServicePrincipalIdHomeRealmDiscoveryPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdHomeRealmDiscoveryPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Home Realm Discovery Policy ID
func (id ServicePrincipalIdHomeRealmDiscoveryPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/homeRealmDiscoveryPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.HomeRealmDiscoveryPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Home Realm Discovery Policy ID
func (id ServicePrincipalIdHomeRealmDiscoveryPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyId"),
	}
}

// String returns a human-readable description of this Service Principal Id Home Realm Discovery Policy ID
func (id ServicePrincipalIdHomeRealmDiscoveryPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
	}
	return fmt.Sprintf("Service Principal Id Home Realm Discovery Policy (%s)", strings.Join(components, "\n"))
}
