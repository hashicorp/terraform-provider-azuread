package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{}

// UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId is a struct representing the Resource ID for a User Id Managed Device Id Security Baseline State Id Setting State
type UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId struct {
	UserId                         string
	ManagedDeviceId                string
	SecurityBaselineStateId        string
	SecurityBaselineSettingStateId string
}

// NewUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID returns a new UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId struct
func NewUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(userId string, managedDeviceId string, securityBaselineStateId string, securityBaselineSettingStateId string) UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId {
	return UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{
		UserId:                         userId,
		ManagedDeviceId:                managedDeviceId,
		SecurityBaselineStateId:        securityBaselineStateId,
		SecurityBaselineSettingStateId: securityBaselineSettingStateId,
	}
}

// ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID parses 'input' into a UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId
func ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(input string) (*UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(input string) (*UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.SecurityBaselineStateId, ok = input.Parsed["securityBaselineStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", input)
	}

	if id.SecurityBaselineSettingStateId, ok = input.Parsed["securityBaselineSettingStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineSettingStateId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID checks that 'input' can be parsed as a User Id Managed Device Id Security Baseline State Id Setting State ID
func ValidateUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Security Baseline State Id Setting State ID
func (id UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/securityBaselineStates/%s/settingStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.SecurityBaselineStateId, id.SecurityBaselineSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Security Baseline State Id Setting State ID
func (id UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
		resourceids.StaticSegment("settingStates", "settingStates", "settingStates"),
		resourceids.UserSpecifiedSegment("securityBaselineSettingStateId", "securityBaselineSettingStateId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Security Baseline State Id Setting State ID
func (id UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
		fmt.Sprintf("Security Baseline Setting State: %q", id.SecurityBaselineSettingStateId),
	}
	return fmt.Sprintf("User Id Managed Device Id Security Baseline State Id Setting State (%s)", strings.Join(components, "\n"))
}
