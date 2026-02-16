package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticGcpFindingId{}

// IdentityGovernancePermissionsAnalyticGcpFindingId is a struct representing the Resource ID for a Identity Governance Permissions Analytic Gcp Finding
type IdentityGovernancePermissionsAnalyticGcpFindingId struct {
	FindingId string
}

// NewIdentityGovernancePermissionsAnalyticGcpFindingID returns a new IdentityGovernancePermissionsAnalyticGcpFindingId struct
func NewIdentityGovernancePermissionsAnalyticGcpFindingID(findingId string) IdentityGovernancePermissionsAnalyticGcpFindingId {
	return IdentityGovernancePermissionsAnalyticGcpFindingId{
		FindingId: findingId,
	}
}

// ParseIdentityGovernancePermissionsAnalyticGcpFindingID parses 'input' into a IdentityGovernancePermissionsAnalyticGcpFindingId
func ParseIdentityGovernancePermissionsAnalyticGcpFindingID(input string) (*IdentityGovernancePermissionsAnalyticGcpFindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticGcpFindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticGcpFindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsAnalyticGcpFindingIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsAnalyticGcpFindingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsAnalyticGcpFindingIDInsensitively(input string) (*IdentityGovernancePermissionsAnalyticGcpFindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticGcpFindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticGcpFindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsAnalyticGcpFindingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FindingId, ok = input.Parsed["findingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "findingId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsAnalyticGcpFindingID checks that 'input' can be parsed as a Identity Governance Permissions Analytic Gcp Finding ID
func ValidateIdentityGovernancePermissionsAnalyticGcpFindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsAnalyticGcpFindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Analytic Gcp Finding ID
func (id IdentityGovernancePermissionsAnalyticGcpFindingId) ID() string {
	fmtString := "/identityGovernance/permissionsAnalytics/gcp/findings/%s"
	return fmt.Sprintf(fmtString, id.FindingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Analytic Gcp Finding ID
func (id IdentityGovernancePermissionsAnalyticGcpFindingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsAnalytics", "permissionsAnalytics", "permissionsAnalytics"),
		resourceids.StaticSegment("gcp", "gcp", "gcp"),
		resourceids.StaticSegment("findings", "findings", "findings"),
		resourceids.UserSpecifiedSegment("findingId", "findingId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Analytic Gcp Finding ID
func (id IdentityGovernancePermissionsAnalyticGcpFindingId) String() string {
	components := []string{
		fmt.Sprintf("Finding: %q", id.FindingId),
	}
	return fmt.Sprintf("Identity Governance Permissions Analytic Gcp Finding (%s)", strings.Join(components, "\n"))
}
