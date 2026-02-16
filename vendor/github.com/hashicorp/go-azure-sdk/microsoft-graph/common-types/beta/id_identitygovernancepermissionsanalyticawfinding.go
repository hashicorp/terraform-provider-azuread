package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAwFindingId{}

// IdentityGovernancePermissionsAnalyticAwFindingId is a struct representing the Resource ID for a Identity Governance Permissions Analytic Aw Finding
type IdentityGovernancePermissionsAnalyticAwFindingId struct {
	FindingId string
}

// NewIdentityGovernancePermissionsAnalyticAwFindingID returns a new IdentityGovernancePermissionsAnalyticAwFindingId struct
func NewIdentityGovernancePermissionsAnalyticAwFindingID(findingId string) IdentityGovernancePermissionsAnalyticAwFindingId {
	return IdentityGovernancePermissionsAnalyticAwFindingId{
		FindingId: findingId,
	}
}

// ParseIdentityGovernancePermissionsAnalyticAwFindingID parses 'input' into a IdentityGovernancePermissionsAnalyticAwFindingId
func ParseIdentityGovernancePermissionsAnalyticAwFindingID(input string) (*IdentityGovernancePermissionsAnalyticAwFindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAwFindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAwFindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsAnalyticAwFindingIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsAnalyticAwFindingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsAnalyticAwFindingIDInsensitively(input string) (*IdentityGovernancePermissionsAnalyticAwFindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsAnalyticAwFindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsAnalyticAwFindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsAnalyticAwFindingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FindingId, ok = input.Parsed["findingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "findingId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsAnalyticAwFindingID checks that 'input' can be parsed as a Identity Governance Permissions Analytic Aw Finding ID
func ValidateIdentityGovernancePermissionsAnalyticAwFindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsAnalyticAwFindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Analytic Aw Finding ID
func (id IdentityGovernancePermissionsAnalyticAwFindingId) ID() string {
	fmtString := "/identityGovernance/permissionsAnalytics/aws/findings/%s"
	return fmt.Sprintf(fmtString, id.FindingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Analytic Aw Finding ID
func (id IdentityGovernancePermissionsAnalyticAwFindingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsAnalytics", "permissionsAnalytics", "permissionsAnalytics"),
		resourceids.StaticSegment("aws", "aws", "aws"),
		resourceids.StaticSegment("findings", "findings", "findings"),
		resourceids.UserSpecifiedSegment("findingId", "findingId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Analytic Aw Finding ID
func (id IdentityGovernancePermissionsAnalyticAwFindingId) String() string {
	components := []string{
		fmt.Sprintf("Finding: %q", id.FindingId),
	}
	return fmt.Sprintf("Identity Governance Permissions Analytic Aw Finding (%s)", strings.Join(components, "\n"))
}
