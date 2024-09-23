package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListColumnId{}

// UserIdDriveIdListColumnId is a struct representing the Resource ID for a User Id Drive Id List Column
type UserIdDriveIdListColumnId struct {
	UserId             string
	DriveId            string
	ColumnDefinitionId string
}

// NewUserIdDriveIdListColumnID returns a new UserIdDriveIdListColumnId struct
func NewUserIdDriveIdListColumnID(userId string, driveId string, columnDefinitionId string) UserIdDriveIdListColumnId {
	return UserIdDriveIdListColumnId{
		UserId:             userId,
		DriveId:            driveId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseUserIdDriveIdListColumnID parses 'input' into a UserIdDriveIdListColumnId
func ParseUserIdDriveIdListColumnID(input string) (*UserIdDriveIdListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListColumnIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListColumnId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListColumnIDInsensitively(input string) (*UserIdDriveIdListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListColumnID checks that 'input' can be parsed as a User Id Drive Id List Column ID
func ValidateUserIdDriveIdListColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Column ID
func (id UserIdDriveIdListColumnId) ID() string {
	fmtString := "/users/%s/drives/%s/list/columns/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Column ID
func (id UserIdDriveIdListColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Column ID
func (id UserIdDriveIdListColumnId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("User Id Drive Id List Column (%s)", strings.Join(components, "\n"))
}
