package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListContentTypeId{}

// UserIdDriveIdListContentTypeId is a struct representing the Resource ID for a User Id Drive Id List Content Type
type UserIdDriveIdListContentTypeId struct {
	UserId        string
	DriveId       string
	ContentTypeId string
}

// NewUserIdDriveIdListContentTypeID returns a new UserIdDriveIdListContentTypeId struct
func NewUserIdDriveIdListContentTypeID(userId string, driveId string, contentTypeId string) UserIdDriveIdListContentTypeId {
	return UserIdDriveIdListContentTypeId{
		UserId:        userId,
		DriveId:       driveId,
		ContentTypeId: contentTypeId,
	}
}

// ParseUserIdDriveIdListContentTypeID parses 'input' into a UserIdDriveIdListContentTypeId
func ParseUserIdDriveIdListContentTypeID(input string) (*UserIdDriveIdListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListContentTypeIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListContentTypeId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListContentTypeIDInsensitively(input string) (*UserIdDriveIdListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListContentTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListContentTypeID checks that 'input' can be parsed as a User Id Drive Id List Content Type ID
func ValidateUserIdDriveIdListContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Content Type ID
func (id UserIdDriveIdListContentTypeId) ID() string {
	fmtString := "/users/%s/drives/%s/list/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Content Type ID
func (id UserIdDriveIdListContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Content Type ID
func (id UserIdDriveIdListContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("User Id Drive Id List Content Type (%s)", strings.Join(components, "\n"))
}
