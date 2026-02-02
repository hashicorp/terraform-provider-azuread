package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{}

// IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId is a struct representing the Resource ID for a Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow Condition Application Include Application
type IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId struct {
	AuthenticationEventsFlowId              string
	AuthenticationConditionApplicationAppId string
}

// NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID returns a new IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId struct
func NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(authenticationEventsFlowId string, authenticationConditionApplicationAppId string) IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId {
	return IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{
		AuthenticationEventsFlowId:              authenticationEventsFlowId,
		AuthenticationConditionApplicationAppId: authenticationConditionApplicationAppId,
	}
}

// ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID parses 'input' into a IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId
func ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(input string) (*IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationIDInsensitively parses 'input' case-insensitively into a IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId
// note: this method should only be used for API response data and not user input
func ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationIDInsensitively(input string) (*IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationEventsFlowId, ok = input.Parsed["authenticationEventsFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationEventsFlowId", input)
	}

	if id.AuthenticationConditionApplicationAppId, ok = input.Parsed["authenticationConditionApplicationAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationConditionApplicationAppId", input)
	}

	return nil
}

// ValidateIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID checks that 'input' can be parsed as a Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow Condition Application Include Application ID
func ValidateIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow Condition Application Include Application ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId) ID() string {
	fmtString := "/identity/authenticationEventsFlows/%s/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationEventsFlowId, id.AuthenticationConditionApplicationAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow Condition Application Include Application ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("authenticationEventsFlows", "authenticationEventsFlows", "authenticationEventsFlows"),
		resourceids.UserSpecifiedSegment("authenticationEventsFlowId", "authenticationEventsFlowId"),
		resourceids.StaticSegment("externalUsersSelfServiceSignUpEventsFlow", "externalUsersSelfServiceSignUpEventsFlow", "externalUsersSelfServiceSignUpEventsFlow"),
		resourceids.StaticSegment("conditions", "conditions", "conditions"),
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.StaticSegment("includeApplications", "includeApplications", "includeApplications"),
		resourceids.UserSpecifiedSegment("authenticationConditionApplicationAppId", "authenticationConditionApplicationAppId"),
	}
}

// String returns a human-readable description of this Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow Condition Application Include Application ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Events Flow: %q", id.AuthenticationEventsFlowId),
		fmt.Sprintf("Authentication Condition Application App: %q", id.AuthenticationConditionApplicationAppId),
	}
	return fmt.Sprintf("Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow Condition Application Include Application (%s)", strings.Join(components, "\n"))
}
