package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAzureFindingId{}

// IdentityGovernancePermissionsAnalyticAzureFindingId is a struct representing the Resource ID for a Identity Governance Permissions Analytic Azure Finding
type IdentityGovernancePermissionsAnalyticAzureFindingId struct {
	FindingId string
}

// NewIdentityGovernancePermissionsAnalyticAzureFindingID returns a new IdentityGovernancePermissionsAnalyticAzureFindingId struct
func NewIdentityGovernancePermissionsAnalyticAzureFindingID(findingId string) IdentityGovernancePermissionsAnalyticAzureFindingId {
	return IdentityGovernancePermissionsAnalyticAzureFindingId{
		FindingId: findingId,
	}
}

// ParseIdentityGovernancePermissionsAnalyticAzureFindingID parses 'input' into a IdentityGovernancePermissionsAnalyticAzureFindingId
func ParseIdentityGovernancePermissionsAnalyticAzureFindingID(input string) (*IdentityGovernancePermissionsAnalyticAzureFindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAzureFindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAzureFindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsAnalyticAzureFindingIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsAnalyticAzureFindingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsAnalyticAzureFindingIDInsensitively(input string) (*IdentityGovernancePermissionsAnalyticAzureFindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAzureFindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAzureFindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsAnalyticAzureFindingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FindingId, ok = input.Parsed["findingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "findingId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsAnalyticAzureFindingID checks that 'input' can be parsed as a Identity Governance Permissions Analytic Azure Finding ID
func ValidateIdentityGovernancePermissionsAnalyticAzureFindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsAnalyticAzureFindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Analytic Azure Finding ID
func (id IdentityGovernancePermissionsAnalyticAzureFindingId) ID() string {
	fmtString := "/identityGovernance/permissionsAnalytics/azure/findings/%s"
	return fmt.Sprintf(fmtString, id.FindingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Analytic Azure Finding ID
func (id IdentityGovernancePermissionsAnalyticAzureFindingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsAnalytics", "permissionsAnalytics", "permissionsAnalytics"),
		resourceids.StaticSegment("azure", "azure", "azure"),
		resourceids.StaticSegment("findings", "findings", "findings"),
		resourceids.UserSpecifiedSegment("findingId", "findingId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Analytic Azure Finding ID
func (id IdentityGovernancePermissionsAnalyticAzureFindingId) String() string {
	components := []string{
		fmt.Sprintf("Finding: %q", id.FindingId),
	}
	return fmt.Sprintf("Identity Governance Permissions Analytic Azure Finding (%s)", strings.Join(components, "\n"))
}
