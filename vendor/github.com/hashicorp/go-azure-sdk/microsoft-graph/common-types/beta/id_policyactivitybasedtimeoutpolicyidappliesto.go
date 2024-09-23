package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyActivityBasedTimeoutPolicyIdAppliesToId{}

// PolicyActivityBasedTimeoutPolicyIdAppliesToId is a struct representing the Resource ID for a Policy Activity Based Timeout Policy Id Applies To
type PolicyActivityBasedTimeoutPolicyIdAppliesToId struct {
	ActivityBasedTimeoutPolicyId string
	DirectoryObjectId            string
}

// NewPolicyActivityBasedTimeoutPolicyIdAppliesToID returns a new PolicyActivityBasedTimeoutPolicyIdAppliesToId struct
func NewPolicyActivityBasedTimeoutPolicyIdAppliesToID(activityBasedTimeoutPolicyId string, directoryObjectId string) PolicyActivityBasedTimeoutPolicyIdAppliesToId {
	return PolicyActivityBasedTimeoutPolicyIdAppliesToId{
		ActivityBasedTimeoutPolicyId: activityBasedTimeoutPolicyId,
		DirectoryObjectId:            directoryObjectId,
	}
}

// ParsePolicyActivityBasedTimeoutPolicyIdAppliesToID parses 'input' into a PolicyActivityBasedTimeoutPolicyIdAppliesToId
func ParsePolicyActivityBasedTimeoutPolicyIdAppliesToID(input string) (*PolicyActivityBasedTimeoutPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyActivityBasedTimeoutPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyActivityBasedTimeoutPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyActivityBasedTimeoutPolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyActivityBasedTimeoutPolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyActivityBasedTimeoutPolicyIdAppliesToIDInsensitively(input string) (*PolicyActivityBasedTimeoutPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyActivityBasedTimeoutPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyActivityBasedTimeoutPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyActivityBasedTimeoutPolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActivityBasedTimeoutPolicyId, ok = input.Parsed["activityBasedTimeoutPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityBasedTimeoutPolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyActivityBasedTimeoutPolicyIdAppliesToID checks that 'input' can be parsed as a Policy Activity Based Timeout Policy Id Applies To ID
func ValidatePolicyActivityBasedTimeoutPolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyActivityBasedTimeoutPolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Activity Based Timeout Policy Id Applies To ID
func (id PolicyActivityBasedTimeoutPolicyIdAppliesToId) ID() string {
	fmtString := "/policies/activityBasedTimeoutPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.ActivityBasedTimeoutPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Activity Based Timeout Policy Id Applies To ID
func (id PolicyActivityBasedTimeoutPolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("activityBasedTimeoutPolicies", "activityBasedTimeoutPolicies", "activityBasedTimeoutPolicies"),
		resourceids.UserSpecifiedSegment("activityBasedTimeoutPolicyId", "activityBasedTimeoutPolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy Activity Based Timeout Policy Id Applies To ID
func (id PolicyActivityBasedTimeoutPolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Activity Based Timeout Policy: %q", id.ActivityBasedTimeoutPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Activity Based Timeout Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
