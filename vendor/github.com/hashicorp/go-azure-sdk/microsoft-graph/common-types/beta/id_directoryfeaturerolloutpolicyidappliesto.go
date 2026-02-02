package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryFeatureRolloutPolicyIdAppliesToId{}

// DirectoryFeatureRolloutPolicyIdAppliesToId is a struct representing the Resource ID for a Directory Feature Rollout Policy Id Applies To
type DirectoryFeatureRolloutPolicyIdAppliesToId struct {
	FeatureRolloutPolicyId string
	DirectoryObjectId      string
}

// NewDirectoryFeatureRolloutPolicyIdAppliesToID returns a new DirectoryFeatureRolloutPolicyIdAppliesToId struct
func NewDirectoryFeatureRolloutPolicyIdAppliesToID(featureRolloutPolicyId string, directoryObjectId string) DirectoryFeatureRolloutPolicyIdAppliesToId {
	return DirectoryFeatureRolloutPolicyIdAppliesToId{
		FeatureRolloutPolicyId: featureRolloutPolicyId,
		DirectoryObjectId:      directoryObjectId,
	}
}

// ParseDirectoryFeatureRolloutPolicyIdAppliesToID parses 'input' into a DirectoryFeatureRolloutPolicyIdAppliesToId
func ParseDirectoryFeatureRolloutPolicyIdAppliesToID(input string) (*DirectoryFeatureRolloutPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryFeatureRolloutPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryFeatureRolloutPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryFeatureRolloutPolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a DirectoryFeatureRolloutPolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParseDirectoryFeatureRolloutPolicyIdAppliesToIDInsensitively(input string) (*DirectoryFeatureRolloutPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryFeatureRolloutPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryFeatureRolloutPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryFeatureRolloutPolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.FeatureRolloutPolicyId, ok = input.Parsed["featureRolloutPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryFeatureRolloutPolicyIdAppliesToID checks that 'input' can be parsed as a Directory Feature Rollout Policy Id Applies To ID
func ValidateDirectoryFeatureRolloutPolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryFeatureRolloutPolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Feature Rollout Policy Id Applies To ID
func (id DirectoryFeatureRolloutPolicyIdAppliesToId) ID() string {
	fmtString := "/directory/featureRolloutPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.FeatureRolloutPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Feature Rollout Policy Id Applies To ID
func (id DirectoryFeatureRolloutPolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("featureRolloutPolicies", "featureRolloutPolicies", "featureRolloutPolicies"),
		resourceids.UserSpecifiedSegment("featureRolloutPolicyId", "featureRolloutPolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Feature Rollout Policy Id Applies To ID
func (id DirectoryFeatureRolloutPolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Feature Rollout Policy: %q", id.FeatureRolloutPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Feature Rollout Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
