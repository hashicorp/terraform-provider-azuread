package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdTokenIssuancePolicyId{}

// ApplicationIdTokenIssuancePolicyId is a struct representing the Resource ID for a Application Id Token Issuance Policy
type ApplicationIdTokenIssuancePolicyId struct {
	ApplicationId         string
	TokenIssuancePolicyId string
}

// NewApplicationIdTokenIssuancePolicyID returns a new ApplicationIdTokenIssuancePolicyId struct
func NewApplicationIdTokenIssuancePolicyID(applicationId string, tokenIssuancePolicyId string) ApplicationIdTokenIssuancePolicyId {
	return ApplicationIdTokenIssuancePolicyId{
		ApplicationId:         applicationId,
		TokenIssuancePolicyId: tokenIssuancePolicyId,
	}
}

// ParseApplicationIdTokenIssuancePolicyID parses 'input' into a ApplicationIdTokenIssuancePolicyId
func ParseApplicationIdTokenIssuancePolicyID(input string) (*ApplicationIdTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdTokenIssuancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdTokenIssuancePolicyIDInsensitively parses 'input' case-insensitively into a ApplicationIdTokenIssuancePolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdTokenIssuancePolicyIDInsensitively(input string) (*ApplicationIdTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdTokenIssuancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdTokenIssuancePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.TokenIssuancePolicyId, ok = input.Parsed["tokenIssuancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", input)
	}

	return nil
}

// ValidateApplicationIdTokenIssuancePolicyID checks that 'input' can be parsed as a Application Id Token Issuance Policy ID
func ValidateApplicationIdTokenIssuancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdTokenIssuancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Token Issuance Policy ID
func (id ApplicationIdTokenIssuancePolicyId) ID() string {
	fmtString := "/applications/%s/tokenIssuancePolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.TokenIssuancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Token Issuance Policy ID
func (id ApplicationIdTokenIssuancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("tokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyId"),
	}
}

// String returns a human-readable description of this Application Id Token Issuance Policy ID
func (id ApplicationIdTokenIssuancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
	}
	return fmt.Sprintf("Application Id Token Issuance Policy (%s)", strings.Join(components, "\n"))
}
