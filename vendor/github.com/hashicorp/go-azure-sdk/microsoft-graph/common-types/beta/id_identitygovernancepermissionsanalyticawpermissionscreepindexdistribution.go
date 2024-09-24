package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{}

// IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId is a struct representing the Resource ID for a Identity Governance Permissions Analytic Aw Permissions Creep Index Distribution
type IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId struct {
	PermissionsCreepIndexDistributionId string
}

// NewIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID returns a new IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId struct
func NewIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(permissionsCreepIndexDistributionId string) IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId {
	return IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{
		PermissionsCreepIndexDistributionId: permissionsCreepIndexDistributionId,
	}
}

// ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID parses 'input' into a IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId
func ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(input string) (*IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionIDInsensitively(input string) (*IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionsCreepIndexDistributionId, ok = input.Parsed["permissionsCreepIndexDistributionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionsCreepIndexDistributionId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID checks that 'input' can be parsed as a Identity Governance Permissions Analytic Aw Permissions Creep Index Distribution ID
func ValidateIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Analytic Aw Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId) ID() string {
	fmtString := "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/%s"
	return fmt.Sprintf(fmtString, id.PermissionsCreepIndexDistributionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Analytic Aw Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsAnalytics", "permissionsAnalytics", "permissionsAnalytics"),
		resourceids.StaticSegment("aws", "aws", "aws"),
		resourceids.StaticSegment("permissionsCreepIndexDistributions", "permissionsCreepIndexDistributions", "permissionsCreepIndexDistributions"),
		resourceids.UserSpecifiedSegment("permissionsCreepIndexDistributionId", "permissionsCreepIndexDistributionId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Analytic Aw Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId) String() string {
	components := []string{
		fmt.Sprintf("Permissions Creep Index Distribution: %q", id.PermissionsCreepIndexDistributionId),
	}
	return fmt.Sprintf("Identity Governance Permissions Analytic Aw Permissions Creep Index Distribution (%s)", strings.Join(components, "\n"))
}
