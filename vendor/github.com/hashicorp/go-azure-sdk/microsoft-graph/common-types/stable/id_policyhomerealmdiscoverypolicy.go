package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyHomeRealmDiscoveryPolicyId{}

// PolicyHomeRealmDiscoveryPolicyId is a struct representing the Resource ID for a Policy Home Realm Discovery Policy
type PolicyHomeRealmDiscoveryPolicyId struct {
	HomeRealmDiscoveryPolicyId string
}

// NewPolicyHomeRealmDiscoveryPolicyID returns a new PolicyHomeRealmDiscoveryPolicyId struct
func NewPolicyHomeRealmDiscoveryPolicyID(homeRealmDiscoveryPolicyId string) PolicyHomeRealmDiscoveryPolicyId {
	return PolicyHomeRealmDiscoveryPolicyId{
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// ParsePolicyHomeRealmDiscoveryPolicyID parses 'input' into a PolicyHomeRealmDiscoveryPolicyId
func ParsePolicyHomeRealmDiscoveryPolicyID(input string) (*PolicyHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyHomeRealmDiscoveryPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyHomeRealmDiscoveryPolicyIDInsensitively parses 'input' case-insensitively into a PolicyHomeRealmDiscoveryPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyHomeRealmDiscoveryPolicyIDInsensitively(input string) (*PolicyHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyHomeRealmDiscoveryPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyHomeRealmDiscoveryPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HomeRealmDiscoveryPolicyId, ok = input.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", input)
	}

	return nil
}

// ValidatePolicyHomeRealmDiscoveryPolicyID checks that 'input' can be parsed as a Policy Home Realm Discovery Policy ID
func ValidatePolicyHomeRealmDiscoveryPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyHomeRealmDiscoveryPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Home Realm Discovery Policy ID
func (id PolicyHomeRealmDiscoveryPolicyId) ID() string {
	fmtString := "/policies/homeRealmDiscoveryPolicies/%s"
	return fmt.Sprintf(fmtString, id.HomeRealmDiscoveryPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Home Realm Discovery Policy ID
func (id PolicyHomeRealmDiscoveryPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyId"),
	}
}

// String returns a human-readable description of this Policy Home Realm Discovery Policy ID
func (id PolicyHomeRealmDiscoveryPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
	}
	return fmt.Sprintf("Policy Home Realm Discovery Policy (%s)", strings.Join(components, "\n"))
}
