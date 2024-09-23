package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdLicenseDetailId{}

// ServicePrincipalIdLicenseDetailId is a struct representing the Resource ID for a Service Principal Id License Detail
type ServicePrincipalIdLicenseDetailId struct {
	ServicePrincipalId string
	LicenseDetailsId   string
}

// NewServicePrincipalIdLicenseDetailID returns a new ServicePrincipalIdLicenseDetailId struct
func NewServicePrincipalIdLicenseDetailID(servicePrincipalId string, licenseDetailsId string) ServicePrincipalIdLicenseDetailId {
	return ServicePrincipalIdLicenseDetailId{
		ServicePrincipalId: servicePrincipalId,
		LicenseDetailsId:   licenseDetailsId,
	}
}

// ParseServicePrincipalIdLicenseDetailID parses 'input' into a ServicePrincipalIdLicenseDetailId
func ParseServicePrincipalIdLicenseDetailID(input string) (*ServicePrincipalIdLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdLicenseDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdLicenseDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdLicenseDetailIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdLicenseDetailId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdLicenseDetailIDInsensitively(input string) (*ServicePrincipalIdLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdLicenseDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdLicenseDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdLicenseDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.LicenseDetailsId, ok = input.Parsed["licenseDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", input)
	}

	return nil
}

// ValidateServicePrincipalIdLicenseDetailID checks that 'input' can be parsed as a Service Principal Id License Detail ID
func ValidateServicePrincipalIdLicenseDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdLicenseDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id License Detail ID
func (id ServicePrincipalIdLicenseDetailId) ID() string {
	fmtString := "/servicePrincipals/%s/licenseDetails/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.LicenseDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id License Detail ID
func (id ServicePrincipalIdLicenseDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("licenseDetails", "licenseDetails", "licenseDetails"),
		resourceids.UserSpecifiedSegment("licenseDetailsId", "licenseDetailsId"),
	}
}

// String returns a human-readable description of this Service Principal Id License Detail ID
func (id ServicePrincipalIdLicenseDetailId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("License Details: %q", id.LicenseDetailsId),
	}
	return fmt.Sprintf("Service Principal Id License Detail (%s)", strings.Join(components, "\n"))
}
