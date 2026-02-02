package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{}

// IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId is a struct representing the Resource ID for a Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Attribute Collection On Attribute Collection External Users Self Service Sign Up Attribute
type IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId struct {
	AuthenticationEventsFlowId  string
	IdentityUserFlowAttributeId string
}

// NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID returns a new IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId struct
func NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(authenticationEventsFlowId string, identityUserFlowAttributeId string) IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId {
	return IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{
		AuthenticationEventsFlowId:  authenticationEventsFlowId,
		IdentityUserFlowAttributeId: identityUserFlowAttributeId,
	}
}

// ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID parses 'input' into a IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId
func ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(input string) (*IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeIDInsensitively parses 'input' case-insensitively into a IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId
// note: this method should only be used for API response data and not user input
func ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeIDInsensitively(input string) (*IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationEventsFlowId, ok = input.Parsed["authenticationEventsFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationEventsFlowId", input)
	}

	if id.IdentityUserFlowAttributeId, ok = input.Parsed["identityUserFlowAttributeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowAttributeId", input)
	}

	return nil
}

// ValidateIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID checks that 'input' can be parsed as a Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Attribute Collection On Attribute Collection External Users Self Service Sign Up Attribute ID
func ValidateIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Attribute Collection On Attribute Collection External Users Self Service Sign Up Attribute ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId) ID() string {
	fmtString := "/identity/authenticationEventsFlows/%s/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationEventsFlowId, id.IdentityUserFlowAttributeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Attribute Collection On Attribute Collection External Users Self Service Sign Up Attribute ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("authenticationEventsFlows", "authenticationEventsFlows", "authenticationEventsFlows"),
		resourceids.UserSpecifiedSegment("authenticationEventsFlowId", "authenticationEventsFlowId"),
		resourceids.StaticSegment("externalUsersSelfServiceSignUpEventsFlow", "externalUsersSelfServiceSignUpEventsFlow", "externalUsersSelfServiceSignUpEventsFlow"),
		resourceids.StaticSegment("onAttributeCollection", "onAttributeCollection", "onAttributeCollection"),
		resourceids.StaticSegment("onAttributeCollectionExternalUsersSelfServiceSignUp", "onAttributeCollectionExternalUsersSelfServiceSignUp", "onAttributeCollectionExternalUsersSelfServiceSignUp"),
		resourceids.StaticSegment("attributes", "attributes", "attributes"),
		resourceids.UserSpecifiedSegment("identityUserFlowAttributeId", "identityUserFlowAttributeId"),
	}
}

// String returns a human-readable description of this Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Attribute Collection On Attribute Collection External Users Self Service Sign Up Attribute ID
func (id IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Events Flow: %q", id.AuthenticationEventsFlowId),
		fmt.Sprintf("Identity User Flow Attribute: %q", id.IdentityUserFlowAttributeId),
	}
	return fmt.Sprintf("Identity Authentication Events Flow Id External Users Self Service Sign Up Events Flow On Attribute Collection On Attribute Collection External Users Self Service Sign Up Attribute (%s)", strings.Join(components, "\n"))
}
