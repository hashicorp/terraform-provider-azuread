package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdSecurityBaselineStateId{}

// UserIdManagedDeviceIdSecurityBaselineStateId is a struct representing the Resource ID for a User Id Managed Device Id Security Baseline State
type UserIdManagedDeviceIdSecurityBaselineStateId struct {
	UserId                  string
	ManagedDeviceId         string
	SecurityBaselineStateId string
}

// NewUserIdManagedDeviceIdSecurityBaselineStateID returns a new UserIdManagedDeviceIdSecurityBaselineStateId struct
func NewUserIdManagedDeviceIdSecurityBaselineStateID(userId string, managedDeviceId string, securityBaselineStateId string) UserIdManagedDeviceIdSecurityBaselineStateId {
	return UserIdManagedDeviceIdSecurityBaselineStateId{
		UserId:                  userId,
		ManagedDeviceId:         managedDeviceId,
		SecurityBaselineStateId: securityBaselineStateId,
	}
}

// ParseUserIdManagedDeviceIdSecurityBaselineStateID parses 'input' into a UserIdManagedDeviceIdSecurityBaselineStateId
func ParseUserIdManagedDeviceIdSecurityBaselineStateID(input string) (*UserIdManagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdSecurityBaselineStateIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdSecurityBaselineStateId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdSecurityBaselineStateIDInsensitively(input string) (*UserIdManagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdSecurityBaselineStateId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdManagedDeviceIdSecurityBaselineStateID checks that 'input' can be parsed as a User Id Managed Device Id Security Baseline State ID
func ValidateUserIdManagedDeviceIdSecurityBaselineStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdSecurityBaselineStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Security Baseline State ID
func (id UserIdManagedDeviceIdSecurityBaselineStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/securityBaselineStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.SecurityBaselineStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Security Baseline State ID
func (id UserIdManagedDeviceIdSecurityBaselineStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Security Baseline State ID
func (id UserIdManagedDeviceIdSecurityBaselineStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
	}
	return fmt.Sprintf("User Id Managed Device Id Security Baseline State (%s)", strings.Join(components, "\n"))
}
