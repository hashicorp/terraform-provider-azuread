package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileAppManagementPolicyIdIncludedGroupId{}

// PolicyMobileAppManagementPolicyIdIncludedGroupId is a struct representing the Resource ID for a Policy Mobile App Management Policy Id Included Group
type PolicyMobileAppManagementPolicyIdIncludedGroupId struct {
	MobilityManagementPolicyId string
	GroupId                    string
}

// NewPolicyMobileAppManagementPolicyIdIncludedGroupID returns a new PolicyMobileAppManagementPolicyIdIncludedGroupId struct
func NewPolicyMobileAppManagementPolicyIdIncludedGroupID(mobilityManagementPolicyId string, groupId string) PolicyMobileAppManagementPolicyIdIncludedGroupId {
	return PolicyMobileAppManagementPolicyIdIncludedGroupId{
		MobilityManagementPolicyId: mobilityManagementPolicyId,
		GroupId:                    groupId,
	}
}

// ParsePolicyMobileAppManagementPolicyIdIncludedGroupID parses 'input' into a PolicyMobileAppManagementPolicyIdIncludedGroupId
func ParsePolicyMobileAppManagementPolicyIdIncludedGroupID(input string) (*PolicyMobileAppManagementPolicyIdIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileAppManagementPolicyIdIncludedGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileAppManagementPolicyIdIncludedGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyMobileAppManagementPolicyIdIncludedGroupIDInsensitively parses 'input' case-insensitively into a PolicyMobileAppManagementPolicyIdIncludedGroupId
// note: this method should only be used for API response data and not user input
func ParsePolicyMobileAppManagementPolicyIdIncludedGroupIDInsensitively(input string) (*PolicyMobileAppManagementPolicyIdIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileAppManagementPolicyIdIncludedGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileAppManagementPolicyIdIncludedGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyMobileAppManagementPolicyIdIncludedGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobilityManagementPolicyId, ok = input.Parsed["mobilityManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", input)
	}

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	return nil
}

// ValidatePolicyMobileAppManagementPolicyIdIncludedGroupID checks that 'input' can be parsed as a Policy Mobile App Management Policy Id Included Group ID
func ValidatePolicyMobileAppManagementPolicyIdIncludedGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMobileAppManagementPolicyIdIncludedGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Mobile App Management Policy Id Included Group ID
func (id PolicyMobileAppManagementPolicyIdIncludedGroupId) ID() string {
	fmtString := "/policies/mobileAppManagementPolicies/%s/includedGroups/%s"
	return fmt.Sprintf(fmtString, id.MobilityManagementPolicyId, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Mobile App Management Policy Id Included Group ID
func (id PolicyMobileAppManagementPolicyIdIncludedGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("mobileAppManagementPolicies", "mobileAppManagementPolicies", "mobileAppManagementPolicies"),
		resourceids.UserSpecifiedSegment("mobilityManagementPolicyId", "mobilityManagementPolicyId"),
		resourceids.StaticSegment("includedGroups", "includedGroups", "includedGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
	}
}

// String returns a human-readable description of this Policy Mobile App Management Policy Id Included Group ID
func (id PolicyMobileAppManagementPolicyIdIncludedGroupId) String() string {
	components := []string{
		fmt.Sprintf("Mobility Management Policy: %q", id.MobilityManagementPolicyId),
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Policy Mobile App Management Policy Id Included Group (%s)", strings.Join(components, "\n"))
}
