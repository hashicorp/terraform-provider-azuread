package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdDetectedAppId{}

// UserIdManagedDeviceIdDetectedAppId is a struct representing the Resource ID for a User Id Managed Device Id Detected App
type UserIdManagedDeviceIdDetectedAppId struct {
	UserId          string
	ManagedDeviceId string
	DetectedAppId   string
}

// NewUserIdManagedDeviceIdDetectedAppID returns a new UserIdManagedDeviceIdDetectedAppId struct
func NewUserIdManagedDeviceIdDetectedAppID(userId string, managedDeviceId string, detectedAppId string) UserIdManagedDeviceIdDetectedAppId {
	return UserIdManagedDeviceIdDetectedAppId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
		DetectedAppId:   detectedAppId,
	}
}

// ParseUserIdManagedDeviceIdDetectedAppID parses 'input' into a UserIdManagedDeviceIdDetectedAppId
func ParseUserIdManagedDeviceIdDetectedAppID(input string) (*UserIdManagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdDetectedAppIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdDetectedAppIDInsensitively(input string) (*UserIdManagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdDetectedAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DetectedAppId, ok = input.Parsed["detectedAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdDetectedAppID checks that 'input' can be parsed as a User Id Managed Device Id Detected App ID
func ValidateUserIdManagedDeviceIdDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Detected App ID
func (id UserIdManagedDeviceIdDetectedAppId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Detected App ID
func (id UserIdManagedDeviceIdDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("detectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Detected App ID
func (id UserIdManagedDeviceIdDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("User Id Managed Device Id Detected App (%s)", strings.Join(components, "\n"))
}
