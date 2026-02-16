package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionSensitivityLabelIdSublabelId{}

// MeInformationProtectionSensitivityLabelIdSublabelId is a struct representing the Resource ID for a Me Information Protection Sensitivity Label Id Sublabel
type MeInformationProtectionSensitivityLabelIdSublabelId struct {
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewMeInformationProtectionSensitivityLabelIdSublabelID returns a new MeInformationProtectionSensitivityLabelIdSublabelId struct
func NewMeInformationProtectionSensitivityLabelIdSublabelID(sensitivityLabelId string, sensitivityLabelId1 string) MeInformationProtectionSensitivityLabelIdSublabelId {
	return MeInformationProtectionSensitivityLabelIdSublabelId{
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseMeInformationProtectionSensitivityLabelIdSublabelID parses 'input' into a MeInformationProtectionSensitivityLabelIdSublabelId
func ParseMeInformationProtectionSensitivityLabelIdSublabelID(input string) (*MeInformationProtectionSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionSensitivityLabelIdSublabelIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionSensitivityLabelIdSublabelId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionSensitivityLabelIdSublabelIDInsensitively(input string) (*MeInformationProtectionSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionSensitivityLabelIdSublabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	if id.SensitivityLabelId1, ok = input.Parsed["sensitivityLabelId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", input)
	}

	return nil
}

// ValidateMeInformationProtectionSensitivityLabelIdSublabelID checks that 'input' can be parsed as a Me Information Protection Sensitivity Label Id Sublabel ID
func ValidateMeInformationProtectionSensitivityLabelIdSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionSensitivityLabelIdSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Sensitivity Label Id Sublabel ID
func (id MeInformationProtectionSensitivityLabelIdSublabelId) ID() string {
	fmtString := "/me/informationProtection/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Sensitivity Label Id Sublabel ID
func (id MeInformationProtectionSensitivityLabelIdSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
		resourceids.StaticSegment("sublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1"),
	}
}

// String returns a human-readable description of this Me Information Protection Sensitivity Label Id Sublabel ID
func (id MeInformationProtectionSensitivityLabelIdSublabelId) String() string {
	components := []string{
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("Me Information Protection Sensitivity Label Id Sublabel (%s)", strings.Join(components, "\n"))
}
