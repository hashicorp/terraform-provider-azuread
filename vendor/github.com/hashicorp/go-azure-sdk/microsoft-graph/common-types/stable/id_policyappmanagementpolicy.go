package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAppManagementPolicyId{}

// PolicyAppManagementPolicyId is a struct representing the Resource ID for a Policy App Management Policy
type PolicyAppManagementPolicyId struct {
	AppManagementPolicyId string
}

// NewPolicyAppManagementPolicyID returns a new PolicyAppManagementPolicyId struct
func NewPolicyAppManagementPolicyID(appManagementPolicyId string) PolicyAppManagementPolicyId {
	return PolicyAppManagementPolicyId{
		AppManagementPolicyId: appManagementPolicyId,
	}
}

// ParsePolicyAppManagementPolicyID parses 'input' into a PolicyAppManagementPolicyId
func ParsePolicyAppManagementPolicyID(input string) (*PolicyAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAppManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAppManagementPolicyIDInsensitively parses 'input' case-insensitively into a PolicyAppManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyAppManagementPolicyIDInsensitively(input string) (*PolicyAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAppManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAppManagementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppManagementPolicyId, ok = input.Parsed["appManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", input)
	}

	return nil
}

// ValidatePolicyAppManagementPolicyID checks that 'input' can be parsed as a Policy App Management Policy ID
func ValidatePolicyAppManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAppManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy App Management Policy ID
func (id PolicyAppManagementPolicyId) ID() string {
	fmtString := "/policies/appManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.AppManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy App Management Policy ID
func (id PolicyAppManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("appManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyId"),
	}
}

// String returns a human-readable description of this Policy App Management Policy ID
func (id PolicyAppManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
	}
	return fmt.Sprintf("Policy App Management Policy (%s)", strings.Join(components, "\n"))
}
