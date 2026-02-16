package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdBundleId{}

// UserIdDriveIdBundleId is a struct representing the Resource ID for a User Id Drive Id Bundle
type UserIdDriveIdBundleId struct {
	UserId      string
	DriveId     string
	DriveItemId string
}

// NewUserIdDriveIdBundleID returns a new UserIdDriveIdBundleId struct
func NewUserIdDriveIdBundleID(userId string, driveId string, driveItemId string) UserIdDriveIdBundleId {
	return UserIdDriveIdBundleId{
		UserId:      userId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseUserIdDriveIdBundleID parses 'input' into a UserIdDriveIdBundleId
func ParseUserIdDriveIdBundleID(input string) (*UserIdDriveIdBundleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdBundleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdBundleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdBundleIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdBundleId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdBundleIDInsensitively(input string) (*UserIdDriveIdBundleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdBundleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdBundleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdBundleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	return nil
}

// ValidateUserIdDriveIdBundleID checks that 'input' can be parsed as a User Id Drive Id Bundle ID
func ValidateUserIdDriveIdBundleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdBundleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Bundle ID
func (id UserIdDriveIdBundleId) ID() string {
	fmtString := "/users/%s/drives/%s/bundles/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Bundle ID
func (id UserIdDriveIdBundleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("bundles", "bundles", "bundles"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Bundle ID
func (id UserIdDriveIdBundleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("User Id Drive Id Bundle (%s)", strings.Join(components, "\n"))
}
