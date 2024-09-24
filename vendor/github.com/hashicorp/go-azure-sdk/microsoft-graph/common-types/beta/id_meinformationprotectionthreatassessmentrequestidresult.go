package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionThreatAssessmentRequestIdResultId{}

// MeInformationProtectionThreatAssessmentRequestIdResultId is a struct representing the Resource ID for a Me Information Protection Threat Assessment Request Id Result
type MeInformationProtectionThreatAssessmentRequestIdResultId struct {
	ThreatAssessmentRequestId string
	ThreatAssessmentResultId  string
}

// NewMeInformationProtectionThreatAssessmentRequestIdResultID returns a new MeInformationProtectionThreatAssessmentRequestIdResultId struct
func NewMeInformationProtectionThreatAssessmentRequestIdResultID(threatAssessmentRequestId string, threatAssessmentResultId string) MeInformationProtectionThreatAssessmentRequestIdResultId {
	return MeInformationProtectionThreatAssessmentRequestIdResultId{
		ThreatAssessmentRequestId: threatAssessmentRequestId,
		ThreatAssessmentResultId:  threatAssessmentResultId,
	}
}

// ParseMeInformationProtectionThreatAssessmentRequestIdResultID parses 'input' into a MeInformationProtectionThreatAssessmentRequestIdResultId
func ParseMeInformationProtectionThreatAssessmentRequestIdResultID(input string) (*MeInformationProtectionThreatAssessmentRequestIdResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionThreatAssessmentRequestIdResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionThreatAssessmentRequestIdResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionThreatAssessmentRequestIdResultIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionThreatAssessmentRequestIdResultId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(input string) (*MeInformationProtectionThreatAssessmentRequestIdResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionThreatAssessmentRequestIdResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionThreatAssessmentRequestIdResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionThreatAssessmentRequestIdResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ThreatAssessmentRequestId, ok = input.Parsed["threatAssessmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", input)
	}

	if id.ThreatAssessmentResultId, ok = input.Parsed["threatAssessmentResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", input)
	}

	return nil
}

// ValidateMeInformationProtectionThreatAssessmentRequestIdResultID checks that 'input' can be parsed as a Me Information Protection Threat Assessment Request Id Result ID
func ValidateMeInformationProtectionThreatAssessmentRequestIdResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionThreatAssessmentRequestIdResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Threat Assessment Request Id Result ID
func (id MeInformationProtectionThreatAssessmentRequestIdResultId) ID() string {
	fmtString := "/me/informationProtection/threatAssessmentRequests/%s/results/%s"
	return fmt.Sprintf(fmtString, id.ThreatAssessmentRequestId, id.ThreatAssessmentResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Threat Assessment Request Id Result ID
func (id MeInformationProtectionThreatAssessmentRequestIdResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("threatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestId"),
		resourceids.StaticSegment("results", "results", "results"),
		resourceids.UserSpecifiedSegment("threatAssessmentResultId", "threatAssessmentResultId"),
	}
}

// String returns a human-readable description of this Me Information Protection Threat Assessment Request Id Result ID
func (id MeInformationProtectionThreatAssessmentRequestIdResultId) String() string {
	components := []string{
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
		fmt.Sprintf("Threat Assessment Result: %q", id.ThreatAssessmentResultId),
	}
	return fmt.Sprintf("Me Information Protection Threat Assessment Request Id Result (%s)", strings.Join(components, "\n"))
}
