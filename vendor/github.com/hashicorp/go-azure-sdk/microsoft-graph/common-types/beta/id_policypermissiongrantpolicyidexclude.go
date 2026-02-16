package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPolicyIdExcludeId{}

// PolicyPermissionGrantPolicyIdExcludeId is a struct representing the Resource ID for a Policy Permission Grant Policy Id Exclude
type PolicyPermissionGrantPolicyIdExcludeId struct {
	PermissionGrantPolicyId       string
	PermissionGrantConditionSetId string
}

// NewPolicyPermissionGrantPolicyIdExcludeID returns a new PolicyPermissionGrantPolicyIdExcludeId struct
func NewPolicyPermissionGrantPolicyIdExcludeID(permissionGrantPolicyId string, permissionGrantConditionSetId string) PolicyPermissionGrantPolicyIdExcludeId {
	return PolicyPermissionGrantPolicyIdExcludeId{
		PermissionGrantPolicyId:       permissionGrantPolicyId,
		PermissionGrantConditionSetId: permissionGrantConditionSetId,
	}
}

// ParsePolicyPermissionGrantPolicyIdExcludeID parses 'input' into a PolicyPermissionGrantPolicyIdExcludeId
func ParsePolicyPermissionGrantPolicyIdExcludeID(input string) (*PolicyPermissionGrantPolicyIdExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPolicyIdExcludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPolicyIdExcludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyPermissionGrantPolicyIdExcludeIDInsensitively parses 'input' case-insensitively into a PolicyPermissionGrantPolicyIdExcludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyPermissionGrantPolicyIdExcludeIDInsensitively(input string) (*PolicyPermissionGrantPolicyIdExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyPermissionGrantPolicyIdExcludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyPermissionGrantPolicyIdExcludeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyPermissionGrantPolicyIdExcludeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionGrantPolicyId, ok = input.Parsed["permissionGrantPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", input)
	}

	if id.PermissionGrantConditionSetId, ok = input.Parsed["permissionGrantConditionSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantConditionSetId", input)
	}

	return nil
}

// ValidatePolicyPermissionGrantPolicyIdExcludeID checks that 'input' can be parsed as a Policy Permission Grant Policy Id Exclude ID
func ValidatePolicyPermissionGrantPolicyIdExcludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyPermissionGrantPolicyIdExcludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Permission Grant Policy Id Exclude ID
func (id PolicyPermissionGrantPolicyIdExcludeId) ID() string {
	fmtString := "/policies/permissionGrantPolicies/%s/excludes/%s"
	return fmt.Sprintf(fmtString, id.PermissionGrantPolicyId, id.PermissionGrantConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Permission Grant Policy Id Exclude ID
func (id PolicyPermissionGrantPolicyIdExcludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("permissionGrantPolicies", "permissionGrantPolicies", "permissionGrantPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPolicyId", "permissionGrantPolicyId"),
		resourceids.StaticSegment("excludes", "excludes", "excludes"),
		resourceids.UserSpecifiedSegment("permissionGrantConditionSetId", "permissionGrantConditionSetId"),
	}
}

// String returns a human-readable description of this Policy Permission Grant Policy Id Exclude ID
func (id PolicyPermissionGrantPolicyIdExcludeId) String() string {
	components := []string{
		fmt.Sprintf("Permission Grant Policy: %q", id.PermissionGrantPolicyId),
		fmt.Sprintf("Permission Grant Condition Set: %q", id.PermissionGrantConditionSetId),
	}
	return fmt.Sprintf("Policy Permission Grant Policy Id Exclude (%s)", strings.Join(components, "\n"))
}
