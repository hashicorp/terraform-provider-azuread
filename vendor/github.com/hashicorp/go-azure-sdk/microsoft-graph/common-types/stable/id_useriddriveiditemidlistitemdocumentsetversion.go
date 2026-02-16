package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemDocumentSetVersionId{}

// UserIdDriveIdItemIdListItemDocumentSetVersionId is a struct representing the Resource ID for a User Id Drive Id Item Id List Item Document Set Version
type UserIdDriveIdItemIdListItemDocumentSetVersionId struct {
	UserId               string
	DriveId              string
	DriveItemId          string
	DocumentSetVersionId string
}

// NewUserIdDriveIdItemIdListItemDocumentSetVersionID returns a new UserIdDriveIdItemIdListItemDocumentSetVersionId struct
func NewUserIdDriveIdItemIdListItemDocumentSetVersionID(userId string, driveId string, driveItemId string, documentSetVersionId string) UserIdDriveIdItemIdListItemDocumentSetVersionId {
	return UserIdDriveIdItemIdListItemDocumentSetVersionId{
		UserId:               userId,
		DriveId:              driveId,
		DriveItemId:          driveItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseUserIdDriveIdItemIdListItemDocumentSetVersionID parses 'input' into a UserIdDriveIdItemIdListItemDocumentSetVersionId
func ParseUserIdDriveIdItemIdListItemDocumentSetVersionID(input string) (*UserIdDriveIdItemIdListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdListItemDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdListItemDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdListItemDocumentSetVersionIDInsensitively(input string) (*UserIdDriveIdItemIdListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdListItemDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdListItemDocumentSetVersionID checks that 'input' can be parsed as a User Id Drive Id Item Id List Item Document Set Version ID
func ValidateUserIdDriveIdItemIdListItemDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdListItemDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id List Item Document Set Version ID
func (id UserIdDriveIdItemIdListItemDocumentSetVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/listItem/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id List Item Document Set Version ID
func (id UserIdDriveIdItemIdListItemDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id List Item Document Set Version ID
func (id UserIdDriveIdItemIdListItemDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id List Item Document Set Version (%s)", strings.Join(components, "\n"))
}
