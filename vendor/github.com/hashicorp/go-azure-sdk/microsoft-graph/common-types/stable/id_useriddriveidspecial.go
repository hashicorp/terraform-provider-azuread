package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdSpecialId{}

// UserIdDriveIdSpecialId is a struct representing the Resource ID for a User Id Drive Id Special
type UserIdDriveIdSpecialId struct {
	UserId      string
	DriveId     string
	DriveItemId string
}

// NewUserIdDriveIdSpecialID returns a new UserIdDriveIdSpecialId struct
func NewUserIdDriveIdSpecialID(userId string, driveId string, driveItemId string) UserIdDriveIdSpecialId {
	return UserIdDriveIdSpecialId{
		UserId:      userId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseUserIdDriveIdSpecialID parses 'input' into a UserIdDriveIdSpecialId
func ParseUserIdDriveIdSpecialID(input string) (*UserIdDriveIdSpecialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdSpecialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdSpecialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdSpecialIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdSpecialId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdSpecialIDInsensitively(input string) (*UserIdDriveIdSpecialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdSpecialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdSpecialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdSpecialId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdSpecialID checks that 'input' can be parsed as a User Id Drive Id Special ID
func ValidateUserIdDriveIdSpecialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdSpecialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Special ID
func (id UserIdDriveIdSpecialId) ID() string {
	fmtString := "/users/%s/drives/%s/special/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Special ID
func (id UserIdDriveIdSpecialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("special", "special", "special"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Special ID
func (id UserIdDriveIdSpecialId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("User Id Drive Id Special (%s)", strings.Join(components, "\n"))
}
