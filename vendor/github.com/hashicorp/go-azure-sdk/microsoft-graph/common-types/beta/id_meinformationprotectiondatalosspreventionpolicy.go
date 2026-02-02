package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionDataLossPreventionPolicyId{}

// MeInformationProtectionDataLossPreventionPolicyId is a struct representing the Resource ID for a Me Information Protection Data Loss Prevention Policy
type MeInformationProtectionDataLossPreventionPolicyId struct {
	DataLossPreventionPolicyId string
}

// NewMeInformationProtectionDataLossPreventionPolicyID returns a new MeInformationProtectionDataLossPreventionPolicyId struct
func NewMeInformationProtectionDataLossPreventionPolicyID(dataLossPreventionPolicyId string) MeInformationProtectionDataLossPreventionPolicyId {
	return MeInformationProtectionDataLossPreventionPolicyId{
		DataLossPreventionPolicyId: dataLossPreventionPolicyId,
	}
}

// ParseMeInformationProtectionDataLossPreventionPolicyID parses 'input' into a MeInformationProtectionDataLossPreventionPolicyId
func ParseMeInformationProtectionDataLossPreventionPolicyID(input string) (*MeInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionDataLossPreventionPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionDataLossPreventionPolicyIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionDataLossPreventionPolicyId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionDataLossPreventionPolicyIDInsensitively(input string) (*MeInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionDataLossPreventionPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionDataLossPreventionPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DataLossPreventionPolicyId, ok = input.Parsed["dataLossPreventionPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataLossPreventionPolicyId", input)
	}

	return nil
}

// ValidateMeInformationProtectionDataLossPreventionPolicyID checks that 'input' can be parsed as a Me Information Protection Data Loss Prevention Policy ID
func ValidateMeInformationProtectionDataLossPreventionPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionDataLossPreventionPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Data Loss Prevention Policy ID
func (id MeInformationProtectionDataLossPreventionPolicyId) ID() string {
	fmtString := "/me/informationProtection/dataLossPreventionPolicies/%s"
	return fmt.Sprintf(fmtString, id.DataLossPreventionPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Data Loss Prevention Policy ID
func (id MeInformationProtectionDataLossPreventionPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("dataLossPreventionPolicies", "dataLossPreventionPolicies", "dataLossPreventionPolicies"),
		resourceids.UserSpecifiedSegment("dataLossPreventionPolicyId", "dataLossPreventionPolicyId"),
	}
}

// String returns a human-readable description of this Me Information Protection Data Loss Prevention Policy ID
func (id MeInformationProtectionDataLossPreventionPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Data Loss Prevention Policy: %q", id.DataLossPreventionPolicyId),
	}
	return fmt.Sprintf("Me Information Protection Data Loss Prevention Policy (%s)", strings.Join(components, "\n"))
}
