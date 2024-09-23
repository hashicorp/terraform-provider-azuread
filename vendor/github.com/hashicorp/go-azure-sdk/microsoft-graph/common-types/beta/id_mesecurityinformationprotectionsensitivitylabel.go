package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSecurityInformationProtectionSensitivityLabelId{}

// MeSecurityInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a Me Security Information Protection Sensitivity Label
type MeSecurityInformationProtectionSensitivityLabelId struct {
	SensitivityLabelId string
}

// NewMeSecurityInformationProtectionSensitivityLabelID returns a new MeSecurityInformationProtectionSensitivityLabelId struct
func NewMeSecurityInformationProtectionSensitivityLabelID(sensitivityLabelId string) MeSecurityInformationProtectionSensitivityLabelId {
	return MeSecurityInformationProtectionSensitivityLabelId{
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseMeSecurityInformationProtectionSensitivityLabelID parses 'input' into a MeSecurityInformationProtectionSensitivityLabelId
func ParseMeSecurityInformationProtectionSensitivityLabelID(input string) (*MeSecurityInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSecurityInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSecurityInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeSecurityInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a MeSecurityInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseMeSecurityInformationProtectionSensitivityLabelIDInsensitively(input string) (*MeSecurityInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSecurityInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSecurityInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeSecurityInformationProtectionSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateMeSecurityInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a Me Security Information Protection Sensitivity Label ID
func ValidateMeSecurityInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeSecurityInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Security Information Protection Sensitivity Label ID
func (id MeSecurityInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/me/security/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Security Information Protection Sensitivity Label ID
func (id MeSecurityInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("security", "security", "security"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this Me Security Information Protection Sensitivity Label ID
func (id MeSecurityInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("Me Security Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
