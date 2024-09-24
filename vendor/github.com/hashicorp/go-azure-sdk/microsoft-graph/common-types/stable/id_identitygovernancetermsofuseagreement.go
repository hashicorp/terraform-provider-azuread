package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementId{}

// IdentityGovernanceTermsOfUseAgreementId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement
type IdentityGovernanceTermsOfUseAgreementId struct {
	AgreementId string
}

// NewIdentityGovernanceTermsOfUseAgreementID returns a new IdentityGovernanceTermsOfUseAgreementId struct
func NewIdentityGovernanceTermsOfUseAgreementID(agreementId string) IdentityGovernanceTermsOfUseAgreementId {
	return IdentityGovernanceTermsOfUseAgreementId{
		AgreementId: agreementId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementID parses 'input' into a IdentityGovernanceTermsOfUseAgreementId
func ParseIdentityGovernanceTermsOfUseAgreementID(input string) (*IdentityGovernanceTermsOfUseAgreementId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementId, ok = input.Parsed["agreementId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementId", input)
	}

	return nil
}

// ValidateIdentityGovernanceTermsOfUseAgreementID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement ID
func ValidateIdentityGovernanceTermsOfUseAgreementID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement ID
func (id IdentityGovernanceTermsOfUseAgreementId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreements/%s"
	return fmt.Sprintf(fmtString, id.AgreementId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement ID
func (id IdentityGovernanceTermsOfUseAgreementId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementId", "agreementId"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement ID
func (id IdentityGovernanceTermsOfUseAgreementId) String() string {
	components := []string{
		fmt.Sprintf("Agreement: %q", id.AgreementId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement (%s)", strings.Join(components, "\n"))
}
