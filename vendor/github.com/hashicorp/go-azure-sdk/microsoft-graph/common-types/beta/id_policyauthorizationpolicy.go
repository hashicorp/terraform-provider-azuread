package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthorizationPolicyId{}

// PolicyAuthorizationPolicyId is a struct representing the Resource ID for a Policy Authorization Policy
type PolicyAuthorizationPolicyId struct {
	AuthorizationPolicyId string
}

// NewPolicyAuthorizationPolicyID returns a new PolicyAuthorizationPolicyId struct
func NewPolicyAuthorizationPolicyID(authorizationPolicyId string) PolicyAuthorizationPolicyId {
	return PolicyAuthorizationPolicyId{
		AuthorizationPolicyId: authorizationPolicyId,
	}
}

// ParsePolicyAuthorizationPolicyID parses 'input' into a PolicyAuthorizationPolicyId
func ParsePolicyAuthorizationPolicyID(input string) (*PolicyAuthorizationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthorizationPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthorizationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAuthorizationPolicyIDInsensitively parses 'input' case-insensitively into a PolicyAuthorizationPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthorizationPolicyIDInsensitively(input string) (*PolicyAuthorizationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthorizationPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthorizationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAuthorizationPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthorizationPolicyId, ok = input.Parsed["authorizationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authorizationPolicyId", input)
	}

	return nil
}

// ValidatePolicyAuthorizationPolicyID checks that 'input' can be parsed as a Policy Authorization Policy ID
func ValidatePolicyAuthorizationPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthorizationPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authorization Policy ID
func (id PolicyAuthorizationPolicyId) ID() string {
	fmtString := "/policies/authorizationPolicy/%s"
	return fmt.Sprintf(fmtString, id.AuthorizationPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authorization Policy ID
func (id PolicyAuthorizationPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("authorizationPolicy", "authorizationPolicy", "authorizationPolicy"),
		resourceids.UserSpecifiedSegment("authorizationPolicyId", "authorizationPolicyId"),
	}
}

// String returns a human-readable description of this Policy Authorization Policy ID
func (id PolicyAuthorizationPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Authorization Policy: %q", id.AuthorizationPolicyId),
	}
	return fmt.Sprintf("Policy Authorization Policy (%s)", strings.Join(components, "\n"))
}
