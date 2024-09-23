package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPolicyIdIncludeId{}

// PolicyPermissionGrantPolicyIdIncludeId is a struct representing the Resource ID for a Policy Permission Grant Policy Id Include
type PolicyPermissionGrantPolicyIdIncludeId struct {
	PermissionGrantPolicyId       string
	PermissionGrantConditionSetId string
}

// NewPolicyPermissionGrantPolicyIdIncludeID returns a new PolicyPermissionGrantPolicyIdIncludeId struct
func NewPolicyPermissionGrantPolicyIdIncludeID(permissionGrantPolicyId string, permissionGrantConditionSetId string) PolicyPermissionGrantPolicyIdIncludeId {
	return PolicyPermissionGrantPolicyIdIncludeId{
		PermissionGrantPolicyId:       permissionGrantPolicyId,
		PermissionGrantConditionSetId: permissionGrantConditionSetId,
	}
}

// ParsePolicyPermissionGrantPolicyIdIncludeID parses 'input' into a PolicyPermissionGrantPolicyIdIncludeId
func ParsePolicyPermissionGrantPolicyIdIncludeID(input string) (*PolicyPermissionGrantPolicyIdIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPolicyIdIncludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPolicyIdIncludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyPermissionGrantPolicyIdIncludeIDInsensitively parses 'input' case-insensitively into a PolicyPermissionGrantPolicyIdIncludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyPermissionGrantPolicyIdIncludeIDInsensitively(input string) (*PolicyPermissionGrantPolicyIdIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPolicyIdIncludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPolicyIdIncludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyPermissionGrantPolicyIdIncludeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionGrantPolicyId, ok = input.Parsed["permissionGrantPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", input)
	}

	if id.PermissionGrantConditionSetId, ok = input.Parsed["permissionGrantConditionSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantConditionSetId", input)
	}

	return nil
}

// ValidatePolicyPermissionGrantPolicyIdIncludeID checks that 'input' can be parsed as a Policy Permission Grant Policy Id Include ID
func ValidatePolicyPermissionGrantPolicyIdIncludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyPermissionGrantPolicyIdIncludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Permission Grant Policy Id Include ID
func (id PolicyPermissionGrantPolicyIdIncludeId) ID() string {
	fmtString := "/policies/permissionGrantPolicies/%s/includes/%s"
	return fmt.Sprintf(fmtString, id.PermissionGrantPolicyId, id.PermissionGrantConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Permission Grant Policy Id Include ID
func (id PolicyPermissionGrantPolicyIdIncludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("permissionGrantPolicies", "permissionGrantPolicies", "permissionGrantPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPolicyId", "permissionGrantPolicyId"),
		resourceids.StaticSegment("includes", "includes", "includes"),
		resourceids.UserSpecifiedSegment("permissionGrantConditionSetId", "permissionGrantConditionSetId"),
	}
}

// String returns a human-readable description of this Policy Permission Grant Policy Id Include ID
func (id PolicyPermissionGrantPolicyIdIncludeId) String() string {
	components := []string{
		fmt.Sprintf("Permission Grant Policy: %q", id.PermissionGrantPolicyId),
		fmt.Sprintf("Permission Grant Condition Set: %q", id.PermissionGrantConditionSetId),
	}
	return fmt.Sprintf("Policy Permission Grant Policy Id Include (%s)", strings.Join(components, "\n"))
}
