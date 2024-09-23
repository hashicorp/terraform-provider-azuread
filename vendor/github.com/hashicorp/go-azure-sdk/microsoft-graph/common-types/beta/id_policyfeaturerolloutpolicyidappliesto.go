package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyFeatureRolloutPolicyIdAppliesToId{}

// PolicyFeatureRolloutPolicyIdAppliesToId is a struct representing the Resource ID for a Policy Feature Rollout Policy Id Applies To
type PolicyFeatureRolloutPolicyIdAppliesToId struct {
	FeatureRolloutPolicyId string
	DirectoryObjectId      string
}

// NewPolicyFeatureRolloutPolicyIdAppliesToID returns a new PolicyFeatureRolloutPolicyIdAppliesToId struct
func NewPolicyFeatureRolloutPolicyIdAppliesToID(featureRolloutPolicyId string, directoryObjectId string) PolicyFeatureRolloutPolicyIdAppliesToId {
	return PolicyFeatureRolloutPolicyIdAppliesToId{
		FeatureRolloutPolicyId: featureRolloutPolicyId,
		DirectoryObjectId:      directoryObjectId,
	}
}

// ParsePolicyFeatureRolloutPolicyIdAppliesToID parses 'input' into a PolicyFeatureRolloutPolicyIdAppliesToId
func ParsePolicyFeatureRolloutPolicyIdAppliesToID(input string) (*PolicyFeatureRolloutPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyFeatureRolloutPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyFeatureRolloutPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyFeatureRolloutPolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyFeatureRolloutPolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyFeatureRolloutPolicyIdAppliesToIDInsensitively(input string) (*PolicyFeatureRolloutPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyFeatureRolloutPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyFeatureRolloutPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyFeatureRolloutPolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FeatureRolloutPolicyId, ok = input.Parsed["featureRolloutPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyFeatureRolloutPolicyIdAppliesToID checks that 'input' can be parsed as a Policy Feature Rollout Policy Id Applies To ID
func ValidatePolicyFeatureRolloutPolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyFeatureRolloutPolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Feature Rollout Policy Id Applies To ID
func (id PolicyFeatureRolloutPolicyIdAppliesToId) ID() string {
	fmtString := "/policies/featureRolloutPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.FeatureRolloutPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Feature Rollout Policy Id Applies To ID
func (id PolicyFeatureRolloutPolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("featureRolloutPolicies", "featureRolloutPolicies", "featureRolloutPolicies"),
		resourceids.UserSpecifiedSegment("featureRolloutPolicyId", "featureRolloutPolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy Feature Rollout Policy Id Applies To ID
func (id PolicyFeatureRolloutPolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Feature Rollout Policy: %q", id.FeatureRolloutPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Feature Rollout Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
