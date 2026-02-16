package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPreApprovalPolicyId{}

// PolicyPermissionGrantPreApprovalPolicyId is a struct representing the Resource ID for a Policy Permission Grant Pre Approval Policy
type PolicyPermissionGrantPreApprovalPolicyId struct {
	PermissionGrantPreApprovalPolicyId string
}

// NewPolicyPermissionGrantPreApprovalPolicyID returns a new PolicyPermissionGrantPreApprovalPolicyId struct
func NewPolicyPermissionGrantPreApprovalPolicyID(permissionGrantPreApprovalPolicyId string) PolicyPermissionGrantPreApprovalPolicyId {
	return PolicyPermissionGrantPreApprovalPolicyId{
		PermissionGrantPreApprovalPolicyId: permissionGrantPreApprovalPolicyId,
	}
}

// ParsePolicyPermissionGrantPreApprovalPolicyID parses 'input' into a PolicyPermissionGrantPreApprovalPolicyId
func ParsePolicyPermissionGrantPreApprovalPolicyID(input string) (*PolicyPermissionGrantPreApprovalPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPreApprovalPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPreApprovalPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyPermissionGrantPreApprovalPolicyIDInsensitively parses 'input' case-insensitively into a PolicyPermissionGrantPreApprovalPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyPermissionGrantPreApprovalPolicyIDInsensitively(input string) (*PolicyPermissionGrantPreApprovalPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPreApprovalPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPreApprovalPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyPermissionGrantPreApprovalPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionGrantPreApprovalPolicyId, ok = input.Parsed["permissionGrantPreApprovalPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPreApprovalPolicyId", input)
	}

	return nil
}

// ValidatePolicyPermissionGrantPreApprovalPolicyID checks that 'input' can be parsed as a Policy Permission Grant Pre Approval Policy ID
func ValidatePolicyPermissionGrantPreApprovalPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyPermissionGrantPreApprovalPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Permission Grant Pre Approval Policy ID
func (id PolicyPermissionGrantPreApprovalPolicyId) ID() string {
	fmtString := "/policies/permissionGrantPreApprovalPolicies/%s"
	return fmt.Sprintf(fmtString, id.PermissionGrantPreApprovalPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Permission Grant Pre Approval Policy ID
func (id PolicyPermissionGrantPreApprovalPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("permissionGrantPreApprovalPolicies", "permissionGrantPreApprovalPolicies", "permissionGrantPreApprovalPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPreApprovalPolicyId", "permissionGrantPreApprovalPolicyId"),
	}
}

// String returns a human-readable description of this Policy Permission Grant Pre Approval Policy ID
func (id PolicyPermissionGrantPreApprovalPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Permission Grant Pre Approval Policy: %q", id.PermissionGrantPreApprovalPolicyId),
	}
	return fmt.Sprintf("Policy Permission Grant Pre Approval Policy (%s)", strings.Join(components, "\n"))
}
