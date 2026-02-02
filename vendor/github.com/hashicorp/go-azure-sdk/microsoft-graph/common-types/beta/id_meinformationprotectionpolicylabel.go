package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionPolicyLabelId{}

// MeInformationProtectionPolicyLabelId is a struct representing the Resource ID for a Me Information Protection Policy Label
type MeInformationProtectionPolicyLabelId struct {
	InformationProtectionLabelId string
}

// NewMeInformationProtectionPolicyLabelID returns a new MeInformationProtectionPolicyLabelId struct
func NewMeInformationProtectionPolicyLabelID(informationProtectionLabelId string) MeInformationProtectionPolicyLabelId {
	return MeInformationProtectionPolicyLabelId{
		InformationProtectionLabelId: informationProtectionLabelId,
	}
}

// ParseMeInformationProtectionPolicyLabelID parses 'input' into a MeInformationProtectionPolicyLabelId
func ParseMeInformationProtectionPolicyLabelID(input string) (*MeInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionPolicyLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionPolicyLabelIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionPolicyLabelId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionPolicyLabelIDInsensitively(input string) (*MeInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionPolicyLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionPolicyLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InformationProtectionLabelId, ok = input.Parsed["informationProtectionLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", input)
	}

	return nil
}

// ValidateMeInformationProtectionPolicyLabelID checks that 'input' can be parsed as a Me Information Protection Policy Label ID
func ValidateMeInformationProtectionPolicyLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionPolicyLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Policy Label ID
func (id MeInformationProtectionPolicyLabelId) ID() string {
	fmtString := "/me/informationProtection/policy/labels/%s"
	return fmt.Sprintf(fmtString, id.InformationProtectionLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Policy Label ID
func (id MeInformationProtectionPolicyLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("policy", "policy", "policy"),
		resourceids.StaticSegment("labels", "labels", "labels"),
		resourceids.UserSpecifiedSegment("informationProtectionLabelId", "informationProtectionLabelId"),
	}
}

// String returns a human-readable description of this Me Information Protection Policy Label ID
func (id MeInformationProtectionPolicyLabelId) String() string {
	components := []string{
		fmt.Sprintf("Information Protection Label: %q", id.InformationProtectionLabelId),
	}
	return fmt.Sprintf("Me Information Protection Policy Label (%s)", strings.Join(components, "\n"))
}
