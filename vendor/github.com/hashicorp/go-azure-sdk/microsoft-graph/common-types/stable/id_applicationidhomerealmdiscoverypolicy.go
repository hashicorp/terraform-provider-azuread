package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdHomeRealmDiscoveryPolicyId{}

// ApplicationIdHomeRealmDiscoveryPolicyId is a struct representing the Resource ID for a Application Id Home Realm Discovery Policy
type ApplicationIdHomeRealmDiscoveryPolicyId struct {
	ApplicationId              string
	HomeRealmDiscoveryPolicyId string
}

// NewApplicationIdHomeRealmDiscoveryPolicyID returns a new ApplicationIdHomeRealmDiscoveryPolicyId struct
func NewApplicationIdHomeRealmDiscoveryPolicyID(applicationId string, homeRealmDiscoveryPolicyId string) ApplicationIdHomeRealmDiscoveryPolicyId {
	return ApplicationIdHomeRealmDiscoveryPolicyId{
		ApplicationId:              applicationId,
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// ParseApplicationIdHomeRealmDiscoveryPolicyID parses 'input' into a ApplicationIdHomeRealmDiscoveryPolicyId
func ParseApplicationIdHomeRealmDiscoveryPolicyID(input string) (*ApplicationIdHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdHomeRealmDiscoveryPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdHomeRealmDiscoveryPolicyIDInsensitively parses 'input' case-insensitively into a ApplicationIdHomeRealmDiscoveryPolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdHomeRealmDiscoveryPolicyIDInsensitively(input string) (*ApplicationIdHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdHomeRealmDiscoveryPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdHomeRealmDiscoveryPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.HomeRealmDiscoveryPolicyId, ok = input.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", input)
	}

	return nil
}

// ValidateApplicationIdHomeRealmDiscoveryPolicyID checks that 'input' can be parsed as a Application Id Home Realm Discovery Policy ID
func ValidateApplicationIdHomeRealmDiscoveryPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdHomeRealmDiscoveryPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Home Realm Discovery Policy ID
func (id ApplicationIdHomeRealmDiscoveryPolicyId) ID() string {
	fmtString := "/applications/%s/homeRealmDiscoveryPolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.HomeRealmDiscoveryPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Home Realm Discovery Policy ID
func (id ApplicationIdHomeRealmDiscoveryPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyId"),
	}
}

// String returns a human-readable description of this Application Id Home Realm Discovery Policy ID
func (id ApplicationIdHomeRealmDiscoveryPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
	}
	return fmt.Sprintf("Application Id Home Realm Discovery Policy (%s)", strings.Join(components, "\n"))
}
