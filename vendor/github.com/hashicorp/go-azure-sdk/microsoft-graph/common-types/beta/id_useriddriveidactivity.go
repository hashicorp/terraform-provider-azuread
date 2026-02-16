package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdActivityId{}

// UserIdDriveIdActivityId is a struct representing the Resource ID for a User Id Drive Id Activity
type UserIdDriveIdActivityId struct {
	UserId            string
	DriveId           string
	ItemActivityOLDId string
}

// NewUserIdDriveIdActivityID returns a new UserIdDriveIdActivityId struct
func NewUserIdDriveIdActivityID(userId string, driveId string, itemActivityOLDId string) UserIdDriveIdActivityId {
	return UserIdDriveIdActivityId{
		UserId:            userId,
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseUserIdDriveIdActivityID parses 'input' into a UserIdDriveIdActivityId
func ParseUserIdDriveIdActivityID(input string) (*UserIdDriveIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdActivityIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdActivityIDInsensitively(input string) (*UserIdDriveIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateUserIdDriveIdActivityID checks that 'input' can be parsed as a User Id Drive Id Activity ID
func ValidateUserIdDriveIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Activity ID
func (id UserIdDriveIdActivityId) ID() string {
	fmtString := "/users/%s/drives/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Activity ID
func (id UserIdDriveIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Activity ID
func (id UserIdDriveIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("User Id Drive Id Activity (%s)", strings.Join(components, "\n"))
}
