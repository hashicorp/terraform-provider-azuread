package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListContentTypeIdColumnLinkId{}

// UserIdDriveIdListContentTypeIdColumnLinkId is a struct representing the Resource ID for a User Id Drive Id List Content Type Id Column Link
type UserIdDriveIdListContentTypeIdColumnLinkId struct {
	UserId        string
	DriveId       string
	ContentTypeId string
	ColumnLinkId  string
}

// NewUserIdDriveIdListContentTypeIdColumnLinkID returns a new UserIdDriveIdListContentTypeIdColumnLinkId struct
func NewUserIdDriveIdListContentTypeIdColumnLinkID(userId string, driveId string, contentTypeId string, columnLinkId string) UserIdDriveIdListContentTypeIdColumnLinkId {
	return UserIdDriveIdListContentTypeIdColumnLinkId{
		UserId:        userId,
		DriveId:       driveId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseUserIdDriveIdListContentTypeIdColumnLinkID parses 'input' into a UserIdDriveIdListContentTypeIdColumnLinkId
func ParseUserIdDriveIdListContentTypeIdColumnLinkID(input string) (*UserIdDriveIdListContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListContentTypeIdColumnLinkIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListContentTypeIdColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListContentTypeIdColumnLinkIDInsensitively(input string) (*UserIdDriveIdListContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListContentTypeIdColumnLinkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ColumnLinkId, ok = input.Parsed["columnLinkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListContentTypeIdColumnLinkID checks that 'input' can be parsed as a User Id Drive Id List Content Type Id Column Link ID
func ValidateUserIdDriveIdListContentTypeIdColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListContentTypeIdColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Content Type Id Column Link ID
func (id UserIdDriveIdListContentTypeIdColumnLinkId) ID() string {
	fmtString := "/users/%s/drives/%s/list/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Content Type Id Column Link ID
func (id UserIdDriveIdListContentTypeIdColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Content Type Id Column Link ID
func (id UserIdDriveIdListContentTypeIdColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("User Id Drive Id List Content Type Id Column Link (%s)", strings.Join(components, "\n"))
}
