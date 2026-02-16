package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCommunicationCallSettingDelegateId{}

// MeCommunicationCallSettingDelegateId is a struct representing the Resource ID for a Me Communication Call Setting Delegate
type MeCommunicationCallSettingDelegateId struct {
	DelegationSettingsId string
}

// NewMeCommunicationCallSettingDelegateID returns a new MeCommunicationCallSettingDelegateId struct
func NewMeCommunicationCallSettingDelegateID(delegationSettingsId string) MeCommunicationCallSettingDelegateId {
	return MeCommunicationCallSettingDelegateId{
		DelegationSettingsId: delegationSettingsId,
	}
}

// ParseMeCommunicationCallSettingDelegateID parses 'input' into a MeCommunicationCallSettingDelegateId
func ParseMeCommunicationCallSettingDelegateID(input string) (*MeCommunicationCallSettingDelegateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCommunicationCallSettingDelegateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCommunicationCallSettingDelegateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCommunicationCallSettingDelegateIDInsensitively parses 'input' case-insensitively into a MeCommunicationCallSettingDelegateId
// note: this method should only be used for API response data and not user input
func ParseMeCommunicationCallSettingDelegateIDInsensitively(input string) (*MeCommunicationCallSettingDelegateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCommunicationCallSettingDelegateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCommunicationCallSettingDelegateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCommunicationCallSettingDelegateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DelegationSettingsId, ok = input.Parsed["delegationSettingsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "delegationSettingsId", input)
	}

	return nil
}

// ValidateMeCommunicationCallSettingDelegateID checks that 'input' can be parsed as a Me Communication Call Setting Delegate ID
func ValidateMeCommunicationCallSettingDelegateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCommunicationCallSettingDelegateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Communication Call Setting Delegate ID
func (id MeCommunicationCallSettingDelegateId) ID() string {
	fmtString := "/me/communications/callSettings/delegates/%s"
	return fmt.Sprintf(fmtString, id.DelegationSettingsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Communication Call Setting Delegate ID
func (id MeCommunicationCallSettingDelegateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("communications", "communications", "communications"),
		resourceids.StaticSegment("callSettings", "callSettings", "callSettings"),
		resourceids.StaticSegment("delegates", "delegates", "delegates"),
		resourceids.UserSpecifiedSegment("delegationSettingsId", "delegationSettingsId"),
	}
}

// String returns a human-readable description of this Me Communication Call Setting Delegate ID
func (id MeCommunicationCallSettingDelegateId) String() string {
	components := []string{
		fmt.Sprintf("Delegation Settings: %q", id.DelegationSettingsId),
	}
	return fmt.Sprintf("Me Communication Call Setting Delegate (%s)", strings.Join(components, "\n"))
}
