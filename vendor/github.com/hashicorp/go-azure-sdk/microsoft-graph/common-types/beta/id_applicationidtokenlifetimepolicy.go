package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdTokenLifetimePolicyId{}

// ApplicationIdTokenLifetimePolicyId is a struct representing the Resource ID for a Application Id Token Lifetime Policy
type ApplicationIdTokenLifetimePolicyId struct {
	ApplicationId         string
	TokenLifetimePolicyId string
}

// NewApplicationIdTokenLifetimePolicyID returns a new ApplicationIdTokenLifetimePolicyId struct
func NewApplicationIdTokenLifetimePolicyID(applicationId string, tokenLifetimePolicyId string) ApplicationIdTokenLifetimePolicyId {
	return ApplicationIdTokenLifetimePolicyId{
		ApplicationId:         applicationId,
		TokenLifetimePolicyId: tokenLifetimePolicyId,
	}
}

// ParseApplicationIdTokenLifetimePolicyID parses 'input' into a ApplicationIdTokenLifetimePolicyId
func ParseApplicationIdTokenLifetimePolicyID(input string) (*ApplicationIdTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdTokenLifetimePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdTokenLifetimePolicyIDInsensitively parses 'input' case-insensitively into a ApplicationIdTokenLifetimePolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdTokenLifetimePolicyIDInsensitively(input string) (*ApplicationIdTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdTokenLifetimePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdTokenLifetimePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.TokenLifetimePolicyId, ok = input.Parsed["tokenLifetimePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", input)
	}

	return nil
}

// ValidateApplicationIdTokenLifetimePolicyID checks that 'input' can be parsed as a Application Id Token Lifetime Policy ID
func ValidateApplicationIdTokenLifetimePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdTokenLifetimePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Token Lifetime Policy ID
func (id ApplicationIdTokenLifetimePolicyId) ID() string {
	fmtString := "/applications/%s/tokenLifetimePolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.TokenLifetimePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Token Lifetime Policy ID
func (id ApplicationIdTokenLifetimePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("tokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyId"),
	}
}

// String returns a human-readable description of this Application Id Token Lifetime Policy ID
func (id ApplicationIdTokenLifetimePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
	}
	return fmt.Sprintf("Application Id Token Lifetime Policy (%s)", strings.Join(components, "\n"))
}
