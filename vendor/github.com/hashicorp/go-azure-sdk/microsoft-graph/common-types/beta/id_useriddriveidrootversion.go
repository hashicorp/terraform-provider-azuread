package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootVersionId{}

// UserIdDriveIdRootVersionId is a struct representing the Resource ID for a User Id Drive Id Root Version
type UserIdDriveIdRootVersionId struct {
	UserId             string
	DriveId            string
	DriveItemVersionId string
}

// NewUserIdDriveIdRootVersionID returns a new UserIdDriveIdRootVersionId struct
func NewUserIdDriveIdRootVersionID(userId string, driveId string, driveItemVersionId string) UserIdDriveIdRootVersionId {
	return UserIdDriveIdRootVersionId{
		UserId:             userId,
		DriveId:            driveId,
		DriveItemVersionId: driveItemVersionId,
	}
}

// ParseUserIdDriveIdRootVersionID parses 'input' into a UserIdDriveIdRootVersionId
func ParseUserIdDriveIdRootVersionID(input string) (*UserIdDriveIdRootVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootVersionIDInsensitively(input string) (*UserIdDriveIdRootVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemVersionId, ok = input.Parsed["driveItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootVersionID checks that 'input' can be parsed as a User Id Drive Id Root Version ID
func ValidateUserIdDriveIdRootVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Version ID
func (id UserIdDriveIdRootVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/versions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Version ID
func (id UserIdDriveIdRootVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("driveItemVersionId", "driveItemVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Version ID
func (id UserIdDriveIdRootVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item Version: %q", id.DriveItemVersionId),
	}
	return fmt.Sprintf("User Id Drive Id Root Version (%s)", strings.Join(components, "\n"))
}
