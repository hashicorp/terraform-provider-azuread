package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootExtensionId{}

// UserIdDriveIdRootExtensionId is a struct representing the Resource ID for a User Id Drive Id Root Extension
type UserIdDriveIdRootExtensionId struct {
	UserId      string
	DriveId     string
	ExtensionId string
}

// NewUserIdDriveIdRootExtensionID returns a new UserIdDriveIdRootExtensionId struct
func NewUserIdDriveIdRootExtensionID(userId string, driveId string, extensionId string) UserIdDriveIdRootExtensionId {
	return UserIdDriveIdRootExtensionId{
		UserId:      userId,
		DriveId:     driveId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdDriveIdRootExtensionID parses 'input' into a UserIdDriveIdRootExtensionId
func ParseUserIdDriveIdRootExtensionID(input string) (*UserIdDriveIdRootExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootExtensionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootExtensionIDInsensitively(input string) (*UserIdDriveIdRootExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootExtensionID checks that 'input' can be parsed as a User Id Drive Id Root Extension ID
func ValidateUserIdDriveIdRootExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Extension ID
func (id UserIdDriveIdRootExtensionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Extension ID
func (id UserIdDriveIdRootExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Extension ID
func (id UserIdDriveIdRootExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Drive Id Root Extension (%s)", strings.Join(components, "\n"))
}
