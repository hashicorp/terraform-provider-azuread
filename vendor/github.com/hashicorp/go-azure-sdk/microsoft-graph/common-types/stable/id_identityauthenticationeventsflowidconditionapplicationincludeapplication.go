package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{}

// IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId is a struct representing the Resource ID for a Identity Authentication Events Flow Id Condition Application Include Application
type IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId struct {
	AuthenticationEventsFlowId              string
	AuthenticationConditionApplicationAppId string
}

// NewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID returns a new IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId struct
func NewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(authenticationEventsFlowId string, authenticationConditionApplicationAppId string) IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId {
	return IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
		AuthenticationEventsFlowId:              authenticationEventsFlowId,
		AuthenticationConditionApplicationAppId: authenticationConditionApplicationAppId,
	}
}

// ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID parses 'input' into a IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId
func ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(input string) (*IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationIDInsensitively parses 'input' case-insensitively into a IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId
// note: this method should only be used for API response data and not user input
func ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationIDInsensitively(input string) (*IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationEventsFlowId, ok = input.Parsed["authenticationEventsFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationEventsFlowId", input)
	}

	if id.AuthenticationConditionApplicationAppId, ok = input.Parsed["authenticationConditionApplicationAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationConditionApplicationAppId", input)
	}

	return nil
}

// ValidateIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID checks that 'input' can be parsed as a Identity Authentication Events Flow Id Condition Application Include Application ID
func ValidateIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Authentication Events Flow Id Condition Application Include Application ID
func (id IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId) ID() string {
	fmtString := "/identity/authenticationEventsFlows/%s/conditions/applications/includeApplications/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationEventsFlowId, id.AuthenticationConditionApplicationAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Authentication Events Flow Id Condition Application Include Application ID
func (id IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("authenticationEventsFlows", "authenticationEventsFlows", "authenticationEventsFlows"),
		resourceids.UserSpecifiedSegment("authenticationEventsFlowId", "authenticationEventsFlowId"),
		resourceids.StaticSegment("conditions", "conditions", "conditions"),
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.StaticSegment("includeApplications", "includeApplications", "includeApplications"),
		resourceids.UserSpecifiedSegment("authenticationConditionApplicationAppId", "authenticationConditionApplicationAppId"),
	}
}

// String returns a human-readable description of this Identity Authentication Events Flow Id Condition Application Include Application ID
func (id IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Events Flow: %q", id.AuthenticationEventsFlowId),
		fmt.Sprintf("Authentication Condition Application App: %q", id.AuthenticationConditionApplicationAppId),
	}
	return fmt.Sprintf("Identity Authentication Events Flow Id Condition Application Include Application (%s)", strings.Join(components, "\n"))
}
