package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCommunicationCallSettingDelegatorId{}

// MeCommunicationCallSettingDelegatorId is a struct representing the Resource ID for a Me Communication Call Setting Delegator
type MeCommunicationCallSettingDelegatorId struct {
	DelegationSettingsId string
}

// NewMeCommunicationCallSettingDelegatorID returns a new MeCommunicationCallSettingDelegatorId struct
func NewMeCommunicationCallSettingDelegatorID(delegationSettingsId string) MeCommunicationCallSettingDelegatorId {
	return MeCommunicationCallSettingDelegatorId{
		DelegationSettingsId: delegationSettingsId,
	}
}

// ParseMeCommunicationCallSettingDelegatorID parses 'input' into a MeCommunicationCallSettingDelegatorId
func ParseMeCommunicationCallSettingDelegatorID(input string) (*MeCommunicationCallSettingDelegatorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCommunicationCallSettingDelegatorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCommunicationCallSettingDelegatorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCommunicationCallSettingDelegatorIDInsensitively parses 'input' case-insensitively into a MeCommunicationCallSettingDelegatorId
// note: this method should only be used for API response data and not user input
func ParseMeCommunicationCallSettingDelegatorIDInsensitively(input string) (*MeCommunicationCallSettingDelegatorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCommunicationCallSettingDelegatorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCommunicationCallSettingDelegatorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCommunicationCallSettingDelegatorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DelegationSettingsId, ok = input.Parsed["delegationSettingsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "delegationSettingsId", input)
	}

	return nil
}

// ValidateMeCommunicationCallSettingDelegatorID checks that 'input' can be parsed as a Me Communication Call Setting Delegator ID
func ValidateMeCommunicationCallSettingDelegatorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCommunicationCallSettingDelegatorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Communication Call Setting Delegator ID
func (id MeCommunicationCallSettingDelegatorId) ID() string {
	fmtString := "/me/communications/callSettings/delegators/%s"
	return fmt.Sprintf(fmtString, id.DelegationSettingsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Communication Call Setting Delegator ID
func (id MeCommunicationCallSettingDelegatorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("communications", "communications", "communications"),
		resourceids.StaticSegment("callSettings", "callSettings", "callSettings"),
		resourceids.StaticSegment("delegators", "delegators", "delegators"),
		resourceids.UserSpecifiedSegment("delegationSettingsId", "delegationSettingsId"),
	}
}

// String returns a human-readable description of this Me Communication Call Setting Delegator ID
func (id MeCommunicationCallSettingDelegatorId) String() string {
	components := []string{
		fmt.Sprintf("Delegation Settings: %q", id.DelegationSettingsId),
	}
	return fmt.Sprintf("Me Communication Call Setting Delegator (%s)", strings.Join(components, "\n"))
}
