package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCommunicationCallSettingDelegateId{}

// UserIdCommunicationCallSettingDelegateId is a struct representing the Resource ID for a User Id Communication Call Setting Delegate
type UserIdCommunicationCallSettingDelegateId struct {
	UserId               string
	DelegationSettingsId string
}

// NewUserIdCommunicationCallSettingDelegateID returns a new UserIdCommunicationCallSettingDelegateId struct
func NewUserIdCommunicationCallSettingDelegateID(userId string, delegationSettingsId string) UserIdCommunicationCallSettingDelegateId {
	return UserIdCommunicationCallSettingDelegateId{
		UserId:               userId,
		DelegationSettingsId: delegationSettingsId,
	}
}

// ParseUserIdCommunicationCallSettingDelegateID parses 'input' into a UserIdCommunicationCallSettingDelegateId
func ParseUserIdCommunicationCallSettingDelegateID(input string) (*UserIdCommunicationCallSettingDelegateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCommunicationCallSettingDelegateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCommunicationCallSettingDelegateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCommunicationCallSettingDelegateIDInsensitively parses 'input' case-insensitively into a UserIdCommunicationCallSettingDelegateId
// note: this method should only be used for API response data and not user input
func ParseUserIdCommunicationCallSettingDelegateIDInsensitively(input string) (*UserIdCommunicationCallSettingDelegateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCommunicationCallSettingDelegateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCommunicationCallSettingDelegateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCommunicationCallSettingDelegateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DelegationSettingsId, ok = input.Parsed["delegationSettingsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "delegationSettingsId", input)
	}

	return nil
}

// ValidateUserIdCommunicationCallSettingDelegateID checks that 'input' can be parsed as a User Id Communication Call Setting Delegate ID
func ValidateUserIdCommunicationCallSettingDelegateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCommunicationCallSettingDelegateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Communication Call Setting Delegate ID
func (id UserIdCommunicationCallSettingDelegateId) ID() string {
	fmtString := "/users/%s/communications/callSettings/delegates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DelegationSettingsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Communication Call Setting Delegate ID
func (id UserIdCommunicationCallSettingDelegateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("communications", "communications", "communications"),
		resourceids.StaticSegment("callSettings", "callSettings", "callSettings"),
		resourceids.StaticSegment("delegates", "delegates", "delegates"),
		resourceids.UserSpecifiedSegment("delegationSettingsId", "delegationSettingsId"),
	}
}

// String returns a human-readable description of this User Id Communication Call Setting Delegate ID
func (id UserIdCommunicationCallSettingDelegateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Delegation Settings: %q", id.DelegationSettingsId),
	}
	return fmt.Sprintf("User Id Communication Call Setting Delegate (%s)", strings.Join(components, "\n"))
}
