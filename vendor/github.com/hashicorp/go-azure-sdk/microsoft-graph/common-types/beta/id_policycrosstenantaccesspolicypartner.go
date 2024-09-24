package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyCrossTenantAccessPolicyPartnerId{}

// PolicyCrossTenantAccessPolicyPartnerId is a struct representing the Resource ID for a Policy Cross Tenant Access Policy Partner
type PolicyCrossTenantAccessPolicyPartnerId struct {
	CrossTenantAccessPolicyConfigurationPartnerTenantId string
}

// NewPolicyCrossTenantAccessPolicyPartnerID returns a new PolicyCrossTenantAccessPolicyPartnerId struct
func NewPolicyCrossTenantAccessPolicyPartnerID(crossTenantAccessPolicyConfigurationPartnerTenantId string) PolicyCrossTenantAccessPolicyPartnerId {
	return PolicyCrossTenantAccessPolicyPartnerId{
		CrossTenantAccessPolicyConfigurationPartnerTenantId: crossTenantAccessPolicyConfigurationPartnerTenantId,
	}
}

// ParsePolicyCrossTenantAccessPolicyPartnerID parses 'input' into a PolicyCrossTenantAccessPolicyPartnerId
func ParsePolicyCrossTenantAccessPolicyPartnerID(input string) (*PolicyCrossTenantAccessPolicyPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyCrossTenantAccessPolicyPartnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyCrossTenantAccessPolicyPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyCrossTenantAccessPolicyPartnerIDInsensitively parses 'input' case-insensitively into a PolicyCrossTenantAccessPolicyPartnerId
// note: this method should only be used for API response data and not user input
func ParsePolicyCrossTenantAccessPolicyPartnerIDInsensitively(input string) (*PolicyCrossTenantAccessPolicyPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyCrossTenantAccessPolicyPartnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyCrossTenantAccessPolicyPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyCrossTenantAccessPolicyPartnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CrossTenantAccessPolicyConfigurationPartnerTenantId, ok = input.Parsed["crossTenantAccessPolicyConfigurationPartnerTenantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "crossTenantAccessPolicyConfigurationPartnerTenantId", input)
	}

	return nil
}

// ValidatePolicyCrossTenantAccessPolicyPartnerID checks that 'input' can be parsed as a Policy Cross Tenant Access Policy Partner ID
func ValidatePolicyCrossTenantAccessPolicyPartnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyCrossTenantAccessPolicyPartnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Cross Tenant Access Policy Partner ID
func (id PolicyCrossTenantAccessPolicyPartnerId) ID() string {
	fmtString := "/policies/crossTenantAccessPolicy/partners/%s"
	return fmt.Sprintf(fmtString, id.CrossTenantAccessPolicyConfigurationPartnerTenantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Cross Tenant Access Policy Partner ID
func (id PolicyCrossTenantAccessPolicyPartnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("crossTenantAccessPolicy", "crossTenantAccessPolicy", "crossTenantAccessPolicy"),
		resourceids.StaticSegment("partners", "partners", "partners"),
		resourceids.UserSpecifiedSegment("crossTenantAccessPolicyConfigurationPartnerTenantId", "crossTenantAccessPolicyConfigurationPartnerTenantId"),
	}
}

// String returns a human-readable description of this Policy Cross Tenant Access Policy Partner ID
func (id PolicyCrossTenantAccessPolicyPartnerId) String() string {
	components := []string{
		fmt.Sprintf("Cross Tenant Access Policy Configuration Partner Tenant: %q", id.CrossTenantAccessPolicyConfigurationPartnerTenantId),
	}
	return fmt.Sprintf("Policy Cross Tenant Access Policy Partner (%s)", strings.Join(components, "\n"))
}
