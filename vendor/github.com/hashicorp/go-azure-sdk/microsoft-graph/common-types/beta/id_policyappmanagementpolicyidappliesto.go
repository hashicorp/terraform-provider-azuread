package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAppManagementPolicyIdAppliesToId{}

// PolicyAppManagementPolicyIdAppliesToId is a struct representing the Resource ID for a Policy App Management Policy Id Applies To
type PolicyAppManagementPolicyIdAppliesToId struct {
	AppManagementPolicyId string
	DirectoryObjectId     string
}

// NewPolicyAppManagementPolicyIdAppliesToID returns a new PolicyAppManagementPolicyIdAppliesToId struct
func NewPolicyAppManagementPolicyIdAppliesToID(appManagementPolicyId string, directoryObjectId string) PolicyAppManagementPolicyIdAppliesToId {
	return PolicyAppManagementPolicyIdAppliesToId{
		AppManagementPolicyId: appManagementPolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyAppManagementPolicyIdAppliesToID parses 'input' into a PolicyAppManagementPolicyIdAppliesToId
func ParsePolicyAppManagementPolicyIdAppliesToID(input string) (*PolicyAppManagementPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAppManagementPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAppManagementPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAppManagementPolicyIdAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyAppManagementPolicyIdAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyAppManagementPolicyIdAppliesToIDInsensitively(input string) (*PolicyAppManagementPolicyIdAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAppManagementPolicyIdAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAppManagementPolicyIdAppliesToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAppManagementPolicyIdAppliesToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppManagementPolicyId, ok = input.Parsed["appManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidatePolicyAppManagementPolicyIdAppliesToID checks that 'input' can be parsed as a Policy App Management Policy Id Applies To ID
func ValidatePolicyAppManagementPolicyIdAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAppManagementPolicyIdAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy App Management Policy Id Applies To ID
func (id PolicyAppManagementPolicyIdAppliesToId) ID() string {
	fmtString := "/policies/appManagementPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.AppManagementPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy App Management Policy Id Applies To ID
func (id PolicyAppManagementPolicyIdAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("appManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyId"),
		resourceids.StaticSegment("appliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Policy App Management Policy Id Applies To ID
func (id PolicyAppManagementPolicyIdAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy App Management Policy Id Applies To (%s)", strings.Join(components, "\n"))
}
