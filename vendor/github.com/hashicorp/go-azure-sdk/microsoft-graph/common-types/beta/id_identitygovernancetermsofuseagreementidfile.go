package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileId{}

// IdentityGovernanceTermsOfUseAgreementIdFileId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement Id File
type IdentityGovernanceTermsOfUseAgreementIdFileId struct {
	AgreementId                 string
	AgreementFileLocalizationId string
}

// NewIdentityGovernanceTermsOfUseAgreementIdFileID returns a new IdentityGovernanceTermsOfUseAgreementIdFileId struct
func NewIdentityGovernanceTermsOfUseAgreementIdFileID(agreementId string, agreementFileLocalizationId string) IdentityGovernanceTermsOfUseAgreementIdFileId {
	return IdentityGovernanceTermsOfUseAgreementIdFileId{
		AgreementId:                 agreementId,
		AgreementFileLocalizationId: agreementFileLocalizationId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileID parses 'input' into a IdentityGovernanceTermsOfUseAgreementIdFileId
func ParseIdentityGovernanceTermsOfUseAgreementIdFileID(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementIdFileIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementIdFileId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementIdFileIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementIdFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementIdFileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementIdFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementIdFileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementId, ok = input.Parsed["agreementId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementId", input)
	}

	if id.AgreementFileLocalizationId, ok = input.Parsed["agreementFileLocalizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementFileLocalizationId", input)
	}

	return nil
}

// ValidateIdentityGovernanceTermsOfUseAgreementIdFileID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement Id File ID
func ValidateIdentityGovernanceTermsOfUseAgreementIdFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement Id File ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreements/%s/files/%s"
	return fmt.Sprintf(fmtString, id.AgreementId, id.AgreementFileLocalizationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement Id File ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementId", "agreementId"),
		resourceids.StaticSegment("files", "files", "files"),
		resourceids.UserSpecifiedSegment("agreementFileLocalizationId", "agreementFileLocalizationId"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement Id File ID
func (id IdentityGovernanceTermsOfUseAgreementIdFileId) String() string {
	components := []string{
		fmt.Sprintf("Agreement: %q", id.AgreementId),
		fmt.Sprintf("Agreement File Localization: %q", id.AgreementFileLocalizationId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement Id File (%s)", strings.Join(components, "\n"))
}
