package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdDelegatedPermissionClassificationId{}

// ServicePrincipalIdDelegatedPermissionClassificationId is a struct representing the Resource ID for a Service Principal Id Delegated Permission Classification
type ServicePrincipalIdDelegatedPermissionClassificationId struct {
	ServicePrincipalId                  string
	DelegatedPermissionClassificationId string
}

// NewServicePrincipalIdDelegatedPermissionClassificationID returns a new ServicePrincipalIdDelegatedPermissionClassificationId struct
func NewServicePrincipalIdDelegatedPermissionClassificationID(servicePrincipalId string, delegatedPermissionClassificationId string) ServicePrincipalIdDelegatedPermissionClassificationId {
	return ServicePrincipalIdDelegatedPermissionClassificationId{
		ServicePrincipalId:                  servicePrincipalId,
		DelegatedPermissionClassificationId: delegatedPermissionClassificationId,
	}
}

// ParseServicePrincipalIdDelegatedPermissionClassificationID parses 'input' into a ServicePrincipalIdDelegatedPermissionClassificationId
func ParseServicePrincipalIdDelegatedPermissionClassificationID(input string) (*ServicePrincipalIdDelegatedPermissionClassificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdDelegatedPermissionClassificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdDelegatedPermissionClassificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdDelegatedPermissionClassificationIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdDelegatedPermissionClassificationId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdDelegatedPermissionClassificationIDInsensitively(input string) (*ServicePrincipalIdDelegatedPermissionClassificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdDelegatedPermissionClassificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdDelegatedPermissionClassificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdDelegatedPermissionClassificationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.DelegatedPermissionClassificationId, ok = input.Parsed["delegatedPermissionClassificationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "delegatedPermissionClassificationId", input)
	}

	return nil
}

// ValidateServicePrincipalIdDelegatedPermissionClassificationID checks that 'input' can be parsed as a Service Principal Id Delegated Permission Classification ID
func ValidateServicePrincipalIdDelegatedPermissionClassificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdDelegatedPermissionClassificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Delegated Permission Classification ID
func (id ServicePrincipalIdDelegatedPermissionClassificationId) ID() string {
	fmtString := "/servicePrincipals/%s/delegatedPermissionClassifications/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DelegatedPermissionClassificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Delegated Permission Classification ID
func (id ServicePrincipalIdDelegatedPermissionClassificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("delegatedPermissionClassifications", "delegatedPermissionClassifications", "delegatedPermissionClassifications"),
		resourceids.UserSpecifiedSegment("delegatedPermissionClassificationId", "delegatedPermissionClassificationId"),
	}
}

// String returns a human-readable description of this Service Principal Id Delegated Permission Classification ID
func (id ServicePrincipalIdDelegatedPermissionClassificationId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Delegated Permission Classification: %q", id.DelegatedPermissionClassificationId),
	}
	return fmt.Sprintf("Service Principal Id Delegated Permission Classification (%s)", strings.Join(components, "\n"))
}
