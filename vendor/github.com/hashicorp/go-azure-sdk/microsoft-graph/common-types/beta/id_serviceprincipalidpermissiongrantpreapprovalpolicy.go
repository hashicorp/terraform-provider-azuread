package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdPermissionGrantPreApprovalPolicyId{}

// ServicePrincipalIdPermissionGrantPreApprovalPolicyId is a struct representing the Resource ID for a Service Principal Id Permission Grant Pre Approval Policy
type ServicePrincipalIdPermissionGrantPreApprovalPolicyId struct {
	ServicePrincipalId                 string
	PermissionGrantPreApprovalPolicyId string
}

// NewServicePrincipalIdPermissionGrantPreApprovalPolicyID returns a new ServicePrincipalIdPermissionGrantPreApprovalPolicyId struct
func NewServicePrincipalIdPermissionGrantPreApprovalPolicyID(servicePrincipalId string, permissionGrantPreApprovalPolicyId string) ServicePrincipalIdPermissionGrantPreApprovalPolicyId {
	return ServicePrincipalIdPermissionGrantPreApprovalPolicyId{
		ServicePrincipalId:                 servicePrincipalId,
		PermissionGrantPreApprovalPolicyId: permissionGrantPreApprovalPolicyId,
	}
}

// ParseServicePrincipalIdPermissionGrantPreApprovalPolicyID parses 'input' into a ServicePrincipalIdPermissionGrantPreApprovalPolicyId
func ParseServicePrincipalIdPermissionGrantPreApprovalPolicyID(input string) (*ServicePrincipalIdPermissionGrantPreApprovalPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdPermissionGrantPreApprovalPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdPermissionGrantPreApprovalPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdPermissionGrantPreApprovalPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdPermissionGrantPreApprovalPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdPermissionGrantPreApprovalPolicyIDInsensitively(input string) (*ServicePrincipalIdPermissionGrantPreApprovalPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdPermissionGrantPreApprovalPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdPermissionGrantPreApprovalPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdPermissionGrantPreApprovalPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.PermissionGrantPreApprovalPolicyId, ok = input.Parsed["permissionGrantPreApprovalPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPreApprovalPolicyId", input)
	}

	return nil
}

// ValidateServicePrincipalIdPermissionGrantPreApprovalPolicyID checks that 'input' can be parsed as a Service Principal Id Permission Grant Pre Approval Policy ID
func ValidateServicePrincipalIdPermissionGrantPreApprovalPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdPermissionGrantPreApprovalPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Permission Grant Pre Approval Policy ID
func (id ServicePrincipalIdPermissionGrantPreApprovalPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/permissionGrantPreApprovalPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.PermissionGrantPreApprovalPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Permission Grant Pre Approval Policy ID
func (id ServicePrincipalIdPermissionGrantPreApprovalPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("permissionGrantPreApprovalPolicies", "permissionGrantPreApprovalPolicies", "permissionGrantPreApprovalPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPreApprovalPolicyId", "permissionGrantPreApprovalPolicyId"),
	}
}

// String returns a human-readable description of this Service Principal Id Permission Grant Pre Approval Policy ID
func (id ServicePrincipalIdPermissionGrantPreApprovalPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Permission Grant Pre Approval Policy: %q", id.PermissionGrantPreApprovalPolicyId),
	}
	return fmt.Sprintf("Service Principal Id Permission Grant Pre Approval Policy (%s)", strings.Join(components, "\n"))
}
