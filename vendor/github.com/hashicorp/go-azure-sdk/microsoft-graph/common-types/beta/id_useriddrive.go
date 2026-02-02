package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveId{}

// UserIdDriveId is a struct representing the Resource ID for a User Id Drive
type UserIdDriveId struct {
	UserId  string
	DriveId string
}

// NewUserIdDriveID returns a new UserIdDriveId struct
func NewUserIdDriveID(userId string, driveId string) UserIdDriveId {
	return UserIdDriveId{
		UserId:  userId,
		DriveId: driveId,
	}
}

// ParseUserIdDriveID parses 'input' into a UserIdDriveId
func ParseUserIdDriveID(input string) (*UserIdDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIDInsensitively parses 'input' case-insensitively into a UserIdDriveId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIDInsensitively(input string) (*UserIdDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	return nil
}

// ValidateUserIdDriveID checks that 'input' can be parsed as a User Id Drive ID
func ValidateUserIdDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive ID
func (id UserIdDriveId) ID() string {
	fmtString := "/users/%s/drives/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive ID
func (id UserIdDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
	}
}

// String returns a human-readable description of this User Id Drive ID
func (id UserIdDriveId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("User Id Drive (%s)", strings.Join(components, "\n"))
}
