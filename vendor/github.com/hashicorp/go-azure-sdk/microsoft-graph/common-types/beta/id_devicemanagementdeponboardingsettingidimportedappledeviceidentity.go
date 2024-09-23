package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{}

// DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId is a struct representing the Resource ID for a Device Management Dep Onboarding Setting Id Imported Apple Device Identity
type DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId struct {
	DepOnboardingSettingId        string
	ImportedAppleDeviceIdentityId string
}

// NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID returns a new DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId struct
func NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(depOnboardingSettingId string, importedAppleDeviceIdentityId string) DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId {
	return DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
		DepOnboardingSettingId:        depOnboardingSettingId,
		ImportedAppleDeviceIdentityId: importedAppleDeviceIdentityId,
	}
}

// ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID parses 'input' into a DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId
func ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(input string) (*DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityIDInsensitively parses 'input' case-insensitively into a DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityIDInsensitively(input string) (*DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DepOnboardingSettingId, ok = input.Parsed["depOnboardingSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "depOnboardingSettingId", input)
	}

	if id.ImportedAppleDeviceIdentityId, ok = input.Parsed["importedAppleDeviceIdentityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "importedAppleDeviceIdentityId", input)
	}

	return nil
}

// ValidateDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID checks that 'input' can be parsed as a Device Management Dep Onboarding Setting Id Imported Apple Device Identity ID
func ValidateDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Dep Onboarding Setting Id Imported Apple Device Identity ID
func (id DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId) ID() string {
	fmtString := "/deviceManagement/depOnboardingSettings/%s/importedAppleDeviceIdentities/%s"
	return fmt.Sprintf(fmtString, id.DepOnboardingSettingId, id.ImportedAppleDeviceIdentityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Dep Onboarding Setting Id Imported Apple Device Identity ID
func (id DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("depOnboardingSettings", "depOnboardingSettings", "depOnboardingSettings"),
		resourceids.UserSpecifiedSegment("depOnboardingSettingId", "depOnboardingSettingId"),
		resourceids.StaticSegment("importedAppleDeviceIdentities", "importedAppleDeviceIdentities", "importedAppleDeviceIdentities"),
		resourceids.UserSpecifiedSegment("importedAppleDeviceIdentityId", "importedAppleDeviceIdentityId"),
	}
}

// String returns a human-readable description of this Device Management Dep Onboarding Setting Id Imported Apple Device Identity ID
func (id DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId) String() string {
	components := []string{
		fmt.Sprintf("Dep Onboarding Setting: %q", id.DepOnboardingSettingId),
		fmt.Sprintf("Imported Apple Device Identity: %q", id.ImportedAppleDeviceIdentityId),
	}
	return fmt.Sprintf("Device Management Dep Onboarding Setting Id Imported Apple Device Identity (%s)", strings.Join(components, "\n"))
}
