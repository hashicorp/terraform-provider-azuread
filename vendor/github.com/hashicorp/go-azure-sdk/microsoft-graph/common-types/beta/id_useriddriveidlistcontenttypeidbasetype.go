package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListContentTypeIdBaseTypeId{}

// UserIdDriveIdListContentTypeIdBaseTypeId is a struct representing the Resource ID for a User Id Drive Id List Content Type Id Base Type
type UserIdDriveIdListContentTypeIdBaseTypeId struct {
	UserId         string
	DriveId        string
	ContentTypeId  string
	ContentTypeId1 string
}

// NewUserIdDriveIdListContentTypeIdBaseTypeID returns a new UserIdDriveIdListContentTypeIdBaseTypeId struct
func NewUserIdDriveIdListContentTypeIdBaseTypeID(userId string, driveId string, contentTypeId string, contentTypeId1 string) UserIdDriveIdListContentTypeIdBaseTypeId {
	return UserIdDriveIdListContentTypeIdBaseTypeId{
		UserId:         userId,
		DriveId:        driveId,
		ContentTypeId:  contentTypeId,
		ContentTypeId1: contentTypeId1,
	}
}

// ParseUserIdDriveIdListContentTypeIdBaseTypeID parses 'input' into a UserIdDriveIdListContentTypeIdBaseTypeId
func ParseUserIdDriveIdListContentTypeIdBaseTypeID(input string) (*UserIdDriveIdListContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListContentTypeIdBaseTypeIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListContentTypeIdBaseTypeId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListContentTypeIdBaseTypeIDInsensitively(input string) (*UserIdDriveIdListContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListContentTypeIdBaseTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContentTypeId1, ok = input.Parsed["contentTypeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", input)
	}

	return nil
}

// ValidateUserIdDriveIdListContentTypeIdBaseTypeID checks that 'input' can be parsed as a User Id Drive Id List Content Type Id Base Type ID
func ValidateUserIdDriveIdListContentTypeIdBaseTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListContentTypeIdBaseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Content Type Id Base Type ID
func (id UserIdDriveIdListContentTypeIdBaseTypeId) ID() string {
	fmtString := "/users/%s/drives/%s/list/contentTypes/%s/baseTypes/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ContentTypeId, id.ContentTypeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Content Type Id Base Type ID
func (id UserIdDriveIdListContentTypeIdBaseTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("baseTypes", "baseTypes", "baseTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId1", "contentTypeId1"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Content Type Id Base Type ID
func (id UserIdDriveIdListContentTypeIdBaseTypeId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Type Id 1: %q", id.ContentTypeId1),
	}
	return fmt.Sprintf("User Id Drive Id List Content Type Id Base Type (%s)", strings.Join(components, "\n"))
}
