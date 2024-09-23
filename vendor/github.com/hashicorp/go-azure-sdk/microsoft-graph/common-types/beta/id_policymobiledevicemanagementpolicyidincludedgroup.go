package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{}

// PolicyMobileDeviceManagementPolicyIdIncludedGroupId is a struct representing the Resource ID for a Policy Mobile Device Management Policy Id Included Group
type PolicyMobileDeviceManagementPolicyIdIncludedGroupId struct {
	MobilityManagementPolicyId string
	GroupId                    string
}

// NewPolicyMobileDeviceManagementPolicyIdIncludedGroupID returns a new PolicyMobileDeviceManagementPolicyIdIncludedGroupId struct
func NewPolicyMobileDeviceManagementPolicyIdIncludedGroupID(mobilityManagementPolicyId string, groupId string) PolicyMobileDeviceManagementPolicyIdIncludedGroupId {
	return PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
		MobilityManagementPolicyId: mobilityManagementPolicyId,
		GroupId:                    groupId,
	}
}

// ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupID parses 'input' into a PolicyMobileDeviceManagementPolicyIdIncludedGroupId
func ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupID(input string) (*PolicyMobileDeviceManagementPolicyIdIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileDeviceManagementPolicyIdIncludedGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileDeviceManagementPolicyIdIncludedGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupIDInsensitively parses 'input' case-insensitively into a PolicyMobileDeviceManagementPolicyIdIncludedGroupId
// note: this method should only be used for API response data and not user input
func ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupIDInsensitively(input string) (*PolicyMobileDeviceManagementPolicyIdIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMobileDeviceManagementPolicyIdIncludedGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMobileDeviceManagementPolicyIdIncludedGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyMobileDeviceManagementPolicyIdIncludedGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobilityManagementPolicyId, ok = input.Parsed["mobilityManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", input)
	}

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	return nil
}

// ValidatePolicyMobileDeviceManagementPolicyIdIncludedGroupID checks that 'input' can be parsed as a Policy Mobile Device Management Policy Id Included Group ID
func ValidatePolicyMobileDeviceManagementPolicyIdIncludedGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Mobile Device Management Policy Id Included Group ID
func (id PolicyMobileDeviceManagementPolicyIdIncludedGroupId) ID() string {
	fmtString := "/policies/mobileDeviceManagementPolicies/%s/includedGroups/%s"
	return fmt.Sprintf(fmtString, id.MobilityManagementPolicyId, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Mobile Device Management Policy Id Included Group ID
func (id PolicyMobileDeviceManagementPolicyIdIncludedGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("mobileDeviceManagementPolicies", "mobileDeviceManagementPolicies", "mobileDeviceManagementPolicies"),
		resourceids.UserSpecifiedSegment("mobilityManagementPolicyId", "mobilityManagementPolicyId"),
		resourceids.StaticSegment("includedGroups", "includedGroups", "includedGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
	}
}

// String returns a human-readable description of this Policy Mobile Device Management Policy Id Included Group ID
func (id PolicyMobileDeviceManagementPolicyIdIncludedGroupId) String() string {
	components := []string{
		fmt.Sprintf("Mobility Management Policy: %q", id.MobilityManagementPolicyId),
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Policy Mobile Device Management Policy Id Included Group (%s)", strings.Join(components, "\n"))
}
