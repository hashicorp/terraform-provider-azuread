package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdAppManagementPolicyId{}

// ServicePrincipalIdAppManagementPolicyId is a struct representing the Resource ID for a Service Principal Id App Management Policy
type ServicePrincipalIdAppManagementPolicyId struct {
	ServicePrincipalId    string
	AppManagementPolicyId string
}

// NewServicePrincipalIdAppManagementPolicyID returns a new ServicePrincipalIdAppManagementPolicyId struct
func NewServicePrincipalIdAppManagementPolicyID(servicePrincipalId string, appManagementPolicyId string) ServicePrincipalIdAppManagementPolicyId {
	return ServicePrincipalIdAppManagementPolicyId{
		ServicePrincipalId:    servicePrincipalId,
		AppManagementPolicyId: appManagementPolicyId,
	}
}

// ParseServicePrincipalIdAppManagementPolicyID parses 'input' into a ServicePrincipalIdAppManagementPolicyId
func ParseServicePrincipalIdAppManagementPolicyID(input string) (*ServicePrincipalIdAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdAppManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdAppManagementPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdAppManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdAppManagementPolicyIDInsensitively(input string) (*ServicePrincipalIdAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdAppManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdAppManagementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.AppManagementPolicyId, ok = input.Parsed["appManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", input)
	}

	return nil
}

// ValidateServicePrincipalIdAppManagementPolicyID checks that 'input' can be parsed as a Service Principal Id App Management Policy ID
func ValidateServicePrincipalIdAppManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdAppManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id App Management Policy ID
func (id ServicePrincipalIdAppManagementPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/appManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.AppManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id App Management Policy ID
func (id ServicePrincipalIdAppManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("appManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyId"),
	}
}

// String returns a human-readable description of this Service Principal Id App Management Policy ID
func (id ServicePrincipalIdAppManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
	}
	return fmt.Sprintf("Service Principal Id App Management Policy (%s)", strings.Join(components, "\n"))
}
