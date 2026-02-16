package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}

// MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId is a struct representing the Resource ID for a Me Data Security And Governance Sensitivity Label Id Sublabel
type MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId struct {
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID returns a new MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId struct
func NewMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(sensitivityLabelId string, sensitivityLabelId1 string) MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId {
	return MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID parses 'input' into a MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId
func ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(input string) (*MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively parses 'input' case-insensitively into a MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId
// note: this method should only be used for API response data and not user input
func ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively(input string) (*MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	if id.SensitivityLabelId1, ok = input.Parsed["sensitivityLabelId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", input)
	}

	return nil
}

// ValidateMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID checks that 'input' can be parsed as a Me Data Security And Governance Sensitivity Label Id Sublabel ID
func ValidateMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Data Security And Governance Sensitivity Label Id Sublabel ID
func (id MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId) ID() string {
	fmtString := "/me/dataSecurityAndGovernance/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Data Security And Governance Sensitivity Label Id Sublabel ID
func (id MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("dataSecurityAndGovernance", "dataSecurityAndGovernance", "dataSecurityAndGovernance"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
		resourceids.StaticSegment("sublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1"),
	}
}

// String returns a human-readable description of this Me Data Security And Governance Sensitivity Label Id Sublabel ID
func (id MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId) String() string {
	components := []string{
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("Me Data Security And Governance Sensitivity Label Id Sublabel (%s)", strings.Join(components, "\n"))
}
