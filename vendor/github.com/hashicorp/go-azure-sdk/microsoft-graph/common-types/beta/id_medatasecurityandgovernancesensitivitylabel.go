package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDataSecurityAndGovernanceSensitivityLabelId{}

// MeDataSecurityAndGovernanceSensitivityLabelId is a struct representing the Resource ID for a Me Data Security And Governance Sensitivity Label
type MeDataSecurityAndGovernanceSensitivityLabelId struct {
	SensitivityLabelId string
}

// NewMeDataSecurityAndGovernanceSensitivityLabelID returns a new MeDataSecurityAndGovernanceSensitivityLabelId struct
func NewMeDataSecurityAndGovernanceSensitivityLabelID(sensitivityLabelId string) MeDataSecurityAndGovernanceSensitivityLabelId {
	return MeDataSecurityAndGovernanceSensitivityLabelId{
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseMeDataSecurityAndGovernanceSensitivityLabelID parses 'input' into a MeDataSecurityAndGovernanceSensitivityLabelId
func ParseMeDataSecurityAndGovernanceSensitivityLabelID(input string) (*MeDataSecurityAndGovernanceSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDataSecurityAndGovernanceSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDataSecurityAndGovernanceSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDataSecurityAndGovernanceSensitivityLabelIDInsensitively parses 'input' case-insensitively into a MeDataSecurityAndGovernanceSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseMeDataSecurityAndGovernanceSensitivityLabelIDInsensitively(input string) (*MeDataSecurityAndGovernanceSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDataSecurityAndGovernanceSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDataSecurityAndGovernanceSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDataSecurityAndGovernanceSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateMeDataSecurityAndGovernanceSensitivityLabelID checks that 'input' can be parsed as a Me Data Security And Governance Sensitivity Label ID
func ValidateMeDataSecurityAndGovernanceSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDataSecurityAndGovernanceSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Data Security And Governance Sensitivity Label ID
func (id MeDataSecurityAndGovernanceSensitivityLabelId) ID() string {
	fmtString := "/me/dataSecurityAndGovernance/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Data Security And Governance Sensitivity Label ID
func (id MeDataSecurityAndGovernanceSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("dataSecurityAndGovernance", "dataSecurityAndGovernance", "dataSecurityAndGovernance"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this Me Data Security And Governance Sensitivity Label ID
func (id MeDataSecurityAndGovernanceSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("Me Data Security And Governance Sensitivity Label (%s)", strings.Join(components, "\n"))
}
