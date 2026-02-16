package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{}

// PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId is a struct representing the Resource ID for a Policy Authorization Policy Id Default User Role Override
type PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId struct {
	AuthorizationPolicyId     string
	DefaultUserRoleOverrideId string
}

// NewPolicyAuthorizationPolicyIdDefaultUserRoleOverrideID returns a new PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId struct
func NewPolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(authorizationPolicyId string, defaultUserRoleOverrideId string) PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId {
	return PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{
		AuthorizationPolicyId:     authorizationPolicyId,
		DefaultUserRoleOverrideId: defaultUserRoleOverrideId,
	}
}

// ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID parses 'input' into a PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId
func ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(input string) (*PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideIDInsensitively parses 'input' case-insensitively into a PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideIDInsensitively(input string) (*PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthorizationPolicyId, ok = input.Parsed["authorizationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authorizationPolicyId", input)
	}

	if id.DefaultUserRoleOverrideId, ok = input.Parsed["defaultUserRoleOverrideId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "defaultUserRoleOverrideId", input)
	}

	return nil
}

// ValidatePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID checks that 'input' can be parsed as a Policy Authorization Policy Id Default User Role Override ID
func ValidatePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authorization Policy Id Default User Role Override ID
func (id PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId) ID() string {
	fmtString := "/policies/authorizationPolicy/%s/defaultUserRoleOverrides/%s"
	return fmt.Sprintf(fmtString, id.AuthorizationPolicyId, id.DefaultUserRoleOverrideId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authorization Policy Id Default User Role Override ID
func (id PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("authorizationPolicy", "authorizationPolicy", "authorizationPolicy"),
		resourceids.UserSpecifiedSegment("authorizationPolicyId", "authorizationPolicyId"),
		resourceids.StaticSegment("defaultUserRoleOverrides", "defaultUserRoleOverrides", "defaultUserRoleOverrides"),
		resourceids.UserSpecifiedSegment("defaultUserRoleOverrideId", "defaultUserRoleOverrideId"),
	}
}

// String returns a human-readable description of this Policy Authorization Policy Id Default User Role Override ID
func (id PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId) String() string {
	components := []string{
		fmt.Sprintf("Authorization Policy: %q", id.AuthorizationPolicyId),
		fmt.Sprintf("Default User Role Override: %q", id.DefaultUserRoleOverrideId),
	}
	return fmt.Sprintf("Policy Authorization Policy Id Default User Role Override (%s)", strings.Join(components, "\n"))
}
