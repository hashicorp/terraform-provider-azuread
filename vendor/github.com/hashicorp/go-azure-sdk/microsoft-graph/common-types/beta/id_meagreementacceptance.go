package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAgreementAcceptanceId{}

// MeAgreementAcceptanceId is a struct representing the Resource ID for a Me Agreement Acceptance
type MeAgreementAcceptanceId struct {
	AgreementAcceptanceId string
}

// NewMeAgreementAcceptanceID returns a new MeAgreementAcceptanceId struct
func NewMeAgreementAcceptanceID(agreementAcceptanceId string) MeAgreementAcceptanceId {
	return MeAgreementAcceptanceId{
		AgreementAcceptanceId: agreementAcceptanceId,
	}
}

// ParseMeAgreementAcceptanceID parses 'input' into a MeAgreementAcceptanceId
func ParseMeAgreementAcceptanceID(input string) (*MeAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAgreementAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAgreementAcceptanceIDInsensitively parses 'input' case-insensitively into a MeAgreementAcceptanceId
// note: this method should only be used for API response data and not user input
func ParseMeAgreementAcceptanceIDInsensitively(input string) (*MeAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAgreementAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAgreementAcceptanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementAcceptanceId, ok = input.Parsed["agreementAcceptanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementAcceptanceId", input)
	}

	return nil
}

// ValidateMeAgreementAcceptanceID checks that 'input' can be parsed as a Me Agreement Acceptance ID
func ValidateMeAgreementAcceptanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAgreementAcceptanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Agreement Acceptance ID
func (id MeAgreementAcceptanceId) ID() string {
	fmtString := "/me/agreementAcceptances/%s"
	return fmt.Sprintf(fmtString, id.AgreementAcceptanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Agreement Acceptance ID
func (id MeAgreementAcceptanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("agreementAcceptances", "agreementAcceptances", "agreementAcceptances"),
		resourceids.UserSpecifiedSegment("agreementAcceptanceId", "agreementAcceptanceId"),
	}
}

// String returns a human-readable description of this Me Agreement Acceptance ID
func (id MeAgreementAcceptanceId) String() string {
	components := []string{
		fmt.Sprintf("Agreement Acceptance: %q", id.AgreementAcceptanceId),
	}
	return fmt.Sprintf("Me Agreement Acceptance (%s)", strings.Join(components, "\n"))
}
