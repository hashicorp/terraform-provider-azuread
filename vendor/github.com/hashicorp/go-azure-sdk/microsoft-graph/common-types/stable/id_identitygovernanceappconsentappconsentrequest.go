package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestId{}

// IdentityGovernanceAppConsentAppConsentRequestId is a struct representing the Resource ID for a Identity Governance App Consent App Consent Request
type IdentityGovernanceAppConsentAppConsentRequestId struct {
	AppConsentRequestId string
}

// NewIdentityGovernanceAppConsentAppConsentRequestID returns a new IdentityGovernanceAppConsentAppConsentRequestId struct
func NewIdentityGovernanceAppConsentAppConsentRequestID(appConsentRequestId string) IdentityGovernanceAppConsentAppConsentRequestId {
	return IdentityGovernanceAppConsentAppConsentRequestId{
		AppConsentRequestId: appConsentRequestId,
	}
}

// ParseIdentityGovernanceAppConsentAppConsentRequestID parses 'input' into a IdentityGovernanceAppConsentAppConsentRequestId
func ParseIdentityGovernanceAppConsentAppConsentRequestID(input string) (*IdentityGovernanceAppConsentAppConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAppConsentAppConsentRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAppConsentAppConsentRequestIDInsensitively(input string) (*IdentityGovernanceAppConsentAppConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAppConsentAppConsentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAppConsentAppConsentRequestID checks that 'input' can be parsed as a Identity Governance App Consent App Consent Request ID
func ValidateIdentityGovernanceAppConsentAppConsentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAppConsentAppConsentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance App Consent App Consent Request ID
func (id IdentityGovernanceAppConsentAppConsentRequestId) ID() string {
	fmtString := "/identityGovernance/appConsent/appConsentRequests/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance App Consent App Consent Request ID
func (id IdentityGovernanceAppConsentAppConsentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("appConsent", "appConsent", "appConsent"),
		resourceids.StaticSegment("appConsentRequests", "appConsentRequests", "appConsentRequests"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance App Consent App Consent Request ID
func (id IdentityGovernanceAppConsentAppConsentRequestId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
	}
	return fmt.Sprintf("Identity Governance App Consent App Consent Request (%s)", strings.Join(components, "\n"))
}
