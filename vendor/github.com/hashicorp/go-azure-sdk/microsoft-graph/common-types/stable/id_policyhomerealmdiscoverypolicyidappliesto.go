package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyHomeRealmDiscoveryPolicyIdAppliesToId{}

// PolicyHomeRealmDiscoveryPolicyIdAppliesToId is a struct representing the Resource ID for a Policy Home Realm Discovery Policy Id Applies To
type PolicyHomeRealmDiscoveryPolicyIdAppliesToId struct {
	HomeRealmDiscoveryPolicyId string
	DirectoryObjectId          string
}

// NewPolicyHomeRealmDiscoveryPolicyIdAppliesToID returns a new PolicyHomeRealmDiscoveryPolicyIdAppliesToId struct
func NewPolicyHomeRealmDiscoveryPolicyIdAppliesToID(homeRealmDiscoveryPolicyId string, directoryObjectId string) PolicyHomeRealmDiscoveryPolicyIdAppliesToId {
	return PolicyHomeRealmDiscoveryPolicyIdAppliesToId{
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
		DirectoryObjectId:          directoryObjectId,
	}
}

// ParsePolicyHomeRealmDiscoveryPolicyIdAppliesToID parses 'input' into a PolicyHomeRealmDiscoveryPolicyIdAppliesToId
func ParsePolicyHomeRealmDiscoveryPolicyIdAppliesToID(input string) (*PolicyHomeRealmDiscoveryPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyHomeRealmDiscoveryPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyHomeRealmDiscoveryPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyHomeRealmDiscoveryPolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyHomeRealmDiscoveryPolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyHomeRealmDiscoveryPolicyIdAppliesToIDInsensitively(input string) (*PolicyHomeRealmDiscoveryPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyHomeRealmDiscoveryPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyHomeRealmDiscoveryPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyHomeRealmDiscoveryPolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HomeRealmDiscoveryPolicyId, ok = input.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyHomeRealmDiscoveryPolicyIdAppliesToID checks that 'input' can be parsed as a Policy Home Realm Discovery Policy Id Applies To ID
func ValidatePolicyHomeRealmDiscoveryPolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyHomeRealmDiscoveryPolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Home Realm Discovery Policy Id Applies To ID
func (id PolicyHomeRealmDiscoveryPolicyIdAppliesToId) ID() string {
	fmtString := "/policies/homeRealmDiscoveryPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.HomeRealmDiscoveryPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Home Realm Discovery Policy Id Applies To ID
func (id PolicyHomeRealmDiscoveryPolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy Home Realm Discovery Policy Id Applies To ID
func (id PolicyHomeRealmDiscoveryPolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Home Realm Discovery Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
