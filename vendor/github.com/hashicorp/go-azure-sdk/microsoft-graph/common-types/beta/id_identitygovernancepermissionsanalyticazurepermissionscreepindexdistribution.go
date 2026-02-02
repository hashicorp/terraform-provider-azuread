package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{}

// IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId is a struct representing the Resource ID for a Identity Governance Permissions Analytic Azure Permissions Creep Index Distribution
type IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId struct {
	PermissionsCreepIndexDistributionId string
}

// NewIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID returns a new IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId struct
func NewIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(permissionsCreepIndexDistributionId string) IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId {
	return IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{
		PermissionsCreepIndexDistributionId: permissionsCreepIndexDistributionId,
	}
}

// ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID parses 'input' into a IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId
func ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(input string) (*IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionIDInsensitively(input string) (*IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionsCreepIndexDistributionId, ok = input.Parsed["permissionsCreepIndexDistributionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionsCreepIndexDistributionId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID checks that 'input' can be parsed as a Identity Governance Permissions Analytic Azure Permissions Creep Index Distribution ID
func ValidateIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Analytic Azure Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId) ID() string {
	fmtString := "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/%s"
	return fmt.Sprintf(fmtString, id.PermissionsCreepIndexDistributionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Analytic Azure Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsAnalytics", "permissionsAnalytics", "permissionsAnalytics"),
		resourceids.StaticSegment("azure", "azure", "azure"),
		resourceids.StaticSegment("permissionsCreepIndexDistributions", "permissionsCreepIndexDistributions", "permissionsCreepIndexDistributions"),
		resourceids.UserSpecifiedSegment("permissionsCreepIndexDistributionId", "permissionsCreepIndexDistributionId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Analytic Azure Permissions Creep Index Distribution ID
func (id IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId) String() string {
	components := []string{
		fmt.Sprintf("Permissions Creep Index Distribution: %q", id.PermissionsCreepIndexDistributionId),
	}
	return fmt.Sprintf("Identity Governance Permissions Analytic Azure Permissions Creep Index Distribution (%s)", strings.Join(components, "\n"))
}
