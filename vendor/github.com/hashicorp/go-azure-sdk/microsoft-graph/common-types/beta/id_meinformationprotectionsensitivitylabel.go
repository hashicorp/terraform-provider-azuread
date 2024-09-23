package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionSensitivityLabelId{}

// MeInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a Me Information Protection Sensitivity Label
type MeInformationProtectionSensitivityLabelId struct {
	SensitivityLabelId string
}

// NewMeInformationProtectionSensitivityLabelID returns a new MeInformationProtectionSensitivityLabelId struct
func NewMeInformationProtectionSensitivityLabelID(sensitivityLabelId string) MeInformationProtectionSensitivityLabelId {
	return MeInformationProtectionSensitivityLabelId{
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseMeInformationProtectionSensitivityLabelID parses 'input' into a MeInformationProtectionSensitivityLabelId
func ParseMeInformationProtectionSensitivityLabelID(input string) (*MeInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionSensitivityLabelIDInsensitively(input string) (*MeInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateMeInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a Me Information Protection Sensitivity Label ID
func ValidateMeInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Sensitivity Label ID
func (id MeInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/me/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Sensitivity Label ID
func (id MeInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this Me Information Protection Sensitivity Label ID
func (id MeInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("Me Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
