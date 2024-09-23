package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdDetectedAppId{}

// MeManagedDeviceIdDetectedAppId is a struct representing the Resource ID for a Me Managed Device Id Detected App
type MeManagedDeviceIdDetectedAppId struct {
	ManagedDeviceId string
	DetectedAppId   string
}

// NewMeManagedDeviceIdDetectedAppID returns a new MeManagedDeviceIdDetectedAppId struct
func NewMeManagedDeviceIdDetectedAppID(managedDeviceId string, detectedAppId string) MeManagedDeviceIdDetectedAppId {
	return MeManagedDeviceIdDetectedAppId{
		ManagedDeviceId: managedDeviceId,
		DetectedAppId:   detectedAppId,
	}
}

// ParseMeManagedDeviceIdDetectedAppID parses 'input' into a MeManagedDeviceIdDetectedAppId
func ParseMeManagedDeviceIdDetectedAppID(input string) (*MeManagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdDetectedAppIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdDetectedAppIDInsensitively(input string) (*MeManagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdDetectedAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DetectedAppId, ok = input.Parsed["detectedAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdDetectedAppID checks that 'input' can be parsed as a Me Managed Device Id Detected App ID
func ValidateMeManagedDeviceIdDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Detected App ID
func (id MeManagedDeviceIdDetectedAppId) ID() string {
	fmtString := "/me/managedDevices/%s/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Detected App ID
func (id MeManagedDeviceIdDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("detectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Detected App ID
func (id MeManagedDeviceIdDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("Me Managed Device Id Detected App (%s)", strings.Join(components, "\n"))
}
