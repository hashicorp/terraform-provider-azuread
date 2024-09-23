package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{}

// IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement Id File Id Version
type IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId struct {
	AgreementId                 string
	AgreementFileLocalizationId string
	AgreementFileVersionId      string
}

// NewIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID returns a new IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId struct
func NewIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(agreementId string, agreementFileLocalizationId string, agreementFileVersionId string) IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId {
	return IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{
		AgreementId:                 agreementId,
		AgreementFileLocalizationId: agreementFileLocalizationId,
		AgreementFileVersionId:      agreementFileVersionId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID parses 'input' into a IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId
func ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement Id File Id Version ID
func ValidateIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement Id File Id Version ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreements/%s/files/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.AgreementId, id.AgreementFileLocalizationId, id.AgreementFileVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement Id File Id Version ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementId", "agreementId"),
		resourceids.StaticSegment("files", "files", "files"),
		resourceids.UserSpecifiedSegment("agreementFileLocalizationId", "agreementFileLocalizationId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("agreementFileVersionId", "agreementFileVersionId"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement Id File Id Version ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Agreement: %q", id.AgreementId),
		fmt.Sprintf("Agreement File Localization: %q", id.AgreementFileLocalizationId),
		fmt.Sprintf("Agreement File Version: %q", id.AgreementFileVersionId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement Id File Id Version (%s)", strings.Join(components, "\n"))
}
