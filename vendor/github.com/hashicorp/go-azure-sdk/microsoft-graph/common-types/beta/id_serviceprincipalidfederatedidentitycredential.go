package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdFederatedIdentityCredentialId{}

// ServicePrincipalIdFederatedIdentityCredentialId is a struct representing the Resource ID for a Service Principal Id Federated Identity Credential
type ServicePrincipalIdFederatedIdentityCredentialId struct {
	ServicePrincipalId            string
	FederatedIdentityCredentialId string
}

// NewServicePrincipalIdFederatedIdentityCredentialID returns a new ServicePrincipalIdFederatedIdentityCredentialId struct
func NewServicePrincipalIdFederatedIdentityCredentialID(servicePrincipalId string, federatedIdentityCredentialId string) ServicePrincipalIdFederatedIdentityCredentialId {
	return ServicePrincipalIdFederatedIdentityCredentialId{
		ServicePrincipalId:            servicePrincipalId,
		FederatedIdentityCredentialId: federatedIdentityCredentialId,
	}
}

// ParseServicePrincipalIdFederatedIdentityCredentialID parses 'input' into a ServicePrincipalIdFederatedIdentityCredentialId
func ParseServicePrincipalIdFederatedIdentityCredentialID(input string) (*ServicePrincipalIdFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdFederatedIdentityCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdFederatedIdentityCredentialIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdFederatedIdentityCredentialId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdFederatedIdentityCredentialIDInsensitively(input string) (*ServicePrincipalIdFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdFederatedIdentityCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdFederatedIdentityCredentialId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.FederatedIdentityCredentialId, ok = input.Parsed["federatedIdentityCredentialId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "federatedIdentityCredentialId", input)
	}

	return nil
}

// ValidateServicePrincipalIdFederatedIdentityCredentialID checks that 'input' can be parsed as a Service Principal Id Federated Identity Credential ID
func ValidateServicePrincipalIdFederatedIdentityCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdFederatedIdentityCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Federated Identity Credential ID
func (id ServicePrincipalIdFederatedIdentityCredentialId) ID() string {
	fmtString := "/servicePrincipals/%s/federatedIdentityCredentials/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.FederatedIdentityCredentialId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Federated Identity Credential ID
func (id ServicePrincipalIdFederatedIdentityCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("federatedIdentityCredentials", "federatedIdentityCredentials", "federatedIdentityCredentials"),
		resourceids.UserSpecifiedSegment("federatedIdentityCredentialId", "federatedIdentityCredentialId"),
	}
}

// String returns a human-readable description of this Service Principal Id Federated Identity Credential ID
func (id ServicePrincipalIdFederatedIdentityCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Federated Identity Credential: %q", id.FederatedIdentityCredentialId),
	}
	return fmt.Sprintf("Service Principal Id Federated Identity Credential (%s)", strings.Join(components, "\n"))
}
