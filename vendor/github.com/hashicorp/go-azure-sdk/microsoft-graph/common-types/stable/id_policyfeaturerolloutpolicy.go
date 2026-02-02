package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyFeatureRolloutPolicyId{}

// PolicyFeatureRolloutPolicyId is a struct representing the Resource ID for a Policy Feature Rollout Policy
type PolicyFeatureRolloutPolicyId struct {
	FeatureRolloutPolicyId string
}

// NewPolicyFeatureRolloutPolicyID returns a new PolicyFeatureRolloutPolicyId struct
func NewPolicyFeatureRolloutPolicyID(featureRolloutPolicyId string) PolicyFeatureRolloutPolicyId {
	return PolicyFeatureRolloutPolicyId{
		FeatureRolloutPolicyId: featureRolloutPolicyId,
	}
}

// ParsePolicyFeatureRolloutPolicyID parses 'input' into a PolicyFeatureRolloutPolicyId
func ParsePolicyFeatureRolloutPolicyID(input string) (*PolicyFeatureRolloutPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyFeatureRolloutPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyFeatureRolloutPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyFeatureRolloutPolicyIDInsensitively parses 'input' case-insensitively into a PolicyFeatureRolloutPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyFeatureRolloutPolicyIDInsensitively(input string) (*PolicyFeatureRolloutPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyFeatureRolloutPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyFeatureRolloutPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyFeatureRolloutPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FeatureRolloutPolicyId, ok = input.Parsed["featureRolloutPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", input)
	}

	return nil
}

// ValidatePolicyFeatureRolloutPolicyID checks that 'input' can be parsed as a Policy Feature Rollout Policy ID
func ValidatePolicyFeatureRolloutPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyFeatureRolloutPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Feature Rollout Policy ID
func (id PolicyFeatureRolloutPolicyId) ID() string {
	fmtString := "/policies/featureRolloutPolicies/%s"
	return fmt.Sprintf(fmtString, id.FeatureRolloutPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Feature Rollout Policy ID
func (id PolicyFeatureRolloutPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("featureRolloutPolicies", "featureRolloutPolicies", "featureRolloutPolicies"),
		resourceids.UserSpecifiedSegment("featureRolloutPolicyId", "featureRolloutPolicyId"),
	}
}

// String returns a human-readable description of this Policy Feature Rollout Policy ID
func (id PolicyFeatureRolloutPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Feature Rollout Policy: %q", id.FeatureRolloutPolicyId),
	}
	return fmt.Sprintf("Policy Feature Rollout Policy (%s)", strings.Join(components, "\n"))
}
