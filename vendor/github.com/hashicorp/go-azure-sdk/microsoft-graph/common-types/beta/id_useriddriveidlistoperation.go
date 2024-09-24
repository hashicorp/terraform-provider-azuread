package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListOperationId{}

// UserIdDriveIdListOperationId is a struct representing the Resource ID for a User Id Drive Id List Operation
type UserIdDriveIdListOperationId struct {
	UserId                     string
	DriveId                    string
	RichLongRunningOperationId string
}

// NewUserIdDriveIdListOperationID returns a new UserIdDriveIdListOperationId struct
func NewUserIdDriveIdListOperationID(userId string, driveId string, richLongRunningOperationId string) UserIdDriveIdListOperationId {
	return UserIdDriveIdListOperationId{
		UserId:                     userId,
		DriveId:                    driveId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseUserIdDriveIdListOperationID parses 'input' into a UserIdDriveIdListOperationId
func ParseUserIdDriveIdListOperationID(input string) (*UserIdDriveIdListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListOperationIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListOperationIDInsensitively(input string) (*UserIdDriveIdListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.RichLongRunningOperationId, ok = input.Parsed["richLongRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListOperationID checks that 'input' can be parsed as a User Id Drive Id List Operation ID
func ValidateUserIdDriveIdListOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Operation ID
func (id UserIdDriveIdListOperationId) ID() string {
	fmtString := "/users/%s/drives/%s/list/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Operation ID
func (id UserIdDriveIdListOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Operation ID
func (id UserIdDriveIdListOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("User Id Drive Id List Operation (%s)", strings.Join(components, "\n"))
}
