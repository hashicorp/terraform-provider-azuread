package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdVersionId{}

// UserIdDriveIdItemIdVersionId is a struct representing the Resource ID for a User Id Drive Id Item Id Version
type UserIdDriveIdItemIdVersionId struct {
	UserId             string
	DriveId            string
	DriveItemId        string
	DriveItemVersionId string
}

// NewUserIdDriveIdItemIdVersionID returns a new UserIdDriveIdItemIdVersionId struct
func NewUserIdDriveIdItemIdVersionID(userId string, driveId string, driveItemId string, driveItemVersionId string) UserIdDriveIdItemIdVersionId {
	return UserIdDriveIdItemIdVersionId{
		UserId:             userId,
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		DriveItemVersionId: driveItemVersionId,
	}
}

// ParseUserIdDriveIdItemIdVersionID parses 'input' into a UserIdDriveIdItemIdVersionId
func ParseUserIdDriveIdItemIdVersionID(input string) (*UserIdDriveIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdVersionIDInsensitively(input string) (*UserIdDriveIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DriveItemVersionId, ok = input.Parsed["driveItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdVersionID checks that 'input' can be parsed as a User Id Drive Id Item Id Version ID
func ValidateUserIdDriveIdItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Version ID
func (id UserIdDriveIdItemIdVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.DriveItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Version ID
func (id UserIdDriveIdItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("driveItemVersionId", "driveItemVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Version ID
func (id UserIdDriveIdItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Drive Item Version: %q", id.DriveItemVersionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Version (%s)", strings.Join(components, "\n"))
}
