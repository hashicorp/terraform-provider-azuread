package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdExtensionId{}

// UserIdDriveIdItemIdExtensionId is a struct representing the Resource ID for a User Id Drive Id Item Id Extension
type UserIdDriveIdItemIdExtensionId struct {
	UserId      string
	DriveId     string
	DriveItemId string
	ExtensionId string
}

// NewUserIdDriveIdItemIdExtensionID returns a new UserIdDriveIdItemIdExtensionId struct
func NewUserIdDriveIdItemIdExtensionID(userId string, driveId string, driveItemId string, extensionId string) UserIdDriveIdItemIdExtensionId {
	return UserIdDriveIdItemIdExtensionId{
		UserId:      userId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdDriveIdItemIdExtensionID parses 'input' into a UserIdDriveIdItemIdExtensionId
func ParseUserIdDriveIdItemIdExtensionID(input string) (*UserIdDriveIdItemIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdExtensionIDInsensitively(input string) (*UserIdDriveIdItemIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdExtensionID checks that 'input' can be parsed as a User Id Drive Id Item Id Extension ID
func ValidateUserIdDriveIdItemIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Extension ID
func (id UserIdDriveIdItemIdExtensionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Extension ID
func (id UserIdDriveIdItemIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Extension ID
func (id UserIdDriveIdItemIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Extension (%s)", strings.Join(components, "\n"))
}
