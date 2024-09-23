package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionThreatAssessmentRequestId{}

// MeInformationProtectionThreatAssessmentRequestId is a struct representing the Resource ID for a Me Information Protection Threat Assessment Request
type MeInformationProtectionThreatAssessmentRequestId struct {
	ThreatAssessmentRequestId string
}

// NewMeInformationProtectionThreatAssessmentRequestID returns a new MeInformationProtectionThreatAssessmentRequestId struct
func NewMeInformationProtectionThreatAssessmentRequestID(threatAssessmentRequestId string) MeInformationProtectionThreatAssessmentRequestId {
	return MeInformationProtectionThreatAssessmentRequestId{
		ThreatAssessmentRequestId: threatAssessmentRequestId,
	}
}

// ParseMeInformationProtectionThreatAssessmentRequestID parses 'input' into a MeInformationProtectionThreatAssessmentRequestId
func ParseMeInformationProtectionThreatAssessmentRequestID(input string) (*MeInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionThreatAssessmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionThreatAssessmentRequestIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionThreatAssessmentRequestId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionThreatAssessmentRequestIDInsensitively(input string) (*MeInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionThreatAssessmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionThreatAssessmentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ThreatAssessmentRequestId, ok = input.Parsed["threatAssessmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", input)
	}

	return nil
}

// ValidateMeInformationProtectionThreatAssessmentRequestID checks that 'input' can be parsed as a Me Information Protection Threat Assessment Request ID
func ValidateMeInformationProtectionThreatAssessmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionThreatAssessmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Threat Assessment Request ID
func (id MeInformationProtectionThreatAssessmentRequestId) ID() string {
	fmtString := "/me/informationProtection/threatAssessmentRequests/%s"
	return fmt.Sprintf(fmtString, id.ThreatAssessmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Threat Assessment Request ID
func (id MeInformationProtectionThreatAssessmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("threatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestId"),
	}
}

// String returns a human-readable description of this Me Information Protection Threat Assessment Request ID
func (id MeInformationProtectionThreatAssessmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
	}
	return fmt.Sprintf("Me Information Protection Threat Assessment Request (%s)", strings.Join(components, "\n"))
}
