package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCommunicationCallSettingDelegatorId{}

// UserIdCommunicationCallSettingDelegatorId is a struct representing the Resource ID for a User Id Communication Call Setting Delegator
type UserIdCommunicationCallSettingDelegatorId struct {
	UserId               string
	DelegationSettingsId string
}

// NewUserIdCommunicationCallSettingDelegatorID returns a new UserIdCommunicationCallSettingDelegatorId struct
func NewUserIdCommunicationCallSettingDelegatorID(userId string, delegationSettingsId string) UserIdCommunicationCallSettingDelegatorId {
	return UserIdCommunicationCallSettingDelegatorId{
		UserId:               userId,
		DelegationSettingsId: delegationSettingsId,
	}
}

// ParseUserIdCommunicationCallSettingDelegatorID parses 'input' into a UserIdCommunicationCallSettingDelegatorId
func ParseUserIdCommunicationCallSettingDelegatorID(input string) (*UserIdCommunicationCallSettingDelegatorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCommunicationCallSettingDelegatorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCommunicationCallSettingDelegatorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCommunicationCallSettingDelegatorIDInsensitively parses 'input' case-insensitively into a UserIdCommunicationCallSettingDelegatorId
// note: this method should only be used for API response data and not user input
func ParseUserIdCommunicationCallSettingDelegatorIDInsensitively(input string) (*UserIdCommunicationCallSettingDelegatorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCommunicationCallSettingDelegatorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCommunicationCallSettingDelegatorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCommunicationCallSettingDelegatorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DelegationSettingsId, ok = input.Parsed["delegationSettingsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "delegationSettingsId", input)
	}

	return nil
}

// ValidateUserIdCommunicationCallSettingDelegatorID checks that 'input' can be parsed as a User Id Communication Call Setting Delegator ID
func ValidateUserIdCommunicationCallSettingDelegatorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCommunicationCallSettingDelegatorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Communication Call Setting Delegator ID
func (id UserIdCommunicationCallSettingDelegatorId) ID() string {
	fmtString := "/users/%s/communications/callSettings/delegators/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DelegationSettingsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Communication Call Setting Delegator ID
func (id UserIdCommunicationCallSettingDelegatorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("communications", "communications", "communications"),
		resourceids.StaticSegment("callSettings", "callSettings", "callSettings"),
		resourceids.StaticSegment("delegators", "delegators", "delegators"),
		resourceids.UserSpecifiedSegment("delegationSettingsId", "delegationSettingsId"),
	}
}

// String returns a human-readable description of this User Id Communication Call Setting Delegator ID
func (id UserIdCommunicationCallSettingDelegatorId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Delegation Settings: %q", id.DelegationSettingsId),
	}
	return fmt.Sprintf("User Id Communication Call Setting Delegator (%s)", strings.Join(components, "\n"))
}
