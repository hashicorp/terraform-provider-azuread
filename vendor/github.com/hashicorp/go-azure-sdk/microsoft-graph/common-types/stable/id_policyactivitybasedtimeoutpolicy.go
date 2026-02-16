package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyActivityBasedTimeoutPolicyId{}

// PolicyActivityBasedTimeoutPolicyId is a struct representing the Resource ID for a Policy Activity Based Timeout Policy
type PolicyActivityBasedTimeoutPolicyId struct {
	ActivityBasedTimeoutPolicyId string
}

// NewPolicyActivityBasedTimeoutPolicyID returns a new PolicyActivityBasedTimeoutPolicyId struct
func NewPolicyActivityBasedTimeoutPolicyID(activityBasedTimeoutPolicyId string) PolicyActivityBasedTimeoutPolicyId {
	return PolicyActivityBasedTimeoutPolicyId{
		ActivityBasedTimeoutPolicyId: activityBasedTimeoutPolicyId,
	}
}

// ParsePolicyActivityBasedTimeoutPolicyID parses 'input' into a PolicyActivityBasedTimeoutPolicyId
func ParsePolicyActivityBasedTimeoutPolicyID(input string) (*PolicyActivityBasedTimeoutPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyActivityBasedTimeoutPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyActivityBasedTimeoutPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyActivityBasedTimeoutPolicyIDInsensitively parses 'input' case-insensitively into a PolicyActivityBasedTimeoutPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyActivityBasedTimeoutPolicyIDInsensitively(input string) (*PolicyActivityBasedTimeoutPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyActivityBasedTimeoutPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyActivityBasedTimeoutPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyActivityBasedTimeoutPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActivityBasedTimeoutPolicyId, ok = input.Parsed["activityBasedTimeoutPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityBasedTimeoutPolicyId", input)
	}

	return nil
}

// ValidatePolicyActivityBasedTimeoutPolicyID checks that 'input' can be parsed as a Policy Activity Based Timeout Policy ID
func ValidatePolicyActivityBasedTimeoutPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyActivityBasedTimeoutPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Activity Based Timeout Policy ID
func (id PolicyActivityBasedTimeoutPolicyId) ID() string {
	fmtString := "/policies/activityBasedTimeoutPolicies/%s"
	return fmt.Sprintf(fmtString, id.ActivityBasedTimeoutPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Activity Based Timeout Policy ID
func (id PolicyActivityBasedTimeoutPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("activityBasedTimeoutPolicies", "activityBasedTimeoutPolicies", "activityBasedTimeoutPolicies"),
		resourceids.UserSpecifiedSegment("activityBasedTimeoutPolicyId", "activityBasedTimeoutPolicyId"),
	}
}

// String returns a human-readable description of this Policy Activity Based Timeout Policy ID
func (id PolicyActivityBasedTimeoutPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Activity Based Timeout Policy: %q", id.ActivityBasedTimeoutPolicyId),
	}
	return fmt.Sprintf("Policy Activity Based Timeout Policy (%s)", strings.Join(components, "\n"))
}
