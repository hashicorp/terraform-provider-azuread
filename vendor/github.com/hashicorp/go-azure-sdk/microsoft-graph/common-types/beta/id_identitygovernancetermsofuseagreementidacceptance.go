package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{}

// IdentityGovernanceTermsOfUseAgreementIdAcceptanceId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement Id Acceptance
type IdentityGovernanceTermsOfUseAgreementIdAcceptanceId struct {
	AgreementId           string
	AgreementAcceptanceId string
}

// NewIdentityGovernanceTermsOfUseAgreementIdAcceptanceID returns a new IdentityGovernanceTermsOfUseAgreementIdAcceptanceId struct
func NewIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(agreementId string, agreementAcceptanceId string) IdentityGovernanceTermsOfUseAgreementIdAcceptanceId {
	return IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{
		AgreementId:           agreementId,
		AgreementAcceptanceId: agreementAcceptanceId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceID parses 'input' into a IdentityGovernanceTermsOfUseAgreementIdAcceptanceId
func ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(input string) (*IdentityGovernanceTermsOfUseAgreementIdAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementIdAcceptanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementIdAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementIdAcceptanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementId, ok = input.Parsed["agreementId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementId", input)
	}

	if id.AgreementAcceptanceId, ok = input.Parsed["agreementAcceptanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementAcceptanceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceTermsOfUseAgreementIdAcceptanceID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement Id Acceptance ID
func ValidateIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement Id Acceptance ID
func (id IdentityGovernanceTermsOfUseAgreementIdAcceptanceId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreements/%s/acceptances/%s"
	return fmt.Sprintf(fmtString, id.AgreementId, id.AgreementAcceptanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement Id Acceptance ID
func (id IdentityGovernanceTermsOfUseAgreementIdAcceptanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementId", "agreementId"),
		resourceids.StaticSegment("acceptances", "acceptances", "acceptances"),
		resourceids.UserSpecifiedSegment("agreementAcceptanceId", "agreementAcceptanceId"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement Id Acceptance ID
func (id IdentityGovernanceTermsOfUseAgreementIdAcceptanceId) String() string {
	components := []string{
		fmt.Sprintf("Agreement: %q", id.AgreementId),
		fmt.Sprintf("Agreement Acceptance: %q", id.AgreementAcceptanceId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement Id Acceptance (%s)", strings.Join(components, "\n"))
}
