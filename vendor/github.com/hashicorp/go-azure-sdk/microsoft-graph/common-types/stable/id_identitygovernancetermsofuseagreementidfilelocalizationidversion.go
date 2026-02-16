package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{}

// IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement Id File Localization Id Version
type IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId struct {
	AgreementId                 string
	AgreementFileLocalizationId string
	AgreementFileVersionId      string
}

// NewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID returns a new IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId struct
func NewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(agreementId string, agreementFileLocalizationId string, agreementFileVersionId string) IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId {
	return IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{
		AgreementId:                 agreementId,
		AgreementFileLocalizationId: agreementFileLocalizationId,
		AgreementFileVersionId:      agreementFileVersionId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID parses 'input' into a IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId
func ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementId, ok = input.Parsed["agreementId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementId", input)
	}

	if id.AgreementFileLocalizationId, ok = input.Parsed["agreementFileLocalizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementFileLocalizationId", input)
	}

	if id.AgreementFileVersionId, ok = input.Parsed["agreementFileVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementFileVersionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement Id File Localization Id Version ID
func ValidateIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement Id File Localization Id Version ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreements/%s/file/localizations/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.AgreementId, id.AgreementFileLocalizationId, id.AgreementFileVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement Id File Localization Id Version ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementId", "agreementId"),
		resourceids.StaticSegment("file", "file", "file"),
		resourceids.StaticSegment("localizations", "localizations", "localizations"),
		resourceids.UserSpecifiedSegment("agreementFileLocalizationId", "agreementFileLocalizationId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("agreementFileVersionId", "agreementFileVersionId"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement Id File Localization Id Version ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Agreement: %q", id.AgreementId),
		fmt.Sprintf("Agreement File Localization: %q", id.AgreementFileLocalizationId),
		fmt.Sprintf("Agreement File Version: %q", id.AgreementFileVersionId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement Id File Localization Id Version (%s)", strings.Join(components, "\n"))
}
