package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{}

// IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId is a struct representing the Resource ID for a Identity Governance App Consent App Consent Request Id User Consent Request
type IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
}

// NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID returns a new IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId struct
func NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(appConsentRequestId string, userConsentRequestId string) IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId {
	return IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
	}
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID parses 'input' into a IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId
func ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(input string) (*IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIDInsensitively(input string) (*IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	if id.UserConsentRequestId, ok = input.Parsed["userConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID checks that 'input' can be parsed as a Identity Governance App Consent App Consent Request Id User Consent Request ID
func ValidateIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance App Consent App Consent Request Id User Consent Request ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId) ID() string {
	fmtString := "/identityGovernance/appConsent/appConsentRequests/%s/userConsentRequests/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance App Consent App Consent Request Id User Consent Request ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("appConsent", "appConsent", "appConsent"),
		resourceids.StaticSegment("appConsentRequests", "appConsentRequests", "appConsentRequests"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance App Consent App Consent Request Id User Consent Request ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
	}
	return fmt.Sprintf("Identity Governance App Consent App Consent Request Id User Consent Request (%s)", strings.Join(components, "\n"))
}
