package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{}

// IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId is a struct representing the Resource ID for a Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Authentication Method Load Start On Authentication Method Load Start External Users Self Service Sign Up Identity Provider
type IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId struct {
	AuthenticationEventsFlowId string
	IdentityProviderBaseId     string
}

// NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID returns a new IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId struct
func NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(authenticationEventsFlowId string, identityProviderBaseId string) IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId {
	return IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{
		AuthenticationEventsFlowId: authenticationEventsFlowId,
		IdentityProviderBaseId:     identityProviderBaseId,
	}
}

// ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID parses 'input' into a IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId
func ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(input string) (*IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderIDInsensitively(input string) (*IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationEventsFlowId, ok = input.Parsed["authenticationEventsFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationEventsFlowId", input)
	}

	if id.IdentityProviderBaseId, ok = input.Parsed["identityProviderBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", input)
	}

	return nil
}

// ValidateIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID checks that 'input' can be parsed as a Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Authentication Method Load Start On Authentication Method Load Start External Users Self Service Sign Up Identity Provider ID
func ValidateIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Authentication Method Load Start On Authentication Method Load Start External Users Self Service Sign Up Identity Provider ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId) ID() string {
	fmtString := "/identity/authenticationEventsFlows/%s/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationEventsFlowId, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Authentication Method Load Start On Authentication Method Load Start External Users Self Service Sign Up Identity Provider ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("authenticationEventsFlows", "authenticationEventsFlows", "authenticationEventsFlows"),
		resourceids.UserSpecifiedSegment("authenticationEventsFlowId", "authenticationEventsFlowId"),
		resourceids.StaticSegment("externalUsersSelfServiceSignUpEventsFlow", "externalUsersSelfServiceSignUpEventsFlow", "externalUsersSelfServiceSignUpEventsFlow"),
		resourceids.StaticSegment("onAuthenticationMethodLoadStart", "onAuthenticationMethodLoadStart", "onAuthenticationMethodLoadStart"),
		resourceids.StaticSegment("onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp", "onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp", "onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp"),
		resourceids.StaticSegment("identityProviders", "identityProviders", "identityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseId"),
	}
}

// String returns a human-readable description of this Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Authentication Method Load Start On Authentication Method Load Start External Users Self Service Sign Up Identity Provider ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Events Flow: %q", id.AuthenticationEventsFlowId),
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Authentication Method Load Start On Authentication Method Load Start External Users Self Service Sign Up Identity Provider (%s)", strings.Join(components, "\n"))
}
