package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryFeatureRolloutPolicyId{}

// DirectoryFeatureRolloutPolicyId is a struct representing the Resource ID for a Directory Feature Rollout Policy
type DirectoryFeatureRolloutPolicyId struct {
	FeatureRolloutPolicyId string
}

// NewDirectoryFeatureRolloutPolicyID returns a new DirectoryFeatureRolloutPolicyId struct
func NewDirectoryFeatureRolloutPolicyID(featureRolloutPolicyId string) DirectoryFeatureRolloutPolicyId {
	return DirectoryFeatureRolloutPolicyId{
		FeatureRolloutPolicyId: featureRolloutPolicyId,
	}
}

// ParseDirectoryFeatureRolloutPolicyID parses 'input' into a DirectoryFeatureRolloutPolicyId
func ParseDirectoryFeatureRolloutPolicyID(input string) (*DirectoryFeatureRolloutPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryFeatureRolloutPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryFeatureRolloutPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryFeatureRolloutPolicyIDInsensitively parses 'input' case-insensitively into a DirectoryFeatureRolloutPolicyId
// note: this method should only be used for API response data and not user input
func ParseDirectoryFeatureRolloutPolicyIDInsensitively(input string) (*DirectoryFeatureRolloutPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryFeatureRolloutPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryFeatureRolloutPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryFeatureRolloutPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FeatureRolloutPolicyId, ok = input.Parsed["featureRolloutPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", input)
	}

	return nil
}

// ValidateDirectoryFeatureRolloutPolicyID checks that 'input' can be parsed as a Directory Feature Rollout Policy ID
func ValidateDirectoryFeatureRolloutPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryFeatureRolloutPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Feature Rollout Policy ID
func (id DirectoryFeatureRolloutPolicyId) ID() string {
	fmtString := "/directory/featureRolloutPolicies/%s"
	return fmt.Sprintf(fmtString, id.FeatureRolloutPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Feature Rollout Policy ID
func (id DirectoryFeatureRolloutPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("featureRolloutPolicies", "featureRolloutPolicies", "featureRolloutPolicies"),
		resourceids.UserSpecifiedSegment("featureRolloutPolicyId", "featureRolloutPolicyId"),
	}
}

// String returns a human-readable description of this Directory Feature Rollout Policy ID
func (id DirectoryFeatureRolloutPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Feature Rollout Policy: %q", id.FeatureRolloutPolicyId),
	}
	return fmt.Sprintf("Directory Feature Rollout Policy (%s)", strings.Join(components, "\n"))
}
