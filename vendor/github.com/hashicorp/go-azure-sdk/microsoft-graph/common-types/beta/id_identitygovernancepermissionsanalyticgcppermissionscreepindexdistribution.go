package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{}

// IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId is a struct representing the Resource ID for a Identity Governance Permissions Analytic Gcp Permissions Creep Index Distribution
type IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId struct {
	PermissionsCreepIndexDistributionId string
}

// NewIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID returns a new IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId struct
func NewIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(permissionsCreepIndexDistributionId string) IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId {
	return IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{
		PermissionsCreepIndexDistributionId: permissionsCreepIndexDistributionId,
	}
}

// ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID parses 'input' into a IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId
func ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(input string) (*IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionIDInsensitively(input string) (*IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionsCreepIndexDistributionId, ok = input.Parsed["permissionsCreepIndexDistributionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionsCreepIndexDistributionId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID checks that 'input' can be parsed as a Identity Governance Permissions Analytic Gcp Permissions Creep Index Distribution ID
func ValidateIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Analytic Gcp Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId) ID() string {
	fmtString := "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/%s"
	return fmt.Sprintf(fmtString, id.PermissionsCreepIndexDistributionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Analytic Gcp Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsAnalytics", "permissionsAnalytics", "permissionsAnalytics"),
		resourceids.StaticSegment("gcp", "gcp", "gcp"),
		resourceids.StaticSegment("permissionsCreepIndexDistributions", "permissionsCreepIndexDistributions", "permissionsCreepIndexDistributions"),
		resourceids.UserSpecifiedSegment("permissionsCreepIndexDistributionId", "permissionsCreepIndexDistributionId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Analytic Gcp Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId) String() string {
	components := []string{
		fmt.Sprintf("Permissions Creep Index Distribution: %q", id.PermissionsCreepIndexDistributionId),
	}
	return fmt.Sprintf("Identity Governance Permissions Analytic Gcp Permissions Creep Index Distribution (%s)", strings.Join(components, "\n"))
}
