package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListItemIdDocumentSetVersionId{}

// UserIdDriveIdListItemIdDocumentSetVersionId is a struct representing the Resource ID for a User Id Drive Id List Item Id Document Set Version
type UserIdDriveIdListItemIdDocumentSetVersionId struct {
	UserId               string
	DriveId              string
	ListItemId           string
	DocumentSetVersionId string
}

// NewUserIdDriveIdListItemIdDocumentSetVersionID returns a new UserIdDriveIdListItemIdDocumentSetVersionId struct
func NewUserIdDriveIdListItemIdDocumentSetVersionID(userId string, driveId string, listItemId string, documentSetVersionId string) UserIdDriveIdListItemIdDocumentSetVersionId {
	return UserIdDriveIdListItemIdDocumentSetVersionId{
		UserId:               userId,
		DriveId:              driveId,
		ListItemId:           listItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseUserIdDriveIdListItemIdDocumentSetVersionID parses 'input' into a UserIdDriveIdListItemIdDocumentSetVersionId
func ParseUserIdDriveIdListItemIdDocumentSetVersionID(input string) (*UserIdDriveIdListItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListItemIdDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListItemIdDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListItemIdDocumentSetVersionIDInsensitively(input string) (*UserIdDriveIdListItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListItemIdDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListItemIdDocumentSetVersionID checks that 'input' can be parsed as a User Id Drive Id List Item Id Document Set Version ID
func ValidateUserIdDriveIdListItemIdDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListItemIdDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Item Id Document Set Version ID
func (id UserIdDriveIdListItemIdDocumentSetVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/list/items/%s/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ListItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Item Id Document Set Version ID
func (id UserIdDriveIdListItemIdDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Item Id Document Set Version ID
func (id UserIdDriveIdListItemIdDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("User Id Drive Id List Item Id Document Set Version (%s)", strings.Join(components, "\n"))
}
