package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPolicyId{}

// PolicyPermissionGrantPolicyId is a struct representing the Resource ID for a Policy Permission Grant Policy
type PolicyPermissionGrantPolicyId struct {
	PermissionGrantPolicyId string
}

// NewPolicyPermissionGrantPolicyID returns a new PolicyPermissionGrantPolicyId struct
func NewPolicyPermissionGrantPolicyID(permissionGrantPolicyId string) PolicyPermissionGrantPolicyId {
	return PolicyPermissionGrantPolicyId{
		PermissionGrantPolicyId: permissionGrantPolicyId,
	}
}

// ParsePolicyPermissionGrantPolicyID parses 'input' into a PolicyPermissionGrantPolicyId
func ParsePolicyPermissionGrantPolicyID(input string) (*PolicyPermissionGrantPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyPermissionGrantPolicyIDInsensitively parses 'input' case-insensitively into a PolicyPermissionGrantPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyPermissionGrantPolicyIDInsensitively(input string) (*PolicyPermissionGrantPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyPermissionGrantPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionGrantPolicyId, ok = input.Parsed["permissionGrantPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", input)
	}

	return nil
}

// ValidatePolicyPermissionGrantPolicyID checks that 'input' can be parsed as a Policy Permission Grant Policy ID
func ValidatePolicyPermissionGrantPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyPermissionGrantPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Permission Grant Policy ID
func (id PolicyPermissionGrantPolicyId) ID() string {
	fmtString := "/policies/permissionGrantPolicies/%s"
	return fmt.Sprintf(fmtString, id.PermissionGrantPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Permission Grant Policy ID
func (id PolicyPermissionGrantPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("permissionGrantPolicies", "permissionGrantPolicies", "permissionGrantPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPolicyId", "permissionGrantPolicyId"),
	}
}

// String returns a human-readable description of this Policy Permission Grant Policy ID
func (id PolicyPermissionGrantPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Permission Grant Policy: %q", id.PermissionGrantPolicyId),
	}
	return fmt.Sprintf("Policy Permission Grant Policy (%s)", strings.Join(components, "\n"))
}
