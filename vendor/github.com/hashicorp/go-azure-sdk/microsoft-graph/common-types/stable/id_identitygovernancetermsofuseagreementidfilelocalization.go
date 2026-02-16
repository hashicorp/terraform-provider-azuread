package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId{}

// IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement Id File Localization
type IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId struct {
	AgreementId                 string
	AgreementFileLocalizationId string
}

// NewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID returns a new IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId struct
func NewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID(agreementId string, agreementFileLocalizationId string) IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId {
	return IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId{
		AgreementId:                 agreementId,
		AgreementFileLocalizationId: agreementFileLocalizationId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID parses 'input' into a IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId
func ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementId, ok = input.Parsed["agreementId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementId", input)
	}

	if id.AgreementFileLocalizationId, ok = input.Parsed["agreementFileLocalizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementFileLocalizationId", input)
	}

	return nil
}

// ValidateIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement Id File Localization ID
func ValidateIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement Id File Localization ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreements/%s/file/localizations/%s"
	return fmt.Sprintf(fmtString, id.AgreementId, id.AgreementFileLocalizationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement Id File Localization ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementId", "agreementId"),
		resourceids.StaticSegment("file", "file", "file"),
		resourceids.StaticSegment("localizations", "localizations", "localizations"),
		resourceids.UserSpecifiedSegment("agreementFileLocalizationId", "agreementFileLocalizationId"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement Id File Localization ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId) String() string {
	components := []string{
		fmt.Sprintf("Agreement: %q", id.AgreementId),
		fmt.Sprintf("Agreement File Localization: %q", id.AgreementFileLocalizationId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement Id File Localization (%s)", strings.Join(components, "\n"))
}
